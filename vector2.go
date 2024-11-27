package core

import "errors"

type (
	// Vector2 interface {
	// 	X() int
	// 	Y() int
	// 	SetX(int)
	// 	SetY(int)
	// 	Add(Vector2)
	// 	Sub(Vector2)
	// 	Multiply(int)
	// 	Divide(int) error
	// 	IsEqual(Vector2) bool
	// }

	Vector2 struct {
		X, Y int
	}
)

func NewVector2(x, y int) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

var (
	Vector2Up = Vector2{
		X: 0,
		Y: -1,
	}

	Vector2Down = Vector2{
		X: 0,
		Y: 1,
	}

	Vector2Left = Vector2{
		X: -1,
		Y: 0,
	}

	Vector2Right = Vector2{
		X: 1,
		Y: 0,
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

func (src *Vector2) Add(dst Vector2) {
	src.X += dst.X
	src.Y += dst.Y
}

func (src *Vector2) Sub(dst Vector2) {
	src.X -= dst.X
	src.Y -= dst.Y
}

func (src *Vector2) Multiply(mult int) {
	src.X *= mult
	src.Y *= mult
}

func (src *Vector2) Divide(div int) error {
	if div == 0 {
		return errors.New("деление на 0")
	}

	src.X /= div
	src.Y /= div

	return nil
}

func (vec *Vector2) IsEqual(vec2 Vector2) bool {
	if vec.X == vec2.X && vec.Y == vec2.Y {
		return true
	}

	return false
}
