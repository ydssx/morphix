package server

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/limit"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, common.MustNewRedisClient, limit.NewRedisLimiter)
