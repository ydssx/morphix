package service

import (
	"github.com/google/wire"
	"github.com/ydssx/morphix/app/job/internal/common"
)

var ProviderSet = wire.NewSet(NewJobService, common.NewAsynqClient)
