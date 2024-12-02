package intermplay

import "time"

type (
	GameEvent struct {
		when time.Time
	}

	// RenderEvent struct {
	// 	scene IScene
	// 	when  time.Time
	// }

	GameOverEvent struct {
		GameEvent
	}

	GamePauseEvent struct {
		GameEvent
	}
)

func (e *GameEvent) When() time.Time {
	return e.when
}

// func (e *RenderEvent) When() time.Time {
// 	return e.when
// }
