package tcli

type Text struct {
	View
	Handler EventHandler
	buf     *TextBuffer
	text    string
}

func NewText() *Text {
	v := new(Text)
	v.buf = NewTextBuffer(0, 0, 0, 0)
	return v
}

// Size - impelement View
func (v *Text) Size() (w, h int) {
	return v.buf.w, v.buf.h
}
func (v *Text) Resize(w, h int) {
	v.buf.SetBounds(v.buf.x, v.buf.y, w, h)
}
func (v *Text) Position() (x, y int) {
	return v.buf.x, v.buf.y
}
func (v *Text) Move(x, y int) {
	v.buf.SetBounds(x, y, v.buf.w, v.buf.h)
}
func (v *Text) Draw() {
	v.buf.SetText(v.text)
	v.buf.Draw()
}
func (v *Text) Clear() {
	v.buf.Clear()
}
func (v *Text) Handle(ev Event) {
	if v.Handler != nil {
		v.Handler(ev)
	}
}

func (v *Text) GetText() string {
	return v.text
}

func (v *Text) SetText(text string) {
	v.text = text
	v.buf.SetText(v.text)
}

func (v *Text) AddText(text string) {
	v.text = v.text + "\n" + text
	v.buf.SetText(v.text)
}
