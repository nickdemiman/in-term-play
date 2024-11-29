package intermplay

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nickdemiman/in-term-play/timer"
)

type (
	IScene interface {
		AddObject(IGameObject)
		awake(IScene)
		update(IScene)
		dispose(IScene)
		Awake()
		Update()
		Dispose()
		timer.TimerSender
	}

	Scene struct {
		GameObjects map[IGameObject]bool
		Bounds      Rect
		Style       tcell.Style
		IScene
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

func (scene *Scene) Awake()   {}
func (scene *Scene) Update()  {}
func (scene *Scene) Dispose() {}

// func NewScene(x, y, width, height int) IScene {
// 	scene := new(S)

// 	style := tcell.StyleDefault.
// 		Foreground(tcell.ColorWhite).
// 		Background(tcell.ColorBlack)

// 	scene.quitq = make(chan struct{})
// 	scene.GameObjects = make(map[GameObject]bool)
// 	scene.Bounds = NewRect(x, y, width, height)
// 	scene.Style = style

// 	return scene
// }
