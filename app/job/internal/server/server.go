package server

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/job/internal/common"
)

var ProviderSet = wire.NewSet(NewJobServer, NewGRPCServer, common.NewServiceClientSet)
