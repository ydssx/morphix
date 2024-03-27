package biz

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/lock"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewOrderUseCase,
	common.NewProductClient,
	common.NewPaymentClient,
	common.NewQuoteClient,
	common.NewJobClient,
	lock.NewLocker,
	wire.Bind(new(lock.Locker), new(*lock.RedisLocker)),
)
