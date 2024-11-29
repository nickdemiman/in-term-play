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
		baseHandleTermEvents(TermEventsListener, tcell.Event)
		handleTermEvents(tcell.Event)
	}

	DefaultTermEventsListener struct{}
)

func (d *DefaultTermEventsListener) baseHandleTermEvents(t TermEventsListener, ev tcell.Event) {
	t.handleTermEvents(ev)
}

var globalEventer TermEventsNotifier

func (t *TermEventsNotifier) Run() {
	for {
		event := GetRenderer().PollEvent()
		t.NotifyListeners(event)
	}
}

func (t *TermEventsNotifier) NotifyListeners(event tcell.Event) {
	for listener := range t.listeners {
		listener.baseHandleTermEvents(listener, event)
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
