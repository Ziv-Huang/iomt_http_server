package delivery

import (
	"iomt_http_server/domain/configs"
	"sync"
)

type IDelivery interface {
	Start()
}

func NewDelivery(wg *sync.WaitGroup, config configs.Configuration) IDelivery {
	return NewHttpFiber(wg, config)
}
