package intermplay

type (
	Transform interface {
		Position() Vector2
		SetPosition(Vector2)
	}

	IMoveable interface {
		Velocity() uint
		SetVelocity(uint)
		MoveDirection() Vector2
		SetMoveDirection(Vector2)
		Move()
	}
)
