package internal

import (
	"github.com/gdamore/tcell/v2"
	engine "github.com/nickdemiman/in-term-play"
)

type (
	Player struct {
		points    uint
		body      []*Cell
		head      *Cell
		styleHead tcell.Style
		styleBody tcell.Style
		engine.GameObject
	}
)

const playerVelocity float32 = 20
const playerLen int = 5

func NewPlayer(position engine.Vector2, length int) *Player {
	player := new(Player)

	player.body = make([]*Cell, 0)

	player.styleHead = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
	player.styleBody = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGreen)

	player.body = append(
		player.body,
		NewCell(position, playerVelocity, engine.Vector2Right, player.styleHead),
	)

	for i := 1; i < length; i++ {

		player.body = append(
			player.body,
			NewCell(engine.NewVector2(position.X-float32(i), position.Y), playerVelocity, engine.Vector2Right, player.styleBody),
		)
	}

	player.head = player.body[0]

	return player
}

func (p *Player) checkSelfCollide(x, y float32) bool {
	for i := 1; i < len(p.body); i++ {
		pos := p.body[i].Position()
		x1, y1 := pos.XY()

		if int(x) == int(x1) && int(y) == int(y1) {
			return true
		}
	}

	return false
}

func (p *Player) Position() engine.Vector2 {
	return p.head.Position()
}

func (p *Player) SetPosition(newPos engine.Vector2) {
	p.head.SetPosition(newPos)
}

func (p *Player) Style() tcell.Style {
	return p.styleHead
}

func (p *Player) nextPosition(dt float32) engine.Vector2 {
	vec := engine.Vector2{
		X: p.MoveDirection().X * p.Velocity() * dt,
		Y: p.MoveDirection().Y * p.Velocity() * dt,
	}
	cp := p.head.Position()
	cp.Add(vec)

	return cp
}

func (p *Player) UpdatePhysics(dt float32) {
	curPos := p.Position()
	nextPos := p.nextPosition(dt)

	if int(curPos.X) != int(nextPos.X) || int(curPos.Y) != int(nextPos.Y) {
		for i := len(p.body) - 1; i > 0; i-- {
			p.body[i].SetPosition(
				p.body[i-1].Position(),
			)
		}
	}

	p.head.UpdatePhysics(dt)
}

func (p *Player) Velocity() float32 {
	return p.head.Velocity()
}

func (p *Player) SetVelocity(velocity float32) {
	for _, playerCell := range p.body {
		playerCell.SetVelocity(velocity)
	}
}

func (p *Player) Update() {
	for _, playerCell := range p.body {
		playerCell.Update()
	}
}

func (p *Player) MoveDirection() engine.Vector2 {
	return p.head.MoveDirection()
}

func (p *Player) SetMoveDirection(newDirection engine.Vector2) {
	defer p.Unlock()
	p.Lock()

	curMoveDirection := p.MoveDirection()

	if newDirection.IsEqual(engine.Vector2Up) && curMoveDirection.IsEqual(engine.Vector2Down) {
		return
	}
	if newDirection.IsEqual(engine.Vector2Down) && curMoveDirection.IsEqual(engine.Vector2Up) {
		return
	}
	if newDirection.IsEqual(engine.Vector2Right) && curMoveDirection.IsEqual(engine.Vector2Left) {
		return
	}
	if newDirection.IsEqual(engine.Vector2Left) && curMoveDirection.IsEqual(engine.Vector2Right) {
		return
	}

	p.head.SetMoveDirection(newDirection)
}

func (p *Player) Eat() {
	defer p.Unlock()
	p.Lock()
	ce := p.body[len(p.body)-1]
	pos := ce.Position()
	pos.Sub(ce.MoveDirection())
	p.body = append(p.body, NewCell(pos, p.Velocity(), ce.MoveDirection(), ce.style))
	p.points++
}

func (p *Player) Points() uint {
	return p.points
}

func (p *Player) handleKeyEvents(key tcell.Key) {
	switch key {
	case tcell.KeyLeft:
		p.SetMoveDirection(engine.Vector2Left)
	case tcell.KeyRight:
		p.SetMoveDirection(engine.Vector2Right)
	case tcell.KeyUp:
		p.SetMoveDirection(engine.Vector2Up)
	case tcell.KeyDown:
		p.SetMoveDirection(engine.Vector2Down)
	}
}

func (p *Player) HandleTermEvents(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		p.handleKeyEvents(ev.Key())
	}
}
