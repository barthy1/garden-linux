package repository_fetcher_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/cloudfoundry-incubator/garden-linux/repository_fetcher"
	"github.com/cloudfoundry-incubator/garden-linux/repository_fetcher/fake_lock"
	"github.com/cloudfoundry-incubator/garden-linux/resource_pool/fake_graph"
	"github.com/docker/distribution/digest"
	"github.com/docker/docker/cliconfig"
	"github.com/docker/docker/image"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/transport"
	"github.com/docker/docker/registry"
	"github.com/pivotal-golang/lager/lagertest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("RemoteV2", func() {
	var (
		fetcher      *RemoteV2Fetcher
		server       *ghttp.Server
		graph        *fake_graph.FakeGraph
		lock         *fake_lock.FakeLock
		logger       *lagertest.TestLogger
		fetchRequest *FetchRequest

		registryAddr string
	)

	BeforeEach(func() {
		graph = fake_graph.New()
		lock = new(fake_lock.FakeLock)

		logger = lagertest.NewTestLogger("test")

		server = ghttp.NewServer()
		server.RouteToHandler(
			"GET", "/v2/",
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Docker-Distribution-API-Version", "registry/2.0")
				w.Write([]byte(`{}`))
			}),
		)
		server.RouteToHandler(
			"GET", "/v1/_ping",
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(404)
			}),
		)

		registryAddr = server.HTTPTestServer.Listener.Addr().String()
		endpoint, err := registry.NewEndpoint(&registry.IndexInfo{
			Name:   registryAddr,
			Secure: false,
		}, nil)
		Expect(err).ToNot(HaveOccurred())

		tr := transport.NewTransport(
			registry.NewTransport(registry.ReceiveTimeout, endpoint.IsSecure),
		)
		session, err := registry.NewSession(registry.HTTPClient(tr), &cliconfig.AuthConfig{}, endpoint)
		Expect(err).ToNot(HaveOccurred())

		fetchRequest = &FetchRequest{
			Session:    session,
			Endpoint:   endpoint,
			Logger:     logger,
			Path:       "some-repo",
			RemotePath: "some-repo",
			Tag:        "some-tag",
		}

		fetcher = &RemoteV2Fetcher{
			Graph:     graph,
			GraphLock: lock,
		}
	})

	Context("when none of the layers already exist", func() {
		BeforeEach(func() {
			setupSuccessfulV2Fetch(server, false)
		})

		It("downloads all layers of the given tag of a repository and returns its image id", func() {
			layers := 0

			graph.WhenRegistering = func(image *image.Image, layer archive.ArchiveReader) error {
				Expect(image.ID).To(Equal(fmt.Sprintf("banana-pie-%d", layers+1)))
				parent := ""
				if layers > 0 {
					parent = fmt.Sprintf("banana-pie-%d", layers)
				}
				Expect(image.Parent).To(Equal(parent))

				layerData, err := ioutil.ReadAll(layer)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(layerData)).To(Equal(fmt.Sprintf("banana-%d-flan", layers+1)))

				layers++

				return nil
			}

			fetchResponse, err := fetcher.Fetch(fetchRequest)

			Expect(err).ToNot(HaveOccurred())
			Expect(fetchResponse.Env).To(BeEmpty())
			Expect(fetchResponse.Volumes).To(BeEmpty())
			Expect(fetchResponse.ImageID).To(Equal("banana-pie-2"))

			Expect(server.ReceivedRequests()).To(HaveLen(4))
			Expect(layers).To(Equal(2))
		})
	})

	Context("when a layer already exists", func() {
		BeforeEach(func() {
			graph.SetExists("banana-pie-1", []byte(`{"id": "banana-pie-1"}`))

			setupSuccessfulV2Fetch(server, true)
		})

		It("is not added to the graph", func() {
			layers := 0

			graph.WhenRegistering = func(image *image.Image, layer archive.ArchiveReader) error {
				Expect(image.ID).To(Equal("banana-pie-2"))
				Expect(image.Parent).To(Equal("banana-pie-1"))

				layerData, err := ioutil.ReadAll(layer)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(layerData)).To(Equal("banana-2-flan"))

				layers++

				return nil
			}

			_, err := fetcher.Fetch(fetchRequest)

			Expect(err).ToNot(HaveOccurred())
			Expect(server.ReceivedRequests()).To(HaveLen(3))
			Expect(layers).To(Equal(1))
		})
	})

	Context("when fetching manifest fails", func() {
		Context("when the manifest endpoint fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/some-repo/manifests/some-tag"),
						http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
							w.WriteHeader(500)
						}),
					),
				)
			})

			It("returns an error", func() {
				_, err := fetcher.Fetch(fetchRequest)
				Expect(err).To(MatchError(ContainSubstring("repository_fetcher: GetV2ImageManifest: could not fetch image some-repo from registry %s:", registryAddr)))
			})
		})

		Context("when the provided manifest is invalid", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/v2/some-repo/manifests/some-tag"),
						http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
							w.Write([]byte(`{\"hello}`))
						}),
					),
				)
			})

			It("returns an error", func() {
				_, err := fetcher.Fetch(fetchRequest)
				Expect(err).To(MatchError(ContainSubstring("repository_fetcher: UnmarshalManifest: could not fetch image some-repo from registry %s:", registryAddr)))
			})
		})
	})

	Context("when fetching a layer fails", func() {
		Context("when the image manifest contains an invalid layer digest", func() {
			BeforeEach(func() {
				setupSuccessfulV2Fetch(server, false)
				server.SetHandler(0, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.Write([]byte(`
					{
					   "name":"some-repo",
					   "tag":"some-tag",
					   "fsLayers":[
						  {
							 "blobSum":"barry"
						  }
					   ],
					   "history":[
						  {
							 "v1Compatibility": "{\"id\":\"banana-pie-2\"}"
						  }
					   ]
					}`))
				}))
			})

			It("returns an error", func() {
				_, err := fetcher.Fetch(fetchRequest)
				Expect(err).To(MatchError(ContainSubstring("invalid checksum digest format")))
			})
		})

		Context("when the image JSON is invalid", func() {
			BeforeEach(func() {
				setupSuccessfulV2Fetch(server, false)
				server.SetHandler(0, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.Write([]byte(`
					{
					   "name":"some-repo",
					   "tag":"some-tag",
					   "fsLayers":[
						  {
							 "blobSum":"sha256:7b3bc336724d50e1ad47bde1458ab57d170a453b0ed734087734a8eaf79c1274"
						  }
					   ],
					   "history":[
						  {
							 "v1Compatibility": "{ba}"
						  }
					   ]
					}`))
				}))
			})

			It("returns an error", func() {
				_, err := fetcher.Fetch(fetchRequest)
				Expect(err).To(MatchError(ContainSubstring("repository_fetcher: NewImgJSON: could not fetch image some-repo from registry %s:", registryAddr)))
			})
		})

		Context("when downloading the blob fails", func() {
			BeforeEach(func() {
				setupSuccessfulV2Fetch(server, false)

				server.SetHandler(1, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.WriteHeader(500)
				}))
			})

			It("returns an error", func() {
				_, err := fetcher.Fetch(fetchRequest)
				Expect(err).To(MatchError(ContainSubstring("repository_fetcher: GetV2ImageBlobReader: could not fetch image some-repo from registry %s:", registryAddr)))
			})
		})

		Context("when registering the layer with the graph fails", func() {
			BeforeEach(func() {
				setupSuccessfulV2Fetch(server, false)
				graph.WhenRegistering = func(image *image.Image, layer archive.ArchiveReader) error {
					return errors.New("oh no!")
				}
			})

			It("returns error", func() {
				_, err := fetcher.Fetch(fetchRequest)
				Expect(err).To(MatchError(ContainSubstring("oh no!")))
			})
		})
	})
})

