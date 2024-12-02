package intermplay

type Rect struct {
	origin Vector2
	w, h   float32
}

func (rect *Rect) Size() (float32, float32) {
	return rect.w, rect.h
}

func (rect *Rect) W() float32 {
	return rect.w
}

func (rect *Rect) H() float32 {
	return rect.h
}

func (rect *Rect) Origin() Vector2 {
	return rect.origin
}

func NewRect(x, y, width, height float32) Rect {
	return Rect{
		origin: NewVector2(float32(x), float32(y)),
		w:      width,
		h:      height,
	}
}
