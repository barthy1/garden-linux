// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/garden-linux/linux_backend"
)

type FakeContainer struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	HasPropertiesStub        func(garden.Properties) bool
	hasPropertiesMutex       sync.RWMutex
	hasPropertiesArgsForCall []struct {
		arg1 garden.Properties
	}
	hasPropertiesReturns struct {
		result1 bool
	}
	GraceTimeStub        func() time.Duration
	graceTimeMutex       sync.RWMutex
	graceTimeArgsForCall []struct{}
	graceTimeReturns     struct {
		result1 time.Duration
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	SnapshotStub        func(io.Writer) error
	snapshotMutex       sync.RWMutex
	snapshotArgsForCall []struct {
		arg1 io.Writer
	}
	snapshotReturns struct {
		result1 error
	}
	ResourceSpecStub        func() linux_backend.LinuxContainerSpec
	resourceSpecMutex       sync.RWMutex
	resourceSpecArgsForCall []struct{}
	resourceSpecReturns     struct {
		result1 linux_backend.LinuxContainerSpec
	}
	RestoreStub        func(linux_backend.LinuxContainerSpec) error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct {
		arg1 linux_backend.LinuxContainerSpec
	}
	restoreReturns struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns     struct {
		result1 string
	}
	StopStub        func(kill bool) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		kill bool
	}
	stopReturns struct {
		result1 error
	}
	InfoStub        func() (garden.ContainerInfo, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct{}
	infoReturns     struct {
		result1 garden.ContainerInfo
		result2 error
	}
	StreamInStub        func(spec garden.StreamInSpec) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		spec garden.StreamInSpec
	}
	streamInReturns struct {
		result1 error
	}
	StreamOutStub        func(spec garden.StreamOutSpec) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		spec garden.StreamOutSpec
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	LimitBandwidthStub        func(limits garden.BandwidthLimits) error
	limitBandwidthMutex       sync.RWMutex
	limitBandwidthArgsForCall []struct {
		limits garden.BandwidthLimits
	}
	limitBandwidthReturns struct {
		result1 error
	}
	CurrentBandwidthLimitsStub        func() (garden.BandwidthLimits, error)
	currentBandwidthLimitsMutex       sync.RWMutex
	currentBandwidthLimitsArgsForCall []struct{}
	currentBandwidthLimitsReturns     struct {
		result1 garden.BandwidthLimits
		result2 error
	}
	LimitCPUStub        func(limits garden.CPULimits) error
	limitCPUMutex       sync.RWMutex
	limitCPUArgsForCall []struct {
		limits garden.CPULimits
	}
	limitCPUReturns struct {
		result1 error
	}
	CurrentCPULimitsStub        func() (garden.CPULimits, error)
	currentCPULimitsMutex       sync.RWMutex
	currentCPULimitsArgsForCall []struct{}
	currentCPULimitsReturns     struct {
		result1 garden.CPULimits
		result2 error
	}
	LimitDiskStub        func(limits garden.DiskLimits) error
	limitDiskMutex       sync.RWMutex
	limitDiskArgsForCall []struct {
		limits garden.DiskLimits
	}
	limitDiskReturns struct {
		result1 error
	}
	CurrentDiskLimitsStub        func() (garden.DiskLimits, error)
	currentDiskLimitsMutex       sync.RWMutex
	currentDiskLimitsArgsForCall []struct{}
	currentDiskLimitsReturns     struct {
		result1 garden.DiskLimits
		result2 error
	}
	LimitMemoryStub        func(limits garden.MemoryLimits) error
	limitMemoryMutex       sync.RWMutex
	limitMemoryArgsForCall []struct {
		limits garden.MemoryLimits
	}
	limitMemoryReturns struct {
		result1 error
	}
	CurrentMemoryLimitsStub        func() (garden.MemoryLimits, error)
	currentMemoryLimitsMutex       sync.RWMutex
	currentMemoryLimitsArgsForCall []struct{}
	currentMemoryLimitsReturns     struct {
		result1 garden.MemoryLimits
		result2 error
	}
	NetInStub        func(hostPort, containerPort uint32) (uint32, uint32, error)
	netInMutex       sync.RWMutex
	netInArgsForCall []struct {
		hostPort      uint32
		containerPort uint32
	}
	netInReturns struct {
		result1 uint32
		result2 uint32
		result3 error
	}
	NetOutStub        func(netOutRule garden.NetOutRule) error
	netOutMutex       sync.RWMutex
	netOutArgsForCall []struct {
		netOutRule garden.NetOutRule
	}
	netOutReturns struct {
		result1 error
	}
	RunStub        func(garden.ProcessSpec, garden.ProcessIO) (garden.Process, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 garden.ProcessSpec
		arg2 garden.ProcessIO
	}
	runReturns struct {
		result1 garden.Process
		result2 error
	}
	AttachStub        func(processID string, io garden.ProcessIO) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		processID string
		io        garden.ProcessIO
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	MetricsStub        func() (garden.Metrics, error)
	metricsMutex       sync.RWMutex
	metricsArgsForCall []struct{}
	metricsReturns     struct {
		result1 garden.Metrics
		result2 error
	}
	SetGraceTimeStub        func(graceTime time.Duration) error
	setGraceTimeMutex       sync.RWMutex
	setGraceTimeArgsForCall []struct {
		graceTime time.Duration
	}
	setGraceTimeReturns struct {
		result1 error
	}
	PropertiesStub        func() (garden.Properties, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct{}
	propertiesReturns     struct {
		result1 garden.Properties
		result2 error
	}
	PropertyStub        func(name string) (string, error)
	propertyMutex       sync.RWMutex
	propertyArgsForCall []struct {
		name string
	}
	propertyReturns struct {
		result1 string
		result2 error
	}
	SetPropertyStub        func(name string, value string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		name  string
		value string
	}
	setPropertyReturns struct {
		result1 error
	}
	RemovePropertyStub        func(name string) error
	removePropertyMutex       sync.RWMutex
	removePropertyArgsForCall []struct {
		name string
	}
	removePropertyReturns struct {
		result1 error
	}
}

