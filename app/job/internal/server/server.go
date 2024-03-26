package server

import (
	"github.com/google/wire"
	ic "github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/common"
)

var ProviderSet = wire.NewSet(
	NewJobServer,
	NewGRPCServer,
	NewListenerServer,
	ic.NewServiceClientSet,
	common.NewNatsConn,
	common.NewCloudEvent,
)
