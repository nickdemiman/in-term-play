package intermplay

import (
	"github.com/gdamore/tcell/v2"
)

type (
	IScene interface {
		updatePhysics(float32)
		awake(IScene)
		update(IScene)
		dispose(IScene)
		Awake()
		Update()
		Dispose()
		AddObject(IGameObject)
	}

	Scene struct {
		GameObjects map[IGameObject]bool
		Bounds      Rect
		Style       tcell.Style
		// IScene
	}
)

func (scene *Scene) AddObject(obj IGameObject) {
	_, ok := scene.GameObjects[obj]

	if !ok {
		scene.GameObjects[obj] = true
		obj.awake(obj)
	}
}

func (scene *Scene) NotifyTimer() {
	scene.update(scene)
}

func (scene *Scene) awake(s IScene) {
	s.Awake()
}

func (scene *Scene) update(s IScene) {
	s.Update()
}
func (scene *Scene) dispose(s IScene) {
	for obj := range scene.GameObjects {
		obj.Dispose()
		delete(scene.GameObjects, obj)
	}

	GetRenderer().Clear()

	s.Dispose()
}

func (scene *Scene) updatePhysics(dt float32) {
	for obj := range scene.GameObjects {
		v, ok := interface{}(obj).(IMoveable)
		if ok {
			v.UpdatePhysics(dt)
		}
	}
}

func (scene *Scene) Awake()   {}
func (scene *Scene) Update()  {}
func (scene *Scene) Dispose() {}
