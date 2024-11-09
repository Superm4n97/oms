package queue

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"k8s.io/klog/v2"
	"os"
	"time"
)

var Client *asynq.Client

func checkClient() error {
	for i := 1; i <= 3; i++ {
		if Client == nil {
			klog.Errorf("no async client is found, probably order-processor server is not running: try %d", i)
			time.Sleep(10 * time.Second)
			continue
		}
		return nil
	}
	return fmt.Errorf("no async client is found, probably order-processor server is not running")
}

func PushToQueue(handler string, c interface{}) error {
	payload, err := json.Marshal(c)
	if err != nil {
		return err
	}
	task := asynq.NewTask(handler, payload)
	info, err := Client.Enqueue(task, asynq.MaxRetry(3), asynq.Retention(5*time.Minute), asynq.ProcessIn(10*time.Second))
	if err != nil {
		return err
	}
	klog.Infof("%s task enqueued...", info.Type)
	return nil
}

func getRedisOpts() *asynq.RedisClientOpt {
	pass, found := os.LookupEnv("REDIS_PASSWORD")
	if !found {
		pass = "redis"
	}
	url, found := os.LookupEnv("REDIS_URL")
	if !found {
		url = "localhost:6379"
	}
	username, found := os.LookupEnv("REDIS_USERNAME")
	if !found {
		username = "default"
	}
	return &asynq.RedisClientOpt{
		Addr:     url,
		Username: username,
		Password: pass,
	}
}

func SetRedisClient() *asynq.RedisClientOpt {
	redisOpt := getRedisOpts()
	Client = asynq.NewClient(redisOpt)
	return redisOpt
}
