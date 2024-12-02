package intermplay

import (
	"github.com/gdamore/tcell/v2"
)

type (
	IScene interface {
		updatePhysics(IScene)
		UpdatePhysics(float32)
		awake(IScene)
		update(IScene, float32)
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

func (scene *Scene) awake(s IScene) {
	s.Awake()
}

func (scene *Scene) update(s IScene, alpha float32) {
	for obj := range scene.GameObjects {
		obj.update(obj, alpha)
	}

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

func (scene *Scene) updatePhysics(s IScene) {
	for obj := range scene.GameObjects {
		obj.updatePhysics(obj, dt)
	}

	s.UpdatePhysics(dt)
}

func (scene *Scene) Awake()   {}
func (scene *Scene) Update()  {}
func (scene *Scene) Dispose() {}
