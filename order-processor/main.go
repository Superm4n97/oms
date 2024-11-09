package main

import (
	"github.com/Superm4n97/oms/order-processor/handlers"
	"github.com/Superm4n97/oms/order-processor/queue"
	"github.com/hibiken/asynq"
	"k8s.io/klog/v2"
)

func main() {
	redisOpt := queue.SetRedisClient()

	jobserver := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 10,
			LogLevel:    asynq.InfoLevel,
		})

	mux := asynq.NewServeMux()

	registerJobHandlers(mux)

	klog.Infof("starting job processor server....")
	if err := jobserver.Run(mux); err != nil {
		klog.Error(err.Error())
		return
	}
}

func registerJobHandlers(w *asynq.ServeMux) {
	w.HandleFunc(handlers.OrderQueue, handlers.ProcessOrder)
	w.HandleFunc(handlers.NotificationQueue, handlers.SetNotification)
}
