package biz

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase, common.NewSMSClient)
