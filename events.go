package intermplay

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	GAMEOVER  = iota
	GAMEPAUSE = iota
)

type (
	BaseEvent struct {
		when time.Time
		tcell.Event
	}

	GameEvent struct {
		Type uint
		BaseEvent
	}

	EventHandler interface {
		HandleEvent(tcell.Event)
	}
)

func (ev *BaseEvent) When() time.Time {
	return ev.when
}

var (
	GameOver = GameEvent{
		Type: GAMEOVER,
	}

	GamePause = GameEvent{
		Type: GAMEPAUSE,
	}
)
