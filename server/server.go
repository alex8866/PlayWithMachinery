package server

import (
	"github.com/RichardKnop/machinery/v1"

	"github.com/zhiruchen/PlayWithMachinery/config"
)

type ITask interface {
	UpdateField(string) error
	SendMsg(string) error
}

// StartServer start server register tasks
func StartServer(t ITask) (server *machinery.Server, err error) {
	// Create server instance
	if config.Config == nil {
		config.LoadConfig("/Users/zhiruchen/go/src/github.com/zhiruchen/PlayWithMachinery/config/config.yaml")
	}
	server, err = machinery.NewServer(config.Config)
	if err != nil {
		return
	}

	tasks := map[string]interface{}{
		"updatefield": t.UpdateField,
		"sendmsg":     t.SendMsg,
	}

	err = server.RegisterTasks(tasks)
	return
}

// LaunchWorker 开启一个worker
func LaunchWorker(workerTag string, t ITask) error {
	s, err := StartServer(t)
	if err != nil {
		return err
	}

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := s.NewWorker(workerTag)
	if err := worker.Launch(); err != nil {
		return err
	}

	return nil
}
