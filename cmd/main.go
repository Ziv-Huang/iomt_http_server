package main

import (
	"fmt"
	"iomt_http_server/domain/configs"
	log "iomt_http_server/domain/logger"
	"iomt_http_server/domain/utils"
	"iomt_http_server/usecase/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var config configs.Configuration
	configs.ReadConfig(&config)
	log.SetLevel(config.Project.Log)
	log.Info(fmt.Sprintf("project: %s  version: %s", config.Project.Name, config.Project.Version))
	utils.ListenExit()

	// Launch server
	wg.Add(1)
	log.Info("Launching server...")
	server.NewServer(&wg, config).LaunchServer()
	wg.Wait()

	log.Info("Server Finished!")
}
