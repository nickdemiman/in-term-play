package internal

import (
	engine "github.com/nickdemiman/in-term-play"

	"github.com/gdamore/tcell/v2"
)

type (
	// Food interface {
	// 	GameObject
	// }

	Food struct {
		position engine.Vector2
		style    tcell.Style
		engine.GameObject
	}
)

func (f *Food) Position() engine.Vector2 {
	return f.position
}

func (f *Food) SetPosition(pos engine.Vector2) {
	f.position = pos
}

func (f *Food) Awake() {}

func (f *Food) Update() {
	engine.GetRenderer().SetContent(
		f.position.X,
		f.position.Y,
		' ',
		nil,
		f.style,
	)
}

func (f *Food) Dispose() {}

func (f *Food) Style() tcell.Style {
	return f.style
}

func (f *Food) HandleEvent(ev tcell.Event) {

}

func NewFood(pos engine.Vector2) *Food {
	f := new(Food)

	f.position = pos
	f.style = tcell.StyleDefault.Background(tcell.ColorYellow)

	return f
}
