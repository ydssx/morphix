package handler

import (
	"fmt"

	"github.com/hibiken/asynq"
)

func RegisterJobHandler(mux *asynq.ServeMux) {
	for k, v := range jobHandlerMap {
		mux.HandleFunc(k, v)
	}
}

func RegisterCronJob(sd *asynq.Scheduler) {
	for k, v := range cronJobMap {
		if _, ok := jobHandlerMap[v]; !ok {
			panic(fmt.Sprintf("the cron job [%s] have not any registered handlers.", v))
		}
		_, err := sd.Register(k, asynq.NewTask(v, nil))
		if err != nil {
			panic(err)
		}
	}
}
