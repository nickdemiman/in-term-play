package intermplay

type (
	Transform interface {
		// positionBeforeRender() Vector2
		Position() Vector2
		SetPosition(Vector2)

		Velocity() float32
		SetVelocity(float32)

		MoveDirection() Vector2
		SetMoveDirection(Vector2)

		updatePhysics(IGameObject, float32)
		UpdatePhysics(float32)

		// interpoletePhysics(IGameObject, float32)
		// positionBeforeRender() Vector2
		// setPositionBeforeRender(Vector2)
	}
)