func (fake *FakeContainer) ID() string {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeContainer) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeContainer) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContainer) HasProperties(arg1 garden.Properties) bool {
	fake.hasPropertiesMutex.Lock()
	fake.hasPropertiesArgsForCall = append(fake.hasPropertiesArgsForCall, struct {
		arg1 garden.Properties
	}{arg1})
	fake.hasPropertiesMutex.Unlock()
	if fake.HasPropertiesStub != nil {
		return fake.HasPropertiesStub(arg1)
	} else {
		return fake.hasPropertiesReturns.result1
	}
}

func (fake *FakeContainer) HasPropertiesCallCount() int {
	fake.hasPropertiesMutex.RLock()
	defer fake.hasPropertiesMutex.RUnlock()
	return len(fake.hasPropertiesArgsForCall)
}

func (fake *FakeContainer) HasPropertiesArgsForCall(i int) garden.Properties {
	fake.hasPropertiesMutex.RLock()
	defer fake.hasPropertiesMutex.RUnlock()
	return fake.hasPropertiesArgsForCall[i].arg1
}

func (fake *FakeContainer) HasPropertiesReturns(result1 bool) {
	fake.HasPropertiesStub = nil
	fake.hasPropertiesReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainer) GraceTime() time.Duration {
	fake.graceTimeMutex.Lock()
	fake.graceTimeArgsForCall = append(fake.graceTimeArgsForCall, struct{}{})
	fake.graceTimeMutex.Unlock()
	if fake.GraceTimeStub != nil {
		return fake.GraceTimeStub()
	} else {
		return fake.graceTimeReturns.result1
	}
}

func (fake *FakeContainer) GraceTimeCallCount() int {
	fake.graceTimeMutex.RLock()
	defer fake.graceTimeMutex.RUnlock()
	return len(fake.graceTimeArgsForCall)
}

