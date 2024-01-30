package biz

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/redis"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewOrderUseCase,
	common.NewProductClient,
	common.NewPaymentClient,
	common.NewQuoteClient,
	common.NewJobClient,
	redis.NewLocker,
	wire.Bind(new(redis.Locker), new(*redis.RedisLocker)),
)
