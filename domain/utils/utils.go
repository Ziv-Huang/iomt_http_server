package utils

import (
	log "iomt_http_server/domain/logger"
	"os"
	"os/signal"
	"syscall"
)

func ListenExit() {
	c := make(chan os.Signal)
	//Listen to signals ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Info("Program Exit...", s)
				os.Exit(0)
			case syscall.SIGUSR1:
				log.Info("usr1 signal", s)
			case syscall.SIGUSR2:
				log.Info("usr2 signal", s)
			default:
				log.Info("other signal", s)
			}
		}
	}()
}
