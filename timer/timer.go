package timer

import (
	"errors"
	"sync"
	"time"
)

type (
	Timer interface {
		Run()
		Stop()
		Register(TimerSender)
		Unregister(TimerSender) error
		searchSender(TimerSender) int
		// Notify(func())
	}

	TimerSender interface {
		NotifyTimer()
	}
	// TimerSender interface {
	// 	Subscribe()
	// 	Unsubscribe()
	// }

	timer struct {
		locker  sync.Mutex
		quitq   chan struct{}
		ticker  *time.Ticker
		senders []TimerSender
	}
)

var (
	_timer   *timer
	interval time.Duration
)

func GetTimer() Timer {
	if _timer == nil {
		_timer = new(timer)
		_timer.ticker = time.NewTicker(interval)

		return _timer
	}

	return _timer
}

func SetInterval(t time.Duration) {
	interval = t
}

func (t *timer) processEvents() {
	for _, sender := range t.senders {
		sender.NotifyTimer()
	}
}

func (t *timer) Run() {
loop:
	for {
		select {
		case <-t.ticker.C:
			t.processEvents()
		case <-t.quitq:
			break loop
		}
	}
}

func (t *timer) Stop() {
	t.quitq <- struct{}{}
}

func (t *timer) Register(sender TimerSender) {
	defer t.locker.Unlock()
	t.locker.Lock()

	t.senders = append(t.senders, sender)
}

func (t *timer) searchSender(sender TimerSender) int {
	for index, s := range t.senders {
		if s == sender {
			return index
		}
	}

	return -1
}

func (t *timer) Unregister(sender TimerSender) error {
	defer t.locker.Unlock()
	t.locker.Lock()

	index := t.searchSender(sender)

	if index == -1 {
		return errors.New("timer:Unregister; sender not in timer senders")
	}

	left := t.senders[:index]
	right := t.senders[index+1:]
	t.senders = append(left, right...)

	return nil
}
