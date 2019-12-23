package tcli

type Cursor struct {
	x    int
	y    int
	cell Cell
}
type CursorModel interface {
	GetCursor() (x, y int)
	MoveCursor(offx, offy int)
	SetCursor(x, y int)
	ShowCursor(bool)
}

type View interface {
	Size() (w, h int)
	Resize(w, h int)
	Position() (x, y int)
	Move(x, y int)
	Draw()
	Clear()
	Handle(ev Event)
}

type Widget struct {
	View
	Handler EventHandler
	x       int
	y       int
	w       int
	h       int
}

// Size - implement View
func (v *Widget) Size() (w, h int) {
	return v.w, v.h
}

// Resize - implement View
func (v *Widget) Resize(w, h int) {
	v.w = w
	v.h = h
}

// Position - implement View
func (v *Widget) Position() (x, y int) {
	return v.x, v.y
}

// Move - implement View
func (v *Widget) Move(x, y int) {
	v.x = x
	v.y = y
}

// Draw - implement View
func (v *Widget) Draw() {}

// Clear - implement View
func (v *Widget) Clear() {
	for y := 0; y < v.h; y++ {
		for x := 0; x < v.w; x++ {
			NewCell(v.x+x, v.y+y, rune(0)).Draw()
		}
	}
}

// Handle - implement View
func (v *Widget) Handle(ev Event) {
	if v.Handler != nil {
		v.Handler(ev)
	}
}
