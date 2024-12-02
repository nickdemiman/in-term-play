package intermplay

import (
	"github.com/gdamore/tcell/v2"
)

func DispatchEvent(ev tcell.Event) {
	switch ev.(type) {
	// case *RenderEvent:
	// 	ev.scene.update(ev.scene)
	case *GameOverEvent:
		select {
		case <-game.quitc:
			return
		default:
			go game.Close()
		}
	}
}
