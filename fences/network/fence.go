package network

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/cloudfoundry-incubator/garden-linux/fences"
	"github.com/cloudfoundry-incubator/garden-linux/fences/network/subnets"
	"github.com/cloudfoundry-incubator/garden/api"
)

type f struct {
	subnets.Subnets
	mtu uint32
}

type FlatFence struct {
	Ipn         string
	ContainerIP string
}

// Builds a (network) Fence from a given network spec. If the network spec
// is empty, dynamically allocates a subnet and IP. Otherwise, if the network
// spec specifies a subnet IP, allocates that subnet, and an available
// dynamic IP address. If the network has non-empty host bits, this exact IP
// address is statically allocated. In all cases, if an IP cannot be allocated which
// meets the requirements, an error is returned.
//
// The given allocation is stored in the returned fence.
func (f *f) Build(spec string) (fences.Fence, error) {
	var ipSelector subnets.IPSelector = subnets.DynamicIPSelector
	var subnetSelector subnets.SubnetSelector = subnets.DynamicSubnetSelector

	if spec != "" {
		specifiedIP, ipn, err := net.ParseCIDR(suffixIfNeeded(spec))
		if err != nil {
			return nil, err
		}

		subnetSelector = subnets.StaticSubnetSelector{ipn}

		if !specifiedIP.Equal(subnets.NetworkIP(ipn)) {
			ipSelector = subnets.StaticIPSelector{specifiedIP}
		}
	}

	subnet, containerIP, err := f.Subnets.Allocate(subnetSelector, ipSelector)
	if err != nil {
		return nil, err
	}

	return &Allocation{subnet, containerIP, f}, nil
}

func suffixIfNeeded(spec string) string {
	if !strings.Contains(spec, "/") {
		spec = spec + "/30"
	}

	return spec
}

// Rebuilds a Fence from the marshalled JSON from an existing Fence's MarshalJSON method.
// Returns an error if any of the allocations stored in the recovered fence are no longer
// available.
func (f *f) Rebuild(rm *json.RawMessage) (fences.Fence, error) {
	ff := FlatFence{}
	if err := json.Unmarshal(*rm, &ff); err != nil {
		return nil, err
	}

	_, ipn, err := net.ParseCIDR(ff.Ipn)
	if err != nil {
		return nil, err
	}

	if err := f.Subnets.Recover(ipn, net.ParseIP(ff.ContainerIP)); err != nil {
		return nil, err
	}

	return &Allocation{ipn, net.ParseIP(ff.ContainerIP), f}, nil
}

type Allocation struct {
	*net.IPNet
	containerIP net.IP
	fence       *f
}

func (a *Allocation) String() string {
	return "Allocation{" + a.IPNet.String() + ", " + a.containerIP.String() + "}"
}

func (a *Allocation) Dismantle() error {
	return a.fence.Release(a.IPNet, a.containerIP)
}

func (a *Allocation) Info(i *api.ContainerInfo) {
	i.HostIP = subnets.GatewayIP(a.IPNet).String()
	i.ContainerIP = a.containerIP.String()
}

func (a *Allocation) MarshalJSON() ([]byte, error) {
	ff := FlatFence{a.IPNet.String(), a.containerIP.String()}
	return json.Marshal(ff)
}

func (a *Allocation) ConfigureProcess(env *[]string) error {
	suff, _ := a.IPNet.Mask.Size()

	*env = append(*env, fmt.Sprintf("network_host_ip=%s", subnets.GatewayIP(a.IPNet)),
		fmt.Sprintf("network_container_ip=%s", a.containerIP),
		fmt.Sprintf("network_cidr_suffix=%d", suff),
		fmt.Sprintf("container_iface_mtu=%d", a.fence.mtu),
		fmt.Sprintf("network_cidr=%s", a.IPNet.String()))

	return nil
}