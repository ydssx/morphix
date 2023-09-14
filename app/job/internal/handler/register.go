package handler

import (
	"fmt"

	"github.com/hibiken/asynq"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
)

func RegisterJobHandler(mux *asynq.ServeMux) {
	for k, v := range jobHandlerMap {
		mux.HandleFunc(jobv1.JobType_name[int32(k)], v)
	}
}

func RegisterCronJob(sd *asynq.Scheduler) {
	for k, jobType := range cronJobMap {
		if _, ok := jobHandlerMap[jobType]; !ok {
			panic(fmt.Sprintf("the cron job [%s] have not any registered handlers.", jobType.String()))
		}
		_, err := sd.Register(k, asynq.NewTask(jobv1.JobType_name[int32(jobType)], nil))
		if err != nil {
			panic(err)
		}
	}
}
