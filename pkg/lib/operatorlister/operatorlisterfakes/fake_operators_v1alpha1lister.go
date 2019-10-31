// Code generated by counterfeiter. DO NOT EDIT.
package operatorlisterfakes

import (
	sync "sync"

	v1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/listers/operators/v1alpha1"
	operatorlister "github.com/operator-framework/operator-lifecycle-manager/pkg/lib/operatorlister"
)

type FakeOperatorsV1alpha1Lister struct {
	CatalogSourceListerStub        func() v1alpha1.CatalogSourceLister
	catalogSourceListerMutex       sync.RWMutex
	catalogSourceListerArgsForCall []struct {
	}
	catalogSourceListerReturns struct {
		result1 v1alpha1.CatalogSourceLister
	}
	catalogSourceListerReturnsOnCall map[int]struct {
		result1 v1alpha1.CatalogSourceLister
	}
	ClusterServiceVersionListerStub        func() v1alpha1.ClusterServiceVersionLister
	clusterServiceVersionListerMutex       sync.RWMutex
	clusterServiceVersionListerArgsForCall []struct {
	}
	clusterServiceVersionListerReturns struct {
		result1 v1alpha1.ClusterServiceVersionLister
	}
	clusterServiceVersionListerReturnsOnCall map[int]struct {
		result1 v1alpha1.ClusterServiceVersionLister
	}
	InstallPlanListerStub        func() v1alpha1.InstallPlanLister
	installPlanListerMutex       sync.RWMutex
	installPlanListerArgsForCall []struct {
	}
	installPlanListerReturns struct {
		result1 v1alpha1.InstallPlanLister
	}
	installPlanListerReturnsOnCall map[int]struct {
		result1 v1alpha1.InstallPlanLister
	}
	RegisterCatalogSourceListerStub        func(string, v1alpha1.CatalogSourceLister)
	registerCatalogSourceListerMutex       sync.RWMutex
	registerCatalogSourceListerArgsForCall []struct {
		arg1 string
		arg2 v1alpha1.CatalogSourceLister
	}
	RegisterClusterServiceVersionListerStub        func(string, v1alpha1.ClusterServiceVersionLister)
	registerClusterServiceVersionListerMutex       sync.RWMutex
	registerClusterServiceVersionListerArgsForCall []struct {
		arg1 string
		arg2 v1alpha1.ClusterServiceVersionLister
	}
	RegisterInstallPlanListerStub        func(string, v1alpha1.InstallPlanLister)
	registerInstallPlanListerMutex       sync.RWMutex
	registerInstallPlanListerArgsForCall []struct {
		arg1 string
		arg2 v1alpha1.InstallPlanLister
	}
	RegisterSubscriptionListerStub        func(string, v1alpha1.SubscriptionLister)
	registerSubscriptionListerMutex       sync.RWMutex
	registerSubscriptionListerArgsForCall []struct {
		arg1 string
		arg2 v1alpha1.SubscriptionLister
	}
	SubscriptionListerStub        func() v1alpha1.SubscriptionLister
	subscriptionListerMutex       sync.RWMutex
	subscriptionListerArgsForCall []struct {
	}
	subscriptionListerReturns struct {
		result1 v1alpha1.SubscriptionLister
	}
	subscriptionListerReturnsOnCall map[int]struct {
		result1 v1alpha1.SubscriptionLister
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOperatorsV1alpha1Lister) CatalogSourceLister() v1alpha1.CatalogSourceLister {
	fake.catalogSourceListerMutex.Lock()
	ret, specificReturn := fake.catalogSourceListerReturnsOnCall[len(fake.catalogSourceListerArgsForCall)]
	fake.catalogSourceListerArgsForCall = append(fake.catalogSourceListerArgsForCall, struct {
	}{})
	fake.recordInvocation("CatalogSourceLister", []interface{}{})
	fake.catalogSourceListerMutex.Unlock()
	if fake.CatalogSourceListerStub != nil {
		return fake.CatalogSourceListerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.catalogSourceListerReturns
	return fakeReturns.result1
}

func (fake *FakeOperatorsV1alpha1Lister) CatalogSourceListerCallCount() int {
	fake.catalogSourceListerMutex.RLock()
	defer fake.catalogSourceListerMutex.RUnlock()
	return len(fake.catalogSourceListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) CatalogSourceListerCalls(stub func() v1alpha1.CatalogSourceLister) {
	fake.catalogSourceListerMutex.Lock()
	defer fake.catalogSourceListerMutex.Unlock()
	fake.CatalogSourceListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) CatalogSourceListerReturns(result1 v1alpha1.CatalogSourceLister) {
	fake.catalogSourceListerMutex.Lock()
	defer fake.catalogSourceListerMutex.Unlock()
	fake.CatalogSourceListerStub = nil
	fake.catalogSourceListerReturns = struct {
		result1 v1alpha1.CatalogSourceLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) CatalogSourceListerReturnsOnCall(i int, result1 v1alpha1.CatalogSourceLister) {
	fake.catalogSourceListerMutex.Lock()
	defer fake.catalogSourceListerMutex.Unlock()
	fake.CatalogSourceListerStub = nil
	if fake.catalogSourceListerReturnsOnCall == nil {
		fake.catalogSourceListerReturnsOnCall = make(map[int]struct {
			result1 v1alpha1.CatalogSourceLister
		})
	}
	fake.catalogSourceListerReturnsOnCall[i] = struct {
		result1 v1alpha1.CatalogSourceLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) ClusterServiceVersionLister() v1alpha1.ClusterServiceVersionLister {
	fake.clusterServiceVersionListerMutex.Lock()
	ret, specificReturn := fake.clusterServiceVersionListerReturnsOnCall[len(fake.clusterServiceVersionListerArgsForCall)]
	fake.clusterServiceVersionListerArgsForCall = append(fake.clusterServiceVersionListerArgsForCall, struct {
	}{})
	fake.recordInvocation("ClusterServiceVersionLister", []interface{}{})
	fake.clusterServiceVersionListerMutex.Unlock()
	if fake.ClusterServiceVersionListerStub != nil {
		return fake.ClusterServiceVersionListerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.clusterServiceVersionListerReturns
	return fakeReturns.result1
}

func (fake *FakeOperatorsV1alpha1Lister) ClusterServiceVersionListerCallCount() int {
	fake.clusterServiceVersionListerMutex.RLock()
	defer fake.clusterServiceVersionListerMutex.RUnlock()
	return len(fake.clusterServiceVersionListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) ClusterServiceVersionListerCalls(stub func() v1alpha1.ClusterServiceVersionLister) {
	fake.clusterServiceVersionListerMutex.Lock()
	defer fake.clusterServiceVersionListerMutex.Unlock()
	fake.ClusterServiceVersionListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) ClusterServiceVersionListerReturns(result1 v1alpha1.ClusterServiceVersionLister) {
	fake.clusterServiceVersionListerMutex.Lock()
	defer fake.clusterServiceVersionListerMutex.Unlock()
	fake.ClusterServiceVersionListerStub = nil
	fake.clusterServiceVersionListerReturns = struct {
		result1 v1alpha1.ClusterServiceVersionLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) ClusterServiceVersionListerReturnsOnCall(i int, result1 v1alpha1.ClusterServiceVersionLister) {
	fake.clusterServiceVersionListerMutex.Lock()
	defer fake.clusterServiceVersionListerMutex.Unlock()
	fake.ClusterServiceVersionListerStub = nil
	if fake.clusterServiceVersionListerReturnsOnCall == nil {
		fake.clusterServiceVersionListerReturnsOnCall = make(map[int]struct {
			result1 v1alpha1.ClusterServiceVersionLister
		})
	}
	fake.clusterServiceVersionListerReturnsOnCall[i] = struct {
		result1 v1alpha1.ClusterServiceVersionLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) InstallPlanLister() v1alpha1.InstallPlanLister {
	fake.installPlanListerMutex.Lock()
	ret, specificReturn := fake.installPlanListerReturnsOnCall[len(fake.installPlanListerArgsForCall)]
	fake.installPlanListerArgsForCall = append(fake.installPlanListerArgsForCall, struct {
	}{})
	fake.recordInvocation("InstallPlanLister", []interface{}{})
	fake.installPlanListerMutex.Unlock()
	if fake.InstallPlanListerStub != nil {
		return fake.InstallPlanListerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.installPlanListerReturns
	return fakeReturns.result1
}

func (fake *FakeOperatorsV1alpha1Lister) InstallPlanListerCallCount() int {
	fake.installPlanListerMutex.RLock()
	defer fake.installPlanListerMutex.RUnlock()
	return len(fake.installPlanListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) InstallPlanListerCalls(stub func() v1alpha1.InstallPlanLister) {
	fake.installPlanListerMutex.Lock()
	defer fake.installPlanListerMutex.Unlock()
	fake.InstallPlanListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) InstallPlanListerReturns(result1 v1alpha1.InstallPlanLister) {
	fake.installPlanListerMutex.Lock()
	defer fake.installPlanListerMutex.Unlock()
	fake.InstallPlanListerStub = nil
	fake.installPlanListerReturns = struct {
		result1 v1alpha1.InstallPlanLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) InstallPlanListerReturnsOnCall(i int, result1 v1alpha1.InstallPlanLister) {
	fake.installPlanListerMutex.Lock()
	defer fake.installPlanListerMutex.Unlock()
	fake.InstallPlanListerStub = nil
	if fake.installPlanListerReturnsOnCall == nil {
		fake.installPlanListerReturnsOnCall = make(map[int]struct {
			result1 v1alpha1.InstallPlanLister
		})
	}
	fake.installPlanListerReturnsOnCall[i] = struct {
		result1 v1alpha1.InstallPlanLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterCatalogSourceLister(arg1 string, arg2 v1alpha1.CatalogSourceLister) {
	fake.registerCatalogSourceListerMutex.Lock()
	fake.registerCatalogSourceListerArgsForCall = append(fake.registerCatalogSourceListerArgsForCall, struct {
		arg1 string
		arg2 v1alpha1.CatalogSourceLister
	}{arg1, arg2})
	fake.recordInvocation("RegisterCatalogSourceLister", []interface{}{arg1, arg2})
	fake.registerCatalogSourceListerMutex.Unlock()
	if fake.RegisterCatalogSourceListerStub != nil {
		fake.RegisterCatalogSourceListerStub(arg1, arg2)
	}
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterCatalogSourceListerCallCount() int {
	fake.registerCatalogSourceListerMutex.RLock()
	defer fake.registerCatalogSourceListerMutex.RUnlock()
	return len(fake.registerCatalogSourceListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterCatalogSourceListerCalls(stub func(string, v1alpha1.CatalogSourceLister)) {
	fake.registerCatalogSourceListerMutex.Lock()
	defer fake.registerCatalogSourceListerMutex.Unlock()
	fake.RegisterCatalogSourceListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterCatalogSourceListerArgsForCall(i int) (string, v1alpha1.CatalogSourceLister) {
	fake.registerCatalogSourceListerMutex.RLock()
	defer fake.registerCatalogSourceListerMutex.RUnlock()
	argsForCall := fake.registerCatalogSourceListerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterClusterServiceVersionLister(arg1 string, arg2 v1alpha1.ClusterServiceVersionLister) {
	fake.registerClusterServiceVersionListerMutex.Lock()
	fake.registerClusterServiceVersionListerArgsForCall = append(fake.registerClusterServiceVersionListerArgsForCall, struct {
		arg1 string
		arg2 v1alpha1.ClusterServiceVersionLister
	}{arg1, arg2})
	fake.recordInvocation("RegisterClusterServiceVersionLister", []interface{}{arg1, arg2})
	fake.registerClusterServiceVersionListerMutex.Unlock()
	if fake.RegisterClusterServiceVersionListerStub != nil {
		fake.RegisterClusterServiceVersionListerStub(arg1, arg2)
	}
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterClusterServiceVersionListerCallCount() int {
	fake.registerClusterServiceVersionListerMutex.RLock()
	defer fake.registerClusterServiceVersionListerMutex.RUnlock()
	return len(fake.registerClusterServiceVersionListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterClusterServiceVersionListerCalls(stub func(string, v1alpha1.ClusterServiceVersionLister)) {
	fake.registerClusterServiceVersionListerMutex.Lock()
	defer fake.registerClusterServiceVersionListerMutex.Unlock()
	fake.RegisterClusterServiceVersionListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterClusterServiceVersionListerArgsForCall(i int) (string, v1alpha1.ClusterServiceVersionLister) {
	fake.registerClusterServiceVersionListerMutex.RLock()
	defer fake.registerClusterServiceVersionListerMutex.RUnlock()
	argsForCall := fake.registerClusterServiceVersionListerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterInstallPlanLister(arg1 string, arg2 v1alpha1.InstallPlanLister) {
	fake.registerInstallPlanListerMutex.Lock()
	fake.registerInstallPlanListerArgsForCall = append(fake.registerInstallPlanListerArgsForCall, struct {
		arg1 string
		arg2 v1alpha1.InstallPlanLister
	}{arg1, arg2})
	fake.recordInvocation("RegisterInstallPlanLister", []interface{}{arg1, arg2})
	fake.registerInstallPlanListerMutex.Unlock()
	if fake.RegisterInstallPlanListerStub != nil {
		fake.RegisterInstallPlanListerStub(arg1, arg2)
	}
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterInstallPlanListerCallCount() int {
	fake.registerInstallPlanListerMutex.RLock()
	defer fake.registerInstallPlanListerMutex.RUnlock()
	return len(fake.registerInstallPlanListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterInstallPlanListerCalls(stub func(string, v1alpha1.InstallPlanLister)) {
	fake.registerInstallPlanListerMutex.Lock()
	defer fake.registerInstallPlanListerMutex.Unlock()
	fake.RegisterInstallPlanListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterInstallPlanListerArgsForCall(i int) (string, v1alpha1.InstallPlanLister) {
	fake.registerInstallPlanListerMutex.RLock()
	defer fake.registerInstallPlanListerMutex.RUnlock()
	argsForCall := fake.registerInstallPlanListerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterSubscriptionLister(arg1 string, arg2 v1alpha1.SubscriptionLister) {
	fake.registerSubscriptionListerMutex.Lock()
	fake.registerSubscriptionListerArgsForCall = append(fake.registerSubscriptionListerArgsForCall, struct {
		arg1 string
		arg2 v1alpha1.SubscriptionLister
	}{arg1, arg2})
	fake.recordInvocation("RegisterSubscriptionLister", []interface{}{arg1, arg2})
	fake.registerSubscriptionListerMutex.Unlock()
	if fake.RegisterSubscriptionListerStub != nil {
		fake.RegisterSubscriptionListerStub(arg1, arg2)
	}
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterSubscriptionListerCallCount() int {
	fake.registerSubscriptionListerMutex.RLock()
	defer fake.registerSubscriptionListerMutex.RUnlock()
	return len(fake.registerSubscriptionListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterSubscriptionListerCalls(stub func(string, v1alpha1.SubscriptionLister)) {
	fake.registerSubscriptionListerMutex.Lock()
	defer fake.registerSubscriptionListerMutex.Unlock()
	fake.RegisterSubscriptionListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) RegisterSubscriptionListerArgsForCall(i int) (string, v1alpha1.SubscriptionLister) {
	fake.registerSubscriptionListerMutex.RLock()
	defer fake.registerSubscriptionListerMutex.RUnlock()
	argsForCall := fake.registerSubscriptionListerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOperatorsV1alpha1Lister) SubscriptionLister() v1alpha1.SubscriptionLister {
	fake.subscriptionListerMutex.Lock()
	ret, specificReturn := fake.subscriptionListerReturnsOnCall[len(fake.subscriptionListerArgsForCall)]
	fake.subscriptionListerArgsForCall = append(fake.subscriptionListerArgsForCall, struct {
	}{})
	fake.recordInvocation("SubscriptionLister", []interface{}{})
	fake.subscriptionListerMutex.Unlock()
	if fake.SubscriptionListerStub != nil {
		return fake.SubscriptionListerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.subscriptionListerReturns
	return fakeReturns.result1
}

func (fake *FakeOperatorsV1alpha1Lister) SubscriptionListerCallCount() int {
	fake.subscriptionListerMutex.RLock()
	defer fake.subscriptionListerMutex.RUnlock()
	return len(fake.subscriptionListerArgsForCall)
}

func (fake *FakeOperatorsV1alpha1Lister) SubscriptionListerCalls(stub func() v1alpha1.SubscriptionLister) {
	fake.subscriptionListerMutex.Lock()
	defer fake.subscriptionListerMutex.Unlock()
	fake.SubscriptionListerStub = stub
}

func (fake *FakeOperatorsV1alpha1Lister) SubscriptionListerReturns(result1 v1alpha1.SubscriptionLister) {
	fake.subscriptionListerMutex.Lock()
	defer fake.subscriptionListerMutex.Unlock()
	fake.SubscriptionListerStub = nil
	fake.subscriptionListerReturns = struct {
		result1 v1alpha1.SubscriptionLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) SubscriptionListerReturnsOnCall(i int, result1 v1alpha1.SubscriptionLister) {
	fake.subscriptionListerMutex.Lock()
	defer fake.subscriptionListerMutex.Unlock()
	fake.SubscriptionListerStub = nil
	if fake.subscriptionListerReturnsOnCall == nil {
		fake.subscriptionListerReturnsOnCall = make(map[int]struct {
			result1 v1alpha1.SubscriptionLister
		})
	}
	fake.subscriptionListerReturnsOnCall[i] = struct {
		result1 v1alpha1.SubscriptionLister
	}{result1}
}

func (fake *FakeOperatorsV1alpha1Lister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.catalogSourceListerMutex.RLock()
	defer fake.catalogSourceListerMutex.RUnlock()
	fake.clusterServiceVersionListerMutex.RLock()
	defer fake.clusterServiceVersionListerMutex.RUnlock()
	fake.installPlanListerMutex.RLock()
	defer fake.installPlanListerMutex.RUnlock()
	fake.registerCatalogSourceListerMutex.RLock()
	defer fake.registerCatalogSourceListerMutex.RUnlock()
	fake.registerClusterServiceVersionListerMutex.RLock()
	defer fake.registerClusterServiceVersionListerMutex.RUnlock()
	fake.registerInstallPlanListerMutex.RLock()
	defer fake.registerInstallPlanListerMutex.RUnlock()
	fake.registerSubscriptionListerMutex.RLock()
	defer fake.registerSubscriptionListerMutex.RUnlock()
	fake.subscriptionListerMutex.RLock()
	defer fake.subscriptionListerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOperatorsV1alpha1Lister) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ operatorlister.OperatorsV1alpha1Lister = new(FakeOperatorsV1alpha1Lister)
