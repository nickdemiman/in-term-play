package intermplay

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nickdemiman/in-term-play/timer"
)

type (
	IScene interface {
		AddObject(GameObject)
		Object
		timer.TimerSender
	}

	Scene struct {
		GameObjects map[GameObject]bool
		Quitq       chan struct{}
		Bounds      Rect
		Style       tcell.Style
		IScene
	}
)

func (scene Scene) AddObject(obj GameObject) {
	_, ok := scene.GameObjects[obj]

	if !ok {
		scene.GameObjects[obj] = true
		obj.awake()
	}
}

func (scene Scene) NotifyTimer() {
	scene.update()
}

func (scene Scene) awake() {
	scene.Awake()

	<-scene.Quitq
}
func (scene Scene) update() {
	scene.Update()
}
func (scene Scene) dispose() {
	scene.Quitq <- struct{}{}

	for obj := range scene.GameObjects {
		obj.Dispose()
		delete(scene.GameObjects, obj)
	}

	timer.GetTimer().Unregister(scene)
	GetRenderer().Clear()

	scene.Dispose()
}

func (scene Scene) Awake()   {}
func (scene Scene) Update()  {}
func (scene Scene) Dispose() {}

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
