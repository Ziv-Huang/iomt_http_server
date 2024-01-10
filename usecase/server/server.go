package server

import (
	"iomt_http_server/delivery"
	"iomt_http_server/domain/configs"
	"sync"
)

type Server struct {
	Delivery delivery.IDelivery
}

func NewServer(wg *sync.WaitGroup, config configs.Configuration) *Server {
	return &Server{Delivery: delivery.NewDelivery(wg, config)}
}

func (s *Server) LaunchServer() {
	s.Delivery.Start()
}
