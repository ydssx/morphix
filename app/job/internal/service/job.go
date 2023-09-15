package service

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JobService struct {
	cli *asynq.Client

	jobv1.UnimplementedJobServiceServer
}

func NewJobService(cli *asynq.Client) *JobService {
	return &JobService{cli: cli}
}

func (j *JobService) Enqueue(ctx context.Context, req *jobv1.EnqueueRequest) (*emptypb.Empty, error) {
	opts := []asynq.Option{asynq.MaxRetry(0), asynq.Retention(time.Hour)}
	if req.RetryTime > 0 {
		opts = append(opts, asynq.MaxRetry(int(req.RetryTime)))
	}
	if req.ProcessAt.IsValid() {
		opts = append(opts, asynq.ProcessAt(req.ProcessAt.AsTime()))
	}
	if req.ProcessIn.IsValid() {
		opts = append(opts, asynq.ProcessIn(req.ProcessIn.AsDuration()))
	}
	_, err := j.cli.EnqueueContext(ctx, asynq.NewTask(req.JobType.String(), req.Payload), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "发送任务失败,err: %v", err)
	}

	return &emptypb.Empty{}, nil
}
