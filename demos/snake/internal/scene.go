package internal

import (
	"fmt"
	"math/rand/v2"

	"github.com/gdamore/tcell/v2"
	engine "github.com/nickdemiman/in-term-play"
)

type MainScene struct {
	engine.Scene
}

var _player *Player
var _food *Food

func randRangeFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func (scene *MainScene) drawPoints() {
	x := scene.Bounds.Origin().X
	y := scene.Bounds.Origin().Y
	_, h := scene.Bounds.Size()

	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	text := fmt.Sprintf("Points: %d", _player.points)

	for i, ch := range text {
		engine.GetRenderer().SetContent(int(x)+i, int(y+h)+1, ch, nil, style)
	}
}

func (scene *MainScene) generateFood() {
	x := scene.Bounds.Origin().X
	y := scene.Bounds.Origin().Y
	w, h := scene.Bounds.Size()

	_food.SetPosition(engine.Vector2{
		X: randRangeFloat32(x, x+w-1),
		Y: randRangeFloat32(y, y+h-1),
	})
}

func (scene *MainScene) drawBorders() {
	x := int(scene.Bounds.Origin().X)
	y := int(scene.Bounds.Origin().Y)
	w := int(scene.Bounds.W())
	h := int(scene.Bounds.H())

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
func (scene *MainScene) checkBounds(obj engine.IGameObject) {
	pos := obj.Position()
	x := scene.Bounds.Origin().X
	y := scene.Bounds.Origin().Y
	w, h := scene.Bounds.Size()

	if pos.X < x+1.0 {
		pos.X = x + w - 1.0
		obj.SetPosition(pos)

		return
	}

	if pos.X > w {
		pos.X = x + 1.0
		obj.SetPosition(pos)

		return
	}

	if pos.Y < y+1.0 {
		pos.Y = y + h - 1.0
		obj.SetPosition(pos)

		return
	}

	if pos.Y > h {
		pos.Y = y + 1.0
		obj.SetPosition(pos)

		return
	}
}

func (scene *MainScene) Awake() {
	_player = NewPlayer(engine.NewVector2(10, 10), 5)
	_food = NewFood(engine.NewVector2(5, 5))
	scene.AddObject(_player)
	scene.AddObject(_food)

	engine.GetRenderer().Sync()
}
func (scene *MainScene) Update() {
	engine.GetRenderer().Clear()
	scene.drawBorders()
	scene.drawPoints()

	pos := _player.Position()

	if int(pos.X) == int(_food.position.X) && int(pos.Y) == int(_food.position.Y) {
		_player.Eat()
		scene.generateFood()
	}

	if _player.checkSelfCollide(pos.XY()) {
		// go engine.GetGame().Close()
		engine.DispatchEvent(&engine.GameOverEvent{})
		return
	}

	for obj := range scene.GameObjects {
		scene.checkBounds(obj)
		obj.Update()
	}

	engine.GetRenderer().Show()
}
func (scene *MainScene) Dispose() {}

func NewMainScene(x, y, width, height int) engine.IScene {
	scene := new(MainScene)

	style := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack)

	scene.GameObjects = make(map[engine.IGameObject]bool)
	scene.Bounds = engine.NewRect(float32(x), float32(y), float32(width), float32(height))
	scene.Style = style

	return scene
}
