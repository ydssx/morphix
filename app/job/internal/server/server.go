package server

import (
	"github.com/google/wire"
	ic "github.com/ydssx/morphix/app/job/internal/common"
	"github.com/ydssx/morphix/common"
	"github.com/ydssx/morphix/pkg/pubsub"
)

var ProviderSet = wire.NewSet(
	NewJobServer,
	NewGRPCServer,
	NewListenerServer,
	ic.NewServiceClientSet,
	common.NewNatsConn,
	common.NewCloudEvent,
	common.NewRedisClient,
	pubsub.NewRedisPubSub,
	wire.Bind(new(pubsub.Subscriber), new(*pubsub.RedisPubSub)),
)
