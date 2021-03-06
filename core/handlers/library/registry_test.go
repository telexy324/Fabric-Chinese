
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
/*
版权所有IBM Corp，SecureKey Technologies Inc.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package library

import (
	"testing"

	"github.com/hyperledger/fabric/core/handlers/auth"
	"github.com/hyperledger/fabric/core/handlers/decoration"
	"github.com/stretchr/testify/assert"
)

func TestInitRegistry(t *testing.T) {
	r := InitRegistry(Config{
		AuthFilters: []*HandlerConfig{{Name: "DefaultAuth"}},
		Decorators:  []*HandlerConfig{{Name: "DefaultDecorator"}},
	})
	assert.NotNil(t, r)
	authHandlers := r.Lookup(Auth)
	assert.NotNil(t, authHandlers)
	filters, isAuthFilters := authHandlers.([]auth.Filter)
	assert.True(t, isAuthFilters)
	assert.Len(t, filters, 1)

	decorationHandlers := r.Lookup(Decoration)
	assert.NotNil(t, decorationHandlers)
	decorators, isDecorators := decorationHandlers.([]decoration.Decorator)
	assert.True(t, isDecorators)
	assert.Len(t, decorators, 1)
}

func TestLoadCompiledInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic with invalid factory method")
		}
	}()

	testReg := registry{}
	testReg.loadCompiled("InvalidFactory", Auth)
}
