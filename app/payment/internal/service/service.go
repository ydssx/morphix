package service

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/common/dapr"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPaymentService, NewEventSender, dapr.NewDaprClient, common.NewCloudEvent, common.NewNatsConn)
