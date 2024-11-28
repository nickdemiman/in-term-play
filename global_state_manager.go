package intermplay

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

type (
	TermEventsNotifier struct {
		listeners map[TermEventsListener]bool
		sync.Mutex
	}
	TermEventsListener interface {
		handleTermEvents(tcell.Event)
	}
)

var globalEventer TermEventsNotifier

func (t *TermEventsNotifier) Run() {
	for {
		event := GetRenderer().PollEvent()
		t.Notify(event)
	}
}

func (t *TermEventsNotifier) Notify(event tcell.Event) {
	for listener := range t.listeners {
		listener.handleTermEvents(event)
	}
}

func (t *TermEventsNotifier) Register(listener TermEventsListener) {
	defer t.Unlock()
	t.Lock()
	t.listeners[listener] = true
}

func (t *TermEventsNotifier) Unregister(listener TermEventsListener) {
	defer t.Unlock()
	t.Lock()
	delete(t.listeners, listener)
}
