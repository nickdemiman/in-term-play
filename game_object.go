package intermplay

type (
	IGameObject interface {
		awake(IGameObject)
		update(IGameObject)
		dispose(IGameObject)
		Awake()
		Update()
		Dispose()
		TermEventsListener
		Transform
	}

	GameObject struct{}
)

//lint:ignore U1000 перегрузка
func (obj *GameObject) awake(i IGameObject) {
	game.Register(i)
	i.Awake()
}

//lint:ignore U1000 перегрузка
func (obj *GameObject) update(i IGameObject) {
	i.Update()
}

//lint:ignore U1000 перегрузка
func (obj *GameObject) dispose(i IGameObject) {
	game.Unregister(i)
	i.Dispose()
}

func (obj *GameObject) Awake()   {}
func (obj *GameObject) Update()  {}
func (obj *GameObject) Dispose() {}

func (obj *GameObject) Position() Vector2 {
	return Vector2Null
}
func (obj *GameObject) SetPosition(vec Vector2) {}
