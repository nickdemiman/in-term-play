package intermplay

import "errors"

type (
	Vector2 struct {
		X, Y float32
	}
)

func NewVector2(x, y float32) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

var (
	Vector2Up = Vector2{
		X: 0.0,
		Y: -yDelta,
	}

	Vector2Down = Vector2{
		X: 0.0,
		Y: yDelta,
	}

	Vector2Left = Vector2{
		X: -xDelta,
		Y: 0.0,
	}

	Vector2Right = Vector2{
		X: xDelta,
		Y: 0.0,
	}

	Vector2Null = Vector2{}
)

func CopyVector2(src Vector2) *Vector2 {
	vec := new(Vector2)

	vec.X = src.X
	vec.Y = src.Y

	return vec
}

func Vector2Add(first, second Vector2) *Vector2 {
	out := CopyVector2(first)

	out.X += second.X
	out.Y += second.Y

	return out
}

func Vector2Sub(first, second Vector2) *Vector2 {
	out := CopyVector2(first)

	out.X -= second.X
	out.Y -= second.Y

	return out
}

func (src Vector2) XY() (float32, float32) {
	return src.X, src.Y
}

func (src *Vector2) Add(dst Vector2) {
	src.X += dst.X
	src.Y += dst.Y
}

func (src *Vector2) Sub(dst Vector2) {
	src.X -= dst.X
	src.Y -= dst.Y
}

func (src *Vector2) Multiply(mult float32) {
	src.X *= mult
	src.Y *= mult
}

func (src *Vector2) Divide(div float32) error {
	if div == 0.0 {
		return errors.New("деление на 0")
	}

	src.X /= div
	src.Y /= div

	return nil
}

func (vec Vector2) IsEqual(vec2 Vector2) bool {
	if vec.X == vec2.X && vec.Y == vec2.Y {
		return true
	}

	return false
}