func (fake *FakeContainer) GraceTimeReturns(result1 time.Duration) {
	fake.GraceTimeStub = nil
	fake.graceTimeReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeContainer) Start() error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeContainer) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeContainer) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Snapshot(arg1 io.Writer) error {
	fake.snapshotMutex.Lock()
	fake.snapshotArgsForCall = append(fake.snapshotArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	fake.snapshotMutex.Unlock()
	if fake.SnapshotStub != nil {
		return fake.SnapshotStub(arg1)
	} else {
		return fake.snapshotReturns.result1
	}
}

func (fake *FakeContainer) SnapshotCallCount() int {
	fake.snapshotMutex.RLock()
	defer fake.snapshotMutex.RUnlock()
	return len(fake.snapshotArgsForCall)
}

func (fake *FakeContainer) SnapshotArgsForCall(i int) io.Writer {
	fake.snapshotMutex.RLock()
	defer fake.snapshotMutex.RUnlock()
	return fake.snapshotArgsForCall[i].arg1
}

func (fake *FakeContainer) SnapshotReturns(result1 error) {
	fake.SnapshotStub = nil
	fake.snapshotReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) ResourceSpec() linux_backend.LinuxContainerSpec {
	fake.resourceSpecMutex.Lock()
	fake.resourceSpecArgsForCall = append(fake.resourceSpecArgsForCall, struct{}{})
	fake.resourceSpecMutex.Unlock()
	if fake.ResourceSpecStub != nil {
		return fake.ResourceSpecStub()
	} else {
		return fake.resourceSpecReturns.result1
	}
}

func (fake *FakeContainer) ResourceSpecCallCount() int {
	fake.resourceSpecMutex.RLock()
	defer fake.resourceSpecMutex.RUnlock()
	return len(fake.resourceSpecArgsForCall)
}

func (fake *FakeContainer) ResourceSpecReturns(result1 linux_backend.LinuxContainerSpec) {
	fake.ResourceSpecStub = nil
	fake.resourceSpecReturns = struct {
		result1 linux_backend.LinuxContainerSpec
	}{result1}
}

func (fake *FakeContainer) Restore(arg1 linux_backend.LinuxContainerSpec) error {
	fake.restoreMutex.Lock()
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct {
		arg1 linux_backend.LinuxContainerSpec
	}{arg1})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub(arg1)
	} else {
		return fake.restoreReturns.result1
	}
}

func (fake *FakeContainer) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeContainer) RestoreArgsForCall(i int) linux_backend.LinuxContainerSpec {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return fake.restoreArgsForCall[i].arg1
}

func (fake *FakeContainer) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Cleanup() error {
	fake.cleanupMutex.Lock()
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	} else {
		return fake.cleanupReturns.result1
	}
}

func (fake *FakeContainer) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeContainer) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Handle() string {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	} else {
		return fake.handleReturns.result1
	}
}

func (fake *FakeContainer) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeContainer) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContainer) Stop(kill bool) error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		kill bool
	}{kill})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(kill)
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeContainer) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeContainer) StopArgsForCall(i int) bool {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].kill
}

func (fake *FakeContainer) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Info() (garden.ContainerInfo, error) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	} else {
		return fake.infoReturns.result1, fake.infoReturns.result2
	}
}

func (fake *FakeContainer) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeContainer) InfoReturns(result1 garden.ContainerInfo, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 garden.ContainerInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) StreamIn(spec garden.StreamInSpec) error {
	fake.streamInMutex.Lock()
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		spec garden.StreamInSpec
	}{spec})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(spec)
	} else {
		return fake.streamInReturns.result1
	}
}

func (fake *FakeContainer) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeContainer) StreamInArgsForCall(i int) garden.StreamInSpec {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].spec
}

func (fake *FakeContainer) StreamInReturns(result1 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) StreamOut(spec garden.StreamOutSpec) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		spec garden.StreamOutSpec
	}{spec})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(spec)
	} else {
		return fake.streamOutReturns.result1, fake.streamOutReturns.result2
	}
}

func (fake *FakeContainer) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeContainer) StreamOutArgsForCall(i int) garden.StreamOutSpec {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].spec
}