func setupSuccessfulV2Fetch(server *ghttp.Server, layer1Cached bool) {
	layer1Data := "banana-1-flan"
	layer1Dgst, _ := digest.FromBytes([]byte(layer1Data))

	layer2Data := "banana-2-flan"
	layer2Dgst, _ := digest.FromBytes([]byte(layer2Data))

	server.AppendHandlers(
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/v2/some-repo/manifests/some-tag"),
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Write([]byte(fmt.Sprintf(`
					{
					   "name":"some-repo",
					   "tag":"some-tag",
					   "fsLayers":[
						  {
							 "blobSum":"%s"
						  },
						  {
							 "blobSum":"%s"
						  }
					   ],
					   "history":[
						  {
							 "v1Compatibility": "{\"id\":\"banana-pie-2\", \"parent\":\"banana-pie-1\"}"
						  },
						  {
							 "v1Compatibility": "{\"id\":\"banana-pie-1\"}"
						  }
					   ]
					}
					`, layer2Dgst.String(), layer1Dgst.String())))
			}),
		),
	)

	if !layer1Cached {
		server.AppendHandlers(
			ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", fmt.Sprintf("/v2/some-repo/blobs/%s", layer1Dgst)),
				http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.Write([]byte(layer1Data))
				}),
			),
		)
	}

	server.AppendHandlers(
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", fmt.Sprintf("/v2/some-repo/blobs/%s", layer2Dgst)),
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Write([]byte(layer2Data))
			}),
		),
	)
}