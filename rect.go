package intermplay

type Rect struct {
	origin Vector2
	w, h   int
}

func (rect *Rect) Size() (int, int) {
	return rect.w, rect.h
}

func (rect *Rect) W() int {
	return rect.w
}

func (rect *Rect) H() int {
	return rect.h
}

func (rect *Rect) Origin() Vector2 {
	return rect.origin
}

func NewRect(x, y, width, height int) Rect {
	return Rect{
		origin: NewVector2(x, y),
		w:      width,
		h:      height,
	}
}
