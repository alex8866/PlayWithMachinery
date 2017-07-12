package main

import (
	"flag"
	"log"


	"github.com/zhiruchen/PlayWithMachinery/client"
	"github.com/zhiruchen/PlayWithMachinery/config"
	"github.com/zhiruchen/PlayWithMachinery/server"
)

func main() {

	cfg := flag.String("config", "config/config.yaml", "config file")
	flag.Parse()
	config.LoadConfig(*cfg)

	workerTag := "delay_worker_1"
	if err := server.LaunchWorker(workerTag, new(client.DelayTask)); err != nil {
		log.Fatalf("start worker error: %v", err)
	}
}
