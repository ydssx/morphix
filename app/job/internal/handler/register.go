package handler

import "github.com/hibiken/asynq"

func RegisterJobHandler(mux *asynq.ServeMux) {
	for k, v := range jobHandlerMap {
		mux.HandleFunc(k, v)
	}
}

func RegisterCronJob(sd *asynq.Scheduler) {
	for k, v := range cronJobMap {
		_, err := sd.Register(k, asynq.NewTask(v, nil))
		if err != nil {
			panic(err)
		}
	}
}
