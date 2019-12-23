package tcli

import (
	"strings"
)

type InputModel struct {
	CursorModel
	cursor Cursor
	buf    *TextBuffer
	text   string
}

// GetCursor - implement CursorModel
func (m *InputModel) GetCursor() (x, y int) {
	return m.cursor.x, m.cursor.y
}

// SetCursor - implement CursorModel
func (m *InputModel) SetCursor(x, y int) {
	m.cursor.x = x
	m.cursor.y = y
	m.limitCursor()
}

// MoveCursor - implement CursorModel
func (m *InputModel) MoveCursor(offx, offy int) {
	m.cursor.x += offx
	m.cursor.y += offy
	m.limitCursor()
}

// ShowCursor - implement CursorModel
func (m *InputModel) ShowCursor(show bool) {
	if show {
		m.cursor.cell.style = CursorStyle()
	} else {
		m.cursor.cell.style = DefaultStyle()
	}
}

// ShowCursor - implement CursorModel
func (m *InputModel) DrawCursor() {
	tSetStyle(m.buf.x+m.cursor.x, m.buf.y+m.cursor.y, CursorStyle())
}

func (m *InputModel) limitCursor() {
	x, y := m.GetCursor()
	lines := strings.Split(m.text, "\n")
	if y < 0 {
		y = 0
		m.SetCursor(x, 0)
		return
	}
	if y >= len(lines) {
		y = len(lines) - 1
		m.SetCursor(x, y)
		return
	}
	if x < 0 {
		x = 0
		m.SetCursor(x, y)
		return
	}
	if maxx := len([]rune(lines[y])); x > maxx {
		m.SetCursor(maxx, y)
		return
	}
	if x > m.buf.w {
		x = m.buf.w
		m.SetCursor(x, y)
	}
}

func (m *InputModel) SetText(text string) {
	m.text = text
	m.buf.SetText(m.text)
	m.limitCursor()
}

func (m *InputModel) InsertText(s string) {
	x, y := m.GetCursor()
	rs := []rune(s)
	ls := strings.Split(m.text, "\n")
	l := []rune(ls[y])
	left := l[:x]
	right := l[x:]
	l = append(left, rs...)
	l = append(l, right...)
	ls[y] = string(l)
	m.text = strings.Join(ls, "\n")
	m.buf.SetText(m.text)
	if s == "\n" {
		m.SetCursor(0, y+1)
	} else {
		m.MoveCursor(len(rs), 0)
	}
}

func (m *InputModel) RemoveText() {
	x, y := m.GetCursor()
	ls := strings.Split(m.text, "\n")
	if x <= 0 && y > 0 {
		ls[y-1] += ls[y]
		ls = append(ls[:y-1], ls[y:]...)
		m.buf.SetText(m.text)
		m.SetCursor(len([]rune(ls[y-1])), y-1)
	} else if x > 0 {
		l := []rune(ls[y])
		if len(l) == x {
			l = l[:x-1]
		} else {
			l = append(l[:x], l[x+1:]...)
		}
		ls[y] = string(l)
		m.text = strings.Join(ls, "\n")
		m.buf.SetText(m.text)
		m.MoveCursor(-1, 0)
	}
}

func (m *InputModel) Draw() {
	m.buf.SetText(m.text)
	m.buf.Draw()
	m.DrawCursor()
}

type Input struct {
	View
	Handler   EventHandler
	Multiline bool
	model     *InputModel
	cursor    Cursor
}

func NewInput() *Input {
	v := new(Input)
	v.model = new(InputModel)
	v.model.buf = NewTextBuffer(0, 0, 0, 0)
	v.model.ShowCursor(true)
	return v
}

// Size - impelement View
func (v *Input) Size() (w, h int) {
	return v.model.buf.w, v.model.buf.h
}
func (v *Input) Resize(w, h int) {
	v.model.buf.SetBounds(v.model.buf.x, v.model.buf.y, w, h)
}
func (v *Input) Position() (x, y int) {
	return v.model.buf.x, v.model.buf.y
}
func (v *Input) Move(x, y int) {
	v.model.buf.SetBounds(x, y, v.model.buf.w, v.model.buf.h)
}
func (v *Input) Draw() {
	v.model.Draw()
}
func (v *Input) Clear() {
	v.model.buf.Clear()
}
func (v *Input) Handle(ev Event) {
	switch ev := ev.(type) {
	case *EventKey:
		switch ev.Key {
		case KeyEnter:
			if v.Multiline {
				v.NewLine()
			}
		case KeyBackspace2:
			fallthrough
		case KeyBackspace:
			v.model.RemoveText()
		case KeyArrowUp:
			v.model.MoveCursor(0, -1)
		case KeyArrowDown:
			v.model.MoveCursor(0, 1)
		case KeyArrowLeft:
			v.model.MoveCursor(-1, 0)
		case KeyArrowRight:
			v.model.MoveCursor(1, 0)
		default:
			if ev.Ch != 0 {
				v.model.InsertText(string(ev.Ch))
			}
		}
	}
	if v.Handler != nil {
		v.Handler(ev)
	}
}

func (v *Input) NewLine() {
	v.model.InsertText("\n")
}

func (v *Input) GetText() string {
	return v.model.text
}
func (v *Input) SetText(text string) {
	v.model.SetText(text)
}
