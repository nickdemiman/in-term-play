package internal

import (
	engine "github.com/nickdemiman/in-term-play"

	"github.com/gdamore/tcell/v2"
)

type (
	Cell struct {
		symbol rune
		style  tcell.Style
		engine.GameObject
	}
)

func NewCell(pos engine.Vector2, velocity float32, moveDirection engine.Vector2, style tcell.Style) *Cell {
	ce := new(Cell)

	ce.SetPosition(pos)
	ce.symbol = ' '
	ce.style = style
	ce.SetVelocity(velocity)
	ce.SetMoveDirection(moveDirection)

	return ce
}

func (ce *Cell) Symbol() rune          { return ce.symbol }
func (ce *Cell) SetSymbol(symbol rune) { ce.symbol = symbol }

func (ce *Cell) UpdatePhysics(dt float32) {
	vec := engine.Vector2{
		X: ce.MoveDirection().X * ce.Velocity() * dt,
		Y: ce.MoveDirection().Y * ce.Velocity() * dt,
	}

	ce.SetPosition(
		*engine.Vector2Add(ce.Position(), vec),
	)
}

func (ce *Cell) Update() {
	engine.GetRenderer().SetContent(
		int(ce.Position().X),
		int(ce.Position().Y),
		ce.symbol,
		nil,
		ce.style,
	)
}

func (ce *Cell) HandleTermEvents(ev tcell.Event) {}
