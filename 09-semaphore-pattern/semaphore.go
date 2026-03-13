package semaphore

import (
	"errors"
	"time"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not acquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release semaphore without acquiring it first")
)

type Interface interface {
	Acquire() error
	Release() error
}

// 定义结构体， 信号量使用chan struct{}
type Semaphore struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *Semaphore) Release() error {
	select {
	case <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &Semaphore{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}
