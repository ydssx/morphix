package service

import (
	"context"

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
	opts := []asynq.Option{asynq.MaxRetry(1)}
	_, err := j.cli.EnqueueContext(ctx, asynq.NewTask(jobv1.JobType_name[int32(req.JobType)], req.Payload), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "发送任务失败,err: %v", err)
	}

	return &emptypb.Empty{}, nil
}
