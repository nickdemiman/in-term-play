package internal

import (
	engine "github.com/nickdemiman/in-term-play"

	"github.com/gdamore/tcell/v2"
)

type (
	Cell struct {
		position      engine.Vector2
		symbol        rune
		style         tcell.Style
		velocity      float32
		moveDirection engine.Vector2
		engine.GameObject
	}
)

func NewCell(pos engine.Vector2, velocity float32, moveDirection engine.Vector2, style tcell.Style) *Cell {
	ce := new(Cell)

	ce.position = pos
	ce.symbol = ' '
	ce.style = style
	ce.velocity = velocity
	ce.moveDirection = moveDirection

	return ce
}

func (ce *Cell) Velocity() float32                   { return ce.velocity }
func (ce *Cell) SetVelocity(velocity float32)        { ce.velocity = velocity }
func (ce *Cell) MoveDirection() engine.Vector2       { return ce.moveDirection }
func (ce *Cell) SetMoveDirection(dst engine.Vector2) { ce.moveDirection = dst }
func (ce *Cell) Symbol() rune                        { return ce.symbol }
func (ce *Cell) SetSymbol(symbol rune)               { ce.symbol = symbol }
func (ce *Cell) Position() engine.Vector2            { return ce.position }
func (ce *Cell) SetPosition(vec engine.Vector2)      { ce.position = vec }

func (ce *Cell) UpdatePhysics(dt float32) {
	vec := engine.Vector2{
		X: ce.moveDirection.X * ce.velocity * dt,
		Y: ce.moveDirection.Y * ce.velocity * dt,
	}

	ce.position.Add(vec)
}

func (ce *Cell) Update() {
	engine.GetRenderer().SetContent(
		int(ce.position.X),
		int(ce.position.Y),
		ce.symbol,
		nil,
		ce.style,
	)
}

func (ce *Cell) HandleTermEvents(ev tcell.Event) {}
