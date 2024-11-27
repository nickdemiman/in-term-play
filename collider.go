package core

import (
	"time"
)

type (
	Collider interface {
		Collides(x, y int) bool
		Bounds() ColliderMap

		EventHandler
	}

	EventCollision struct {
		when      time.Time
		initiator Collider
		target    Collider
	}
)

func (e *EventCollision) When() time.Time {
	return e.when
}

func (e *EventCollision) Initiator() Collider {
	return e.initiator
}

func (e *EventCollision) Target() Collider {
	return e.target
}

func HandleCollision(init, target Collider) {
	when := time.Now()

	eventInit := &EventCollision{when: when, initiator: init, target: target}
	eventTarget := &EventCollision{when: when, initiator: target, target: init}

	init.HandleEvent(eventInit)
	target.HandleEvent(eventTarget)
}
