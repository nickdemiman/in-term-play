package intermplay

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

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
		sync.Locker
	}

	GameObject struct {
		position      Vector2
		moveDirection Vector2
		velocity      float32
		sync.Mutex
	}
)

//lint:ignore U1000 скрытый метод
func (obj *GameObject) awake(i IGameObject) {
	game.Register(i)
	i.Awake()
}

//lint:ignore U1000 скрытый метод
func (obj *GameObject) update(i IGameObject) {
	i.Update()
}

//lint:ignore U1000 скрытый метод
func (obj *GameObject) dispose(i IGameObject) {
	game.Unregister(i)
	i.Dispose()
}

func (obj *GameObject) Awake()   {}
func (obj *GameObject) Update()  {}
func (obj *GameObject) Dispose() {}

func (obj *GameObject) Position() Vector2 { return obj.position }
func (obj *GameObject) SetPosition(vec Vector2) {
	defer obj.Unlock()
	obj.Lock()
	obj.position = vec
}

func (obj *GameObject) MoveDirection() Vector2 { return obj.moveDirection }
func (obj *GameObject) SetMoveDirection(vec Vector2) {
	defer obj.Unlock()
	obj.Lock()
	obj.moveDirection = vec
}

func (obj *GameObject) Velocity() float32 { return obj.velocity }
func (obj *GameObject) SetVelocity(velocity float32) {
	defer obj.Unlock()
	obj.Lock()
	obj.velocity = velocity
}

func (obj *GameObject) updatePhysics(gm IGameObject, dt float32) {
	defer gm.Unlock()
	gm.Lock()
	gm.UpdatePhysics(dt)
}

func (obj *GameObject) UpdatePhysics(dt float32)        {}
func (obj *GameObject) HandleTermEvents(ev tcell.Event) {}
