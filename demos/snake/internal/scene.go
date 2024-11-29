package internal

import (
	"math/rand/v2"

	"github.com/gdamore/tcell/v2"
	engine "github.com/nickdemiman/in-term-play"
)

type MainScene struct {
	// engine.IScene

	engine.Scene
}

var _player *Player

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func (scene *MainScene) generateFood() {
	x := scene.Bounds.Origin().X
	y := scene.Bounds.Origin().Y
	w, h := scene.Bounds.Size()

	_x := randRange(x, x+w-1)
	_y := randRange(y, y+h-1)

	food := NewFood(engine.NewVector2(_x, _y))

	scene.AddObject(food)
}

func (scene *MainScene) drawBorders() {
	x := scene.Bounds.Origin().X
	y := scene.Bounds.Origin().Y
	w, h := scene.Bounds.Size()

	engine.GetRenderer().SetContent(x, y, '┌', nil, scene.Style)
	engine.GetRenderer().SetContent(x, y+h, '└', nil, scene.Style)
	engine.GetRenderer().SetContent(x+w, y, '┐', nil, scene.Style)
	engine.GetRenderer().SetContent(x+w, y+h, '┘', nil, scene.Style)

	for i := y + 1; i < h; i++ {
		engine.GetRenderer().SetContent(x, i, '│', nil, scene.Style)
	}
	for i := y + 1; i < h; i++ {
		engine.GetRenderer().SetContent(x+w, i, '│', nil, scene.Style)
	}
	for i := x + 1; i < w; i++ {
		engine.GetRenderer().SetContent(i, y, '─', nil, scene.Style)
	}
	for i := x + 1; i < w; i++ {
		engine.GetRenderer().SetContent(i, y+h, '─', nil, scene.Style)
	}
}

func (scene *MainScene) Awake() {
	_player = NewPlayer(engine.NewVector2(10, 10), 5)
	food := NewFood(engine.NewVector2(5, 5))
	scene.AddObject(_player)
	scene.AddObject(food)
}
func (scene *MainScene) Update() {
	engine.GetRenderer().Clear()
	scene.drawBorders()

	pos := _player.Position()
	pos.Add(_player.MoveDirection())

loop:
	for obj := range scene.GameObjects {
		switch obj.(type) {
		case *Food:
			if pos.IsEqual(obj.Position()) {
				_player.Move()
				_player.Eat()
				delete(scene.GameObjects, obj)
				scene.generateFood()

				break loop
			}
		}
	}

	_player.Move()

	for obj := range scene.GameObjects {
		scene.checkBounds(obj)
		obj.Update()
	}

	engine.GetRenderer().Sync()
}
func (scene *MainScene) Dispose() {}

func (scene *MainScene) checkBounds(obj engine.IGameObject) {
	pos := obj.Position()
	x := scene.Bounds.Origin().X
	y := scene.Bounds.Origin().Y
	w, h := scene.Bounds.Size()

	if pos.X < x+1 {
		pos.X = x + w - 1
		obj.SetPosition(pos)

		return
	}

	if pos.X > w-1 {
		pos.X = x + 1
		obj.SetPosition(pos)

		return
	}

	if pos.Y < y+1 {
		pos.Y = y + h - 1
		obj.SetPosition(pos)

		return
	}

	if pos.Y > h-1 {
		pos.Y = y + 1
		obj.SetPosition(pos)

		return
	}

}

func NewMainScene(x, y, width, height int) engine.IScene {
	scene := new(MainScene)

	style := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack)

	scene.GameObjects = make(map[engine.IGameObject]bool)
	scene.Bounds = engine.NewRect(x, y, width, height)
	scene.Style = style

	return scene
}
