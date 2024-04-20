package biz

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/pkg/cache"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewSmsUseCase, cache.NewRedisCache)
