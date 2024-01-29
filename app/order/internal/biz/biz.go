package biz

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewOrderUseCase,
	common.NewProductClient,
	common.NewPaymentClient,
	common.NewQuoteClient,
	common.NewJobClient,
)
