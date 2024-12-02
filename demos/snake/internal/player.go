package internal

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	engine "github.com/nickdemiman/in-term-play"
)

type (
	Player struct {
		points        uint
		body          []*Cell
		head          *Cell
		styleHead     tcell.Style
		styleBody     tcell.Style
		velocity      float32
		moveDirection engine.Vector2
		engine.GameObject
		engine.IMoveable
		sync.Mutex
	}
)

var playerVelocity float32 = 1.0

func NewPlayer(position engine.Vector2, length int) *Player {
	player := new(Player)

	player.body = make([]*Cell, 0)

	player.styleHead = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
	player.styleBody = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGreen)

	player.velocity = playerVelocity
	player.moveDirection = engine.Vector2Right

	player.body = append(
		player.body,
		NewCell(position, player.velocity, player.moveDirection, player.styleHead),
	)

	for i := 1; i < length; i++ {

		player.body = append(
			player.body,
			NewCell(engine.NewVector2(position.X-float32(i), position.Y), player.velocity, player.moveDirection, player.styleBody),
		)
	}

	player.head = player.body[0]

	return player
}

func (p *Player) checkSelfCollide(x, y float32) bool {
	for i := 1; i < len(p.body); i++ {
		x1, y1 := p.body[i].position.XY()
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

func (p *Player) UpdatePhysics(dt float32) {
	before := p.head.position

	vec := engine.Vector2{
		X: p.moveDirection.X * p.velocity * dt,
		Y: p.moveDirection.Y * p.velocity * dt,
	}
	cp := p.head.position
	cp.Add(vec)

	if int(before.X) != int(cp.X) || int(before.Y) != int(cp.Y) {
		for i := len(p.body) - 1; i > 0; i-- {
			p.body[i].position = p.body[i-1].position
		}
	}

	p.head.position.Add(vec)
}

func (p *Player) Velocity() float32 {
	return p.velocity
}

func (p *Player) SetVelocity(newVelocity float32) {
	p.velocity = newVelocity
	for i := 1; i < len(_player.body); i++ {
		_player.body[i].velocity = newVelocity
	}
}

func (p *Player) Move() {}

func (p *Player) Update() {
	for _, playerCell := range p.body {
		playerCell.Update()
	}
}

func (p *Player) MoveDirection() engine.Vector2 {
	return p.moveDirection
}

func (p *Player) SetMoveDirection(newDirection engine.Vector2) {
	defer p.Unlock()
	p.Lock()

	if newDirection.IsEqual(engine.Vector2Up) && p.moveDirection.IsEqual(engine.Vector2Down) {
		return
	}
	if newDirection.IsEqual(engine.Vector2Down) && p.moveDirection.IsEqual(engine.Vector2Up) {
		return
	}
	if newDirection.IsEqual(engine.Vector2Right) && p.moveDirection.IsEqual(engine.Vector2Left) {
		return
	}
	if newDirection.IsEqual(engine.Vector2Left) && p.moveDirection.IsEqual(engine.Vector2Right) {
		return
	}

	p.moveDirection = newDirection
}

func (p *Player) Eat() {
	defer p.Unlock()
	p.Lock()
	ce := p.body[len(p.body)-1]
	pos := ce.position
	pos.Sub(ce.moveDirection)
	p.body = append(p.body, NewCell(pos, p.velocity, ce.moveDirection, ce.style))
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
