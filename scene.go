package intermplay

import (
	"github.com/gdamore/tcell/v2"
)

type (
	IScene interface {
		updatePhysics(IScene, float32)
		UpdatePhysics(float32)
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

//lint:ignore U1000 скрытый метод
func (scene *Scene) awake(s IScene) {
	s.Awake()
}

//lint:ignore U1000 скрытый метод
func (scene *Scene) update(s IScene) {
	for obj := range scene.GameObjects {
		obj.update(obj)
	}

	s.Update()
}

//lint:ignore U1000 скрытый метод
func (scene *Scene) dispose(s IScene) {
	for obj := range scene.GameObjects {
		obj.Dispose()
		delete(scene.GameObjects, obj)
	}

	GetRenderer().Clear()

	s.Dispose()
}

//lint:ignore U1000 скрытый метод
func (scene *Scene) updatePhysics(s IScene, dt float32) {
	for obj := range scene.GameObjects {
		obj.updatePhysics(obj, dt)
	}

	s.UpdatePhysics(dt)
}

func (scene *Scene) Awake()   {}
func (scene *Scene) Update()  {}
func (scene *Scene) Dispose() {}
