package internal

import (
	"github.com/gdamore/tcell/v2"
	engine "github.com/nickdemiman/in-term-play"
)

type (
	// Player interface {
	// 	Eat()
	// 	GameObject
	// 	IMoveable
	// }

	//	player struct {
	//		name          string
	//		body          []Cell
	//		head          Cell
	//		styleHead     tcell.Style
	//		styleBody     tcell.Style
	//		moveDirection Vector2
	//		colliderMap   ColliderMap
	//		velocity      uint
	//	}

	Player struct {
		name          string
		body          []*Cell
		head          *Cell
		styleHead     tcell.Style
		styleBody     tcell.Style
		moveDirection engine.Vector2
		// colliderMap   ColliderMap
		velocity uint
		engine.GameObject
		engine.IMoveable
	}
)

func NewPlayer(position engine.Vector2, length int) *Player {
	player := new(Player)

	player.name = "player"
	player.body = make([]*Cell, 0)

	player.styleHead = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
	player.styleBody = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGreen)

	player.body = append(
		player.body,
		NewCell(position, ' ', player.styleHead),
	)

	for i := 1; i < length; i++ {

		player.body = append(
			player.body,
			NewCell(engine.NewVector2(position.X-i, position.Y), ' ', player.styleBody),
		)
	}

	player.velocity = 1
	player.moveDirection = engine.Vector2Right
	player.head = player.body[0]

	return player
}

func (p *Player) Name() string {
	return p.name
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

func (p *Player) Velocity() uint {
	return p.velocity
}

func (p *Player) SetVelocity(newVelocity uint) {
	p.velocity = newVelocity
}

func (p *Player) Move() {
	for i := len(p.body) - 1; i > 0; i-- {
		p.body[i].SetPosition(p.body[i-1].Position())
	}

	p.body[0].SetPosition(
		*engine.Vector2Add(p.body[0].Position(), p.moveDirection),
	)
}

func (p *Player) Awake() {}

func (p *Player) Update() {
	for _, playerCell := range p.body {
		playerCell.Update()
	}
}

func (p *Player) Dispose() {}

func (p *Player) MoveDirection() engine.Vector2 {
	return p.moveDirection
}

func (p *Player) SetMoveDirection(newDirection engine.Vector2) {
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

func (p *Player) HandleEvent(ev tcell.Event) {}

func (p *Player) Eat() {
	pos := p.body[len(p.body)-1].Position()
	pos.X -= 1
	p.body = append(p.body, NewCell(pos, ' ', p.styleBody))
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

func (p *Player) handleTermEvents(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		p.handleKeyEvents(ev.Key())
	}
}
