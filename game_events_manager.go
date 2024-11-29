package intermplay

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type (
	gameEventManager struct {
	}

	RenderSceneEvent struct {
		when  time.Time
		scene IScene
	}
)

func (ev *RenderSceneEvent) When() time.Time {
	return ev.when
}

// func (gm *gameEventManager) Register()   {}
// func (gm *gameEventManager) Unregister() {}
func (gm *gameEventManager) DispatchEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *RenderSceneEvent:
		ev.scene.update(ev.scene)
	}
}
