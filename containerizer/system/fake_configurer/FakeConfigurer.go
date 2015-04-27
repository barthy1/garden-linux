// This file was generated by counterfeiter
package fake_configurer

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/containerizer/system"
)

type FakeConfigurer struct {
	ConfigureStub        func() error
	configureMutex       sync.RWMutex
	configureArgsForCall []struct{}
	configureReturns     struct {
		result1 error
	}
}

func (fake *FakeConfigurer) Configure() error {
	fake.configureMutex.Lock()
	fake.configureArgsForCall = append(fake.configureArgsForCall, struct{}{})
	fake.configureMutex.Unlock()
	if fake.ConfigureStub != nil {
		return fake.ConfigureStub()
	} else {
		return fake.configureReturns.result1
	}
}

func (fake *FakeConfigurer) ConfigureCallCount() int {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return len(fake.configureArgsForCall)
}

func (fake *FakeConfigurer) ConfigureReturns(result1 error) {
	fake.ConfigureStub = nil
	fake.configureReturns = struct {
		result1 error
	}{result1}
}

var _ system.Configurer = new(FakeConfigurer)