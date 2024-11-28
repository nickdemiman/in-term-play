package internal

import (
	engine "github.com/nickdemiman/in-term-play"

	"github.com/gdamore/tcell/v2"
)

type (
	// Cell interface {
	// 	Symbol() rune
	// 	SetSymbol(rune)
	// 	GameObject
	// }

	Cell struct {
		position engine.Vector2
		symbol   rune
		style    tcell.Style
		engine.GameObject
	}
)

func NewCell(pos engine.Vector2, symbol rune, style tcell.Style) *Cell {
	ce := new(Cell)

	ce.position = pos
	ce.symbol = symbol
	ce.style = style

	return ce
}

func (ce *Cell) Position() engine.Vector2 {
	return ce.position
}

func (ce *Cell) SetPosition(pos engine.Vector2) {
	ce.position = pos
}

func (ce *Cell) Style() tcell.Style {
	return ce.style
}

func (ce *Cell) Symbol() rune {
	return ce.symbol
}

func (ce *Cell) SetSymbol(symbol rune) {
	ce.symbol = symbol
}

func (ce *Cell) HandleEvent(ev tcell.Event) {

}

func (ce *Cell) Awake() {}

func (ce *Cell) Update() {
	engine.GetRenderer().SetContent(
		ce.position.X,
		ce.position.Y,
		ce.symbol,
		nil,
		ce.style,
	)
}

func (ce *Cell) Dispose() {}

func (ce *Cell) Collides(x, y int) bool {
	return false
}
