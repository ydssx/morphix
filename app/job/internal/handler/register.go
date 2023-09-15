package handler

import (
	"fmt"

	"github.com/hibiken/asynq"
)

func RegisterJobHandler(mux *asynq.ServeMux) {
	for k, v := range jobHandlerMap {
		mux.HandleFunc(k.String(), v)
	}
}

func RegisterCronJob(sd *asynq.Scheduler) {
	for k, jobType := range cronJobMap {
		if _, ok := jobHandlerMap[jobType]; !ok {
			panic(fmt.Sprintf("the cron job [%s] have not any registered handlers.", jobType.String()))
		}
		_, err := sd.Register(k, asynq.NewTask(jobType.String(), nil))
		if err != nil {
			panic(err)
		}
	}
}
