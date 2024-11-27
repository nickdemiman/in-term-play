package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nickdemiman/in-term-play/screen"
	"github.com/nickdemiman/in-term-play/timer"
)

type (
	Scene interface {
		AddObject(GameObject) Scene
		Object
		timer.TimerSender
	}

	scene struct {
		gameObjects   map[GameObject]bool
		quitq         chan struct{}
		gameEventChan *chan GameEvent
		bounds        Rect
		style         tcell.Style
	}
)

func NewScene(x, y, width, height int, gameEventChan *chan GameEvent) (Scene, error) {
	scene := new(scene)
	scene.quitq = make(chan struct{})
	scene.gameObjects = make(map[GameObject]bool)
	scene.gameEventChan = gameEventChan
	scene.bounds = NewRect(x, y, width, height)
	scene.style = tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack)

	return scene, nil
}

func (scene *scene) AddObject(obj GameObject) Scene {
	_, ok := scene.gameObjects[obj]

	if !ok {
		scene.gameObjects[obj] = true
		obj.Awake()
	}

	return scene
}

func (scene *scene) NotifyTimer() {
	scene.Update()
}

func (scene *scene) Awake() {

	<-scene.quitq

	for obj := range scene.gameObjects {
		obj.Dispose()
		delete(scene.gameObjects, obj)
	}
	timer.GetTimer().Unregister(scene)
	screen.GetGameScreen().Screen.Clear()
}

func (scene *scene) Update() {}

func (scene *scene) Dispose() {
	scene.quitq <- struct{}{}
}
