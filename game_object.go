package intermplay

type (
	Object interface {
		awake(Object)
		update(Object)
		dispose(Object)
		Awake()
		Update()
		Dispose()
		TermEventsListener
	}

	IGameObject interface {
		Object
		Transform
	}

	GameObject struct {
		RootPosition Vector2
		// Style() tcell.Style
		IGameObject
		DefaultTermEventsListener
	}
)

func (obj *GameObject) awake(i Object) {
	globalEventer.Register(i)
	i.Awake()
}
func (obj *GameObject) update(i Object) {
	i.Update()
}
func (obj *GameObject) dispose(i Object) {
	globalEventer.Unregister(i)
	i.Dispose()
}

func (obj *GameObject) Awake()   {}
func (obj *GameObject) Update()  {}
func (obj *GameObject) Dispose() {}

// func (obj *GameObject) Position() Vector2 {
// 	return obj.RootPosition
// }
// func (obj *GameObject) SetPosition(vec Vector2) {}
