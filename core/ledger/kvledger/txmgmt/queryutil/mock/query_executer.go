
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/queryutil"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
)

type QueryExecuter struct {
	GetStateStub        func(namespace, key string) (*statedb.VersionedValue, error)
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct {
		namespace string
		key       string
	}
	getStateReturns struct {
		result1 *statedb.VersionedValue
		result2 error
	}
	getStateReturnsOnCall map[int]struct {
		result1 *statedb.VersionedValue
		result2 error
	}
	GetStateRangeScanIteratorStub        func(namespace, startKey, endKey string) (statedb.ResultsIterator, error)
	getStateRangeScanIteratorMutex       sync.RWMutex
	getStateRangeScanIteratorArgsForCall []struct {
		namespace string
		startKey  string
		endKey    string
	}
	getStateRangeScanIteratorReturns struct {
		result1 statedb.ResultsIterator
		result2 error
	}
	getStateRangeScanIteratorReturnsOnCall map[int]struct {
		result1 statedb.ResultsIterator
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *QueryExecuter) GetState(namespace string, key string) (*statedb.VersionedValue, error) {
	fake.getStateMutex.Lock()
	ret, specificReturn := fake.getStateReturnsOnCall[len(fake.getStateArgsForCall)]
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct {
		namespace string
		key       string
	}{namespace, key})
	fake.recordInvocation("GetState", []interface{}{namespace, key})
	fake.getStateMutex.Unlock()
	if fake.GetStateStub != nil {
		return fake.GetStateStub(namespace, key)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStateReturns.result1, fake.getStateReturns.result2
}

func (fake *QueryExecuter) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *QueryExecuter) GetStateArgsForCall(i int) (string, string) {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return fake.getStateArgsForCall[i].namespace, fake.getStateArgsForCall[i].key
}

func (fake *QueryExecuter) GetStateReturns(result1 *statedb.VersionedValue, result2 error) {
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 *statedb.VersionedValue
		result2 error
	}{result1, result2}
}

func (fake *QueryExecuter) GetStateReturnsOnCall(i int, result1 *statedb.VersionedValue, result2 error) {
	fake.GetStateStub = nil
	if fake.getStateReturnsOnCall == nil {
		fake.getStateReturnsOnCall = make(map[int]struct {
			result1 *statedb.VersionedValue
			result2 error
		})
	}
	fake.getStateReturnsOnCall[i] = struct {
		result1 *statedb.VersionedValue
		result2 error
	}{result1, result2}
}

func (fake *QueryExecuter) GetStateRangeScanIterator(namespace string, startKey string, endKey string) (statedb.ResultsIterator, error) {
	fake.getStateRangeScanIteratorMutex.Lock()
	ret, specificReturn := fake.getStateRangeScanIteratorReturnsOnCall[len(fake.getStateRangeScanIteratorArgsForCall)]
	fake.getStateRangeScanIteratorArgsForCall = append(fake.getStateRangeScanIteratorArgsForCall, struct {
		namespace string
		startKey  string
		endKey    string
	}{namespace, startKey, endKey})
	fake.recordInvocation("GetStateRangeScanIterator", []interface{}{namespace, startKey, endKey})
	fake.getStateRangeScanIteratorMutex.Unlock()
	if fake.GetStateRangeScanIteratorStub != nil {
		return fake.GetStateRangeScanIteratorStub(namespace, startKey, endKey)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStateRangeScanIteratorReturns.result1, fake.getStateRangeScanIteratorReturns.result2
}

func (fake *QueryExecuter) GetStateRangeScanIteratorCallCount() int {
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	return len(fake.getStateRangeScanIteratorArgsForCall)
}

func (fake *QueryExecuter) GetStateRangeScanIteratorArgsForCall(i int) (string, string, string) {
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	return fake.getStateRangeScanIteratorArgsForCall[i].namespace, fake.getStateRangeScanIteratorArgsForCall[i].startKey, fake.getStateRangeScanIteratorArgsForCall[i].endKey
}

func (fake *QueryExecuter) GetStateRangeScanIteratorReturns(result1 statedb.ResultsIterator, result2 error) {
	fake.GetStateRangeScanIteratorStub = nil
	fake.getStateRangeScanIteratorReturns = struct {
		result1 statedb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *QueryExecuter) GetStateRangeScanIteratorReturnsOnCall(i int, result1 statedb.ResultsIterator, result2 error) {
	fake.GetStateRangeScanIteratorStub = nil
	if fake.getStateRangeScanIteratorReturnsOnCall == nil {
		fake.getStateRangeScanIteratorReturnsOnCall = make(map[int]struct {
			result1 statedb.ResultsIterator
			result2 error
		})
	}
	fake.getStateRangeScanIteratorReturnsOnCall[i] = struct {
		result1 statedb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *QueryExecuter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *QueryExecuter) recordInvocation(key string, args []interface{}) {
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

var _ queryutil.QueryExecuter = new(QueryExecuter)