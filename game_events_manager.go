package intermplay

import (
	"github.com/gdamore/tcell/v2"
)

func DispatchEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *RenderEvent:
		ev.scene.update(ev.scene)
	case *GameOverEvent:
		go game.Close()
	}
}
