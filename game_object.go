package intermplay

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

type (
	IGameObject interface {
		awake(IGameObject)
		update(IGameObject, float32)
		dispose(IGameObject)
		Awake()
		Update()
		Dispose()
		TermEventsListener
		Transform
		sync.Locker
	}

	GameObject struct {
		position        Vector2
		posBeforeRender Vector2
		moveDirection   Vector2
		velocity        float32
		sync.Mutex
	}
)

//lint:ignore U1000 перегрузка
func (obj *GameObject) awake(i IGameObject) {
	game.Register(i)
	i.Awake()
	// i.setPositionBeforeRender(i.Position())
}

//lint:ignore U1000 перегрузка
func (obj *GameObject) update(i IGameObject, alpha float32) {
	// obj.interpoletePhysics(i, alpha)
	i.Update()
	// i.setPositionBeforeRender(obj.position)
}

//lint:ignore U1000 перегрузка
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

//	func (obj *GameObject) updatePhysics(dt float32) {
//		obj.velocity =
//	}
func (obj *GameObject) positionBeforeRender() Vector2 {
	return obj.posBeforeRender
}

func (obj *GameObject) setPositionBeforeRender(dst Vector2) {
	obj.posBeforeRender = dst
}

func (obj *GameObject) interpoletePhysics(gm IGameObject, alpha float32) {
	x0, y0 := gm.positionBeforeRender().XY()
	x1, y1 := gm.Position().XY()

	if x0 == x1 && y0 == y1 {
		return
	}

	timeBefore := float32(timeBeforeRender.Nanosecond())
	timeAfter := float32(timeAfterRender.Nanosecond())

	interpolatedX := x0 + (x1-x0)/(timeAfter-timeBefore)*(alpha-timeBefore)
	interpolatedY := y0 + (y1-y0)/(timeAfter-timeBefore)*(alpha-timeBefore)

	gm.SetPosition(Vector2{interpolatedX, interpolatedY})
	// res := y0 + (y1-y0)/(x1-x0)*(alpha-x0)

	// if math.IsNaN(float64(res)) {
	// 	gm.SetPosition(
	// 		Vector2{X: x1, Y: 0},
	// 	)
	// } else {
	// 	gm.SetPosition(
	// 		Vector2{X: x1, Y: res},
	// 	)
	// }
}

func (obj *GameObject) updatePhysics(gm IGameObject, dt float32) {
	defer gm.Unlock()
	gm.Lock()
	gm.UpdatePhysics(dt)
}

func (obj *GameObject) UpdatePhysics(dt float32)        {}
func (obj *GameObject) HandleTermEvents(ev tcell.Event) {}

func NewGameObject() IGameObject {
	return new(GameObject)
}
