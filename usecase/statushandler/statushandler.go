package statushandler

import (
	"sync"
	"time"
)

// StatusHandler ...
type StatusHandler interface {
	GetStatus() bool
	Countdown()
	// SetStatus(bool)
}

// NewStatusHandler ...
func NewStatusHandler(waittime int) StatusHandler {
	return &statusHandler{status: false, waittime: waittime}
}

type statusHandler struct {
	status   bool
	waittime int
	lock     sync.Mutex
}

func (s *statusHandler) GetStatus() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.status
}

func (s *statusHandler) Countdown() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.status {
		return
	}

	s.status = true
	go func() {
		defer func() {
			s.lock.Lock()
			s.status = false
			s.lock.Unlock()
		}()
		<-time.After(time.Duration(s.waittime) * time.Second)
	}()
}
