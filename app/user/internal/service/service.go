package service

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService, common.NewSMSClient)