func (fake *FakeContainer) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitBandwidth(limits garden.BandwidthLimits) error {
	fake.limitBandwidthMutex.Lock()
	fake.limitBandwidthArgsForCall = append(fake.limitBandwidthArgsForCall, struct {
		limits garden.BandwidthLimits
	}{limits})
	fake.limitBandwidthMutex.Unlock()
	if fake.LimitBandwidthStub != nil {
		return fake.LimitBandwidthStub(limits)
	} else {
		return fake.limitBandwidthReturns.result1
	}
}

func (fake *FakeContainer) LimitBandwidthCallCount() int {
	fake.limitBandwidthMutex.RLock()
	defer fake.limitBandwidthMutex.RUnlock()
	return len(fake.limitBandwidthArgsForCall)
}

func (fake *FakeContainer) LimitBandwidthArgsForCall(i int) garden.BandwidthLimits {
	fake.limitBandwidthMutex.RLock()
	defer fake.limitBandwidthMutex.RUnlock()
	return fake.limitBandwidthArgsForCall[i].limits
}

func (fake *FakeContainer) LimitBandwidthReturns(result1 error) {
	fake.LimitBandwidthStub = nil
	fake.limitBandwidthReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentBandwidthLimits() (garden.BandwidthLimits, error) {
	fake.currentBandwidthLimitsMutex.Lock()
	fake.currentBandwidthLimitsArgsForCall = append(fake.currentBandwidthLimitsArgsForCall, struct{}{})
	fake.currentBandwidthLimitsMutex.Unlock()
	if fake.CurrentBandwidthLimitsStub != nil {
		return fake.CurrentBandwidthLimitsStub()
	} else {
		return fake.currentBandwidthLimitsReturns.result1, fake.currentBandwidthLimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentBandwidthLimitsCallCount() int {
	fake.currentBandwidthLimitsMutex.RLock()
	defer fake.currentBandwidthLimitsMutex.RUnlock()
	return len(fake.currentBandwidthLimitsArgsForCall)
}

func (fake *FakeContainer) CurrentBandwidthLimitsReturns(result1 garden.BandwidthLimits, result2 error) {
	fake.CurrentBandwidthLimitsStub = nil
	fake.currentBandwidthLimitsReturns = struct {
		result1 garden.BandwidthLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitCPU(limits garden.CPULimits) error {
	fake.limitCPUMutex.Lock()
	fake.limitCPUArgsForCall = append(fake.limitCPUArgsForCall, struct {
		limits garden.CPULimits
	}{limits})
	fake.limitCPUMutex.Unlock()
	if fake.LimitCPUStub != nil {
		return fake.LimitCPUStub(limits)
	} else {
		return fake.limitCPUReturns.result1
	}
}

func (fake *FakeContainer) LimitCPUCallCount() int {
	fake.limitCPUMutex.RLock()
	defer fake.limitCPUMutex.RUnlock()
	return len(fake.limitCPUArgsForCall)
}

func (fake *FakeContainer) LimitCPUArgsForCall(i int) garden.CPULimits {
	fake.limitCPUMutex.RLock()
	defer fake.limitCPUMutex.RUnlock()
	return fake.limitCPUArgsForCall[i].limits
}

func (fake *FakeContainer) LimitCPUReturns(result1 error) {
	fake.LimitCPUStub = nil
	fake.limitCPUReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentCPULimits() (garden.CPULimits, error) {
	fake.currentCPULimitsMutex.Lock()
	fake.currentCPULimitsArgsForCall = append(fake.currentCPULimitsArgsForCall, struct{}{})
	fake.currentCPULimitsMutex.Unlock()
	if fake.CurrentCPULimitsStub != nil {
		return fake.CurrentCPULimitsStub()
	} else {
		return fake.currentCPULimitsReturns.result1, fake.currentCPULimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentCPULimitsCallCount() int {
	fake.currentCPULimitsMutex.RLock()
	defer fake.currentCPULimitsMutex.RUnlock()
	return len(fake.currentCPULimitsArgsForCall)
}

func (fake *FakeContainer) CurrentCPULimitsReturns(result1 garden.CPULimits, result2 error) {
	fake.CurrentCPULimitsStub = nil
	fake.currentCPULimitsReturns = struct {
		result1 garden.CPULimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitDisk(limits garden.DiskLimits) error {
	fake.limitDiskMutex.Lock()
	fake.limitDiskArgsForCall = append(fake.limitDiskArgsForCall, struct {
		limits garden.DiskLimits
	}{limits})
	fake.limitDiskMutex.Unlock()
	if fake.LimitDiskStub != nil {
		return fake.LimitDiskStub(limits)
	} else {
		return fake.limitDiskReturns.result1
	}
}

func (fake *FakeContainer) LimitDiskCallCount() int {
	fake.limitDiskMutex.RLock()
	defer fake.limitDiskMutex.RUnlock()
	return len(fake.limitDiskArgsForCall)
}

func (fake *FakeContainer) LimitDiskArgsForCall(i int) garden.DiskLimits {
	fake.limitDiskMutex.RLock()
	defer fake.limitDiskMutex.RUnlock()
	return fake.limitDiskArgsForCall[i].limits
}

func (fake *FakeContainer) LimitDiskReturns(result1 error) {
	fake.LimitDiskStub = nil
	fake.limitDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentDiskLimits() (garden.DiskLimits, error) {
	fake.currentDiskLimitsMutex.Lock()
	fake.currentDiskLimitsArgsForCall = append(fake.currentDiskLimitsArgsForCall, struct{}{})
	fake.currentDiskLimitsMutex.Unlock()
	if fake.CurrentDiskLimitsStub != nil {
		return fake.CurrentDiskLimitsStub()
	} else {
		return fake.currentDiskLimitsReturns.result1, fake.currentDiskLimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentDiskLimitsCallCount() int {
	fake.currentDiskLimitsMutex.RLock()
	defer fake.currentDiskLimitsMutex.RUnlock()
	return len(fake.currentDiskLimitsArgsForCall)
}

func (fake *FakeContainer) CurrentDiskLimitsReturns(result1 garden.DiskLimits, result2 error) {
	fake.CurrentDiskLimitsStub = nil
	fake.currentDiskLimitsReturns = struct {
		result1 garden.DiskLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitMemory(limits garden.MemoryLimits) error {
	fake.limitMemoryMutex.Lock()
	fake.limitMemoryArgsForCall = append(fake.limitMemoryArgsForCall, struct {
		limits garden.MemoryLimits
	}{limits})
	fake.limitMemoryMutex.Unlock()
	if fake.LimitMemoryStub != nil {
		return fake.LimitMemoryStub(limits)
	} else {
		return fake.limitMemoryReturns.result1
	}
}

func (fake *FakeContainer) LimitMemoryCallCount() int {
	fake.limitMemoryMutex.RLock()
	defer fake.limitMemoryMutex.RUnlock()
	return len(fake.limitMemoryArgsForCall)
}

func (fake *FakeContainer) LimitMemoryArgsForCall(i int) garden.MemoryLimits {
	fake.limitMemoryMutex.RLock()
	defer fake.limitMemoryMutex.RUnlock()
	return fake.limitMemoryArgsForCall[i].limits
}

func (fake *FakeContainer) LimitMemoryReturns(result1 error) {
	fake.LimitMemoryStub = nil
	fake.limitMemoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentMemoryLimits() (garden.MemoryLimits, error) {
	fake.currentMemoryLimitsMutex.Lock()
	fake.currentMemoryLimitsArgsForCall = append(fake.currentMemoryLimitsArgsForCall, struct{}{})
	fake.currentMemoryLimitsMutex.Unlock()
	if fake.CurrentMemoryLimitsStub != nil {
		return fake.CurrentMemoryLimitsStub()
	} else {
		return fake.currentMemoryLimitsReturns.result1, fake.currentMemoryLimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentMemoryLimitsCallCount() int {
	fake.currentMemoryLimitsMutex.RLock()
	defer fake.currentMemoryLimitsMutex.RUnlock()
	return len(fake.currentMemoryLimitsArgsForCall)
}

func (fake *FakeContainer) CurrentMemoryLimitsReturns(result1 garden.MemoryLimits, result2 error) {
	fake.CurrentMemoryLimitsStub = nil
	fake.currentMemoryLimitsReturns = struct {
		result1 garden.MemoryLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) NetIn(hostPort uint32, containerPort uint32) (uint32, uint32, error) {
	fake.netInMutex.Lock()
	fake.netInArgsForCall = append(fake.netInArgsForCall, struct {
		hostPort      uint32
		containerPort uint32
	}{hostPort, containerPort})
	fake.netInMutex.Unlock()
	if fake.NetInStub != nil {
		return fake.NetInStub(hostPort, containerPort)
	} else {
		return fake.netInReturns.result1, fake.netInReturns.result2, fake.netInReturns.result3
	}
}

func (fake *FakeContainer) NetInCallCount() int {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return len(fake.netInArgsForCall)
}

func (fake *FakeContainer) NetInArgsForCall(i int) (uint32, uint32) {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return fake.netInArgsForCall[i].hostPort, fake.netInArgsForCall[i].containerPort
}

func (fake *FakeContainer) NetInReturns(result1 uint32, result2 uint32, result3 error) {
	fake.NetInStub = nil
	fake.netInReturns = struct {
		result1 uint32
		result2 uint32
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainer) NetOut(netOutRule garden.NetOutRule) error {
	fake.netOutMutex.Lock()
	fake.netOutArgsForCall = append(fake.netOutArgsForCall, struct {
		netOutRule garden.NetOutRule
	}{netOutRule})
	fake.netOutMutex.Unlock()
	if fake.NetOutStub != nil {
		return fake.NetOutStub(netOutRule)
	} else {
		return fake.netOutReturns.result1
	}
}

func (fake *FakeContainer) NetOutCallCount() int {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return len(fake.netOutArgsForCall)
}

func (fake *FakeContainer) NetOutArgsForCall(i int) garden.NetOutRule {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return fake.netOutArgsForCall[i].netOutRule
}

func (fake *FakeContainer) NetOutReturns(result1 error) {
	fake.NetOutStub = nil
	fake.netOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Run(arg1 garden.ProcessSpec, arg2 garden.ProcessIO) (garden.Process, error) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 garden.ProcessSpec
		arg2 garden.ProcessIO
	}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	} else {
		return fake.runReturns.result1, fake.runReturns.result2
	}
}

func (fake *FakeContainer) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeContainer) RunArgsForCall(i int) (garden.ProcessSpec, garden.ProcessIO) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2
}

func (fake *FakeContainer) RunReturns(result1 garden.Process, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Attach(processID string, io garden.ProcessIO) (garden.Process, error) {
	fake.attachMutex.Lock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		processID string
		io        garden.ProcessIO
	}{processID, io})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(processID, io)
	} else {
		return fake.attachReturns.result1, fake.attachReturns.result2
	}
}

func (fake *FakeContainer) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeContainer) AttachArgsForCall(i int) (string, garden.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].processID, fake.attachArgsForCall[i].io
}

func (fake *FakeContainer) AttachReturns(result1 garden.Process, result2 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Metrics() (garden.Metrics, error) {
	fake.metricsMutex.Lock()
	fake.metricsArgsForCall = append(fake.metricsArgsForCall, struct{}{})
	fake.metricsMutex.Unlock()
	if fake.MetricsStub != nil {
		return fake.MetricsStub()
	} else {
		return fake.metricsReturns.result1, fake.metricsReturns.result2
	}
}

func (fake *FakeContainer) MetricsCallCount() int {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return len(fake.metricsArgsForCall)
}

func (fake *FakeContainer) MetricsReturns(result1 garden.Metrics, result2 error) {
	fake.MetricsStub = nil
	fake.metricsReturns = struct {
		result1 garden.Metrics
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) SetGraceTime(graceTime time.Duration) error {
	fake.setGraceTimeMutex.Lock()
	fake.setGraceTimeArgsForCall = append(fake.setGraceTimeArgsForCall, struct {
		graceTime time.Duration
	}{graceTime})
	fake.setGraceTimeMutex.Unlock()
	if fake.SetGraceTimeStub != nil {
		return fake.SetGraceTimeStub(graceTime)
	} else {
		return fake.setGraceTimeReturns.result1
	}
}

func (fake *FakeContainer) SetGraceTimeCallCount() int {
	fake.setGraceTimeMutex.RLock()
	defer fake.setGraceTimeMutex.RUnlock()
	return len(fake.setGraceTimeArgsForCall)
}

func (fake *FakeContainer) SetGraceTimeArgsForCall(i int) time.Duration {
	fake.setGraceTimeMutex.RLock()
	defer fake.setGraceTimeMutex.RUnlock()
	return fake.setGraceTimeArgsForCall[i].graceTime
}

func (fake *FakeContainer) SetGraceTimeReturns(result1 error) {
	fake.SetGraceTimeStub = nil
	fake.setGraceTimeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Properties() (garden.Properties, error) {
	fake.propertiesMutex.Lock()
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct{}{})
	fake.propertiesMutex.Unlock()
	if fake.PropertiesStub != nil {
		return fake.PropertiesStub()
	} else {
		return fake.propertiesReturns.result1, fake.propertiesReturns.result2
	}
}

func (fake *FakeContainer) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *FakeContainer) PropertiesReturns(result1 garden.Properties, result2 error) {
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 garden.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Property(name string) (string, error) {
	fake.propertyMutex.Lock()
	fake.propertyArgsForCall = append(fake.propertyArgsForCall, struct {
		name string
	}{name})
	fake.propertyMutex.Unlock()
	if fake.PropertyStub != nil {
		return fake.PropertyStub(name)
	} else {
		return fake.propertyReturns.result1, fake.propertyReturns.result2
	}
}

func (fake *FakeContainer) PropertyCallCount() int {
	fake.propertyMutex.RLock()
	defer fake.propertyMutex.RUnlock()
	return len(fake.propertyArgsForCall)
}

func (fake *FakeContainer) PropertyArgsForCall(i int) string {
	fake.propertyMutex.RLock()
	defer fake.propertyMutex.RUnlock()
	return fake.propertyArgsForCall[i].name
}

func (fake *FakeContainer) PropertyReturns(result1 string, result2 error) {
	fake.PropertyStub = nil
	fake.propertyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) SetProperty(name string, value string) error {
	fake.setPropertyMutex.Lock()
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		name  string
		value string
	}{name, value})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(name, value)
	} else {
		return fake.setPropertyReturns.result1
	}
}

func (fake *FakeContainer) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeContainer) SetPropertyArgsForCall(i int) (string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return fake.setPropertyArgsForCall[i].name, fake.setPropertyArgsForCall[i].value
}

func (fake *FakeContainer) SetPropertyReturns(result1 error) {
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) RemoveProperty(name string) error {
	fake.removePropertyMutex.Lock()
	fake.removePropertyArgsForCall = append(fake.removePropertyArgsForCall, struct {
		name string
	}{name})
	fake.removePropertyMutex.Unlock()
	if fake.RemovePropertyStub != nil {
		return fake.RemovePropertyStub(name)
	} else {
		return fake.removePropertyReturns.result1
	}
}

func (fake *FakeContainer) RemovePropertyCallCount() int {
	fake.removePropertyMutex.RLock()
	defer fake.removePropertyMutex.RUnlock()
	return len(fake.removePropertyArgsForCall)
}

func (fake *FakeContainer) RemovePropertyArgsForCall(i int) string {
	fake.removePropertyMutex.RLock()
	defer fake.removePropertyMutex.RUnlock()
	return fake.removePropertyArgsForCall[i].name
}

func (fake *FakeContainer) RemovePropertyReturns(result1 error) {
	fake.RemovePropertyStub = nil
	fake.removePropertyReturns = struct {
		result1 error
	}{result1}
}

var _ linux_backend.Container = new(FakeContainer)
