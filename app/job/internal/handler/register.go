package handler

import (
	"fmt"

	"github.com/hibiken/asynq"
	jobv1 "github.com/ydssx/morphix/api/job/v1"
)

// RegisterJobHandler 注册 jobHandlerMap 中定义的所有 job 的处理函数到 ServeMux。
// 它会遍历 jobHandlerMap,并为每个 job 注册对应的处理函数到 mux。
// mux 会根据请求中的 job name 来路由到相应的处理函数。
func RegisterJobHandler(mux *asynq.ServeMux) {
	for k, v := range jobHandlerMap {
		mux.HandleFunc(k.String(), v)
	}
}

// RegisterCronJob 注册定时任务的处理函数。
// 它会遍历 cronJobMap 中定义的所有定时任务,并在调度器 sd 中注册对应的处理函数。
// 如果某个定时任务在 jobHandlerMap 中没有找到对应的处理函数,会 panic。
// 注册成功后,定时任务会按照 cronJobMap 中定义的时间表定期执行。
func RegisterCronJob(sd *asynq.Scheduler) {
	for k, jobType := range cronJobMap {
		err := ValidateTask(jobType)
		if err != nil {
			panic(err)
		}
		_, err = sd.Register(k, asynq.NewTask(jobType.String(), nil))
		if err != nil {
			panic(err)
		}
	}
}

// ValidateTask validates that the given jobType has a registered handler.
// It checks if there is a handler function registered in jobHandlerMap for
// the given jobType. If not, it returns an error.
func ValidateTask(jobType jobv1.JobType) error {
	if _, ok := jobHandlerMap[jobType]; !ok {
		return fmt.Errorf("the cron job [%s] does not have any registered handlers", jobType.String())
	}
	return nil
}
