package service

import (
	"context"

	"github.com/hibiken/asynq"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JobService struct {
	cli *asynq.Client
	ipt *asynq.Inspector

	jobv1.UnimplementedJobServiceServer
}

func NewJobService(cli *asynq.Client, ipt *asynq.Inspector) *JobService {
	return &JobService{cli: cli, ipt: ipt}
}

func (j *JobService) Enqueue(ctx context.Context, req *jobv1.EnqueueRequest) (*jobv1.EnqueueResponse, error) {
	opts := []asynq.Option{asynq.MaxRetry(0)}
	if req.RetryTime > 0 {
		opts = append(opts, asynq.MaxRetry(int(req.RetryTime)))
	}
	if req.ProcessAt.IsValid() {
		opts = append(opts, asynq.ProcessAt(req.ProcessAt.AsTime()))
	}
	if req.ProcessIn.IsValid() {
		opts = append(opts, asynq.ProcessIn(req.ProcessIn.AsDuration()))
	}
	if req.Retention.IsValid() {
		opts = append(opts, asynq.Retention(req.Retention.AsDuration()))
	}
	taskInfo, err := j.cli.EnqueueContext(ctx, asynq.NewTask(req.JobType.String(), req.Payload), opts...)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "发送任务失败,err: %v", err)
	}

	return &jobv1.EnqueueResponse{TaskId: taskInfo.ID}, nil
}

func (j *JobService) QueryTasks(ctx context.Context, req *jobv1.QueryTasksRequest) (resp *jobv1.QueryTasksResponse, err error) {
	resp = new(jobv1.QueryTasksResponse)
	for _, taskId := range req.TaskIds {
		taskInfo, err := j.ipt.GetTaskInfo("default", taskId)
		if err != nil {
			return nil, err
		}
		resp.Tasks = append(resp.Tasks, &jobv1.QueryTasksResponse_TaskInfo{
			TaskId: taskId,
			Result: taskInfo.Result,
			Status: taskInfo.State.String(),
		})
	}

	return
}
