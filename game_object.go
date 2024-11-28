package intermplay

import "github.com/gdamore/tcell/v2"

type (
	Object interface {
		awake()
		update()
		dispose()
		Awake()
		Update()
		Dispose()
	}

	GameObject interface {
		// Style() tcell.Style
		HandleEvent(tcell.Event)
		Object
		Transform
		TermEventsListener
	}

	gameObject struct {
		RootPosition Vector2
	}
	// GameObject struct {
	// }
)

func (obj *gameObject) handleTermEvents(tcell.Event) {}

func (obj *gameObject) awake() {
	globalEventer.Register(obj)
	obj.Awake()
}
func (obj *gameObject) update() {
	obj.Update()
}
func (obj *gameObject) dispose() {
	globalEventer.Unregister(obj)
	obj.Dispose()
}

func (obj *gameObject) Awake()   {}
func (obj *gameObject) Update()  {}
func (obj *gameObject) Dispose() {}

// func (obj *gameObject) Position() Vector2 {
// 	return obj.RootPosition
// }
// func (obj *gameObject) SetPosition(vec Vector2) {}
