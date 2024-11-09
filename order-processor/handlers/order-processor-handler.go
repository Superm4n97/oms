package handlers

import (
	"context"
	"github.com/Superm4n97/oms/order-processor/queue"
	"github.com/goccy/go-json"
	"github.com/hibiken/asynq"
	"k8s.io/klog/v2"
	"time"
)

const (
	OrderQueue        = "order-queue"
	NotificationQueue = "notification-queue"
)

type Order struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

func ProcessOrder(ctx context.Context, task *asynq.Task) error {
	var order Order
	if err := json.Unmarshal(task.Payload(), &order); err != nil {
		klog.Error(err.Error())
		return err
	}
	for i := 1; i <= 3; i++ {
		klog.Infof("processing order for user %s", order.Username)
		time.Sleep(10 * time.Second)
	}
	klog.Infof("[order processed for user %s]", order.Username)

	if err := queue.PushToQueue(NotificationQueue, order); err != nil {
		klog.Error(err.Error())
		return err
	}
	return nil
}

func SetNotification(ctx context.Context, task *asynq.Task) error {
	var order Order
	if err := json.Unmarshal(task.Payload(), &order); err != nil {
		klog.Error(err.Error())
		return err
	}
	klog.Infof("*******Notification*******\n order ready for user %s", order.Username)
	return nil
}
