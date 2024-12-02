package intermplay

type (
	Transform interface {
		Position() Vector2
		SetPosition(Vector2)
		Velocity() float32
		SetVelocity(float32)
		MoveDirection() Vector2
		SetMoveDirection(Vector2)
		UpdatePhysics(float32)
	}
)
