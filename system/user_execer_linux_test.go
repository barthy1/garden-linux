package system_test

import (
	"io"
	"os/exec"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("UserExecer", func() {

	var (
		id int
	)

	BeforeEach(func() {
		id = 100000 + GinkgoParallelNode()
		Expect(exec.Command("groupadd", "-g", fmt.Sprintf("%d", id), "banana").Run()).To(Succeed())
		Expect(exec.Command(
			"useradd",
			"-g", fmt.Sprintf("%d", id),
			"-u", fmt.Sprintf("%d", id),
			"banana",
		).Run()).To(Succeed())
	})

	AfterEach(func() {
		Expect(exec.Command("userdel", "banana").Run()).To(Succeed())
	})

	It("execs a process as specified user", func() {
		out := gbytes.NewBuffer()
		runningTest, err := gexec.Start(
			exec.Command(testUserExecerPath, fmt.Sprintf("-uid=%d", id), fmt.Sprintf("-gid=%d", id)),
			io.MultiWriter(GinkgoWriter, out),
			io.MultiWriter(GinkgoWriter, out),
		)
		Expect(err).NotTo(HaveOccurred())
		runningTest.Wait()
		Expect(string(out.Contents())).To(Equal(fmt.Sprintf("%d\n%d\n", id, id)))
	})
})