package tcli

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// Cell - cell
type Cell struct {
	ch    rune
	x     int
	y     int
	w     int
	style Style
}

func NewCell(x, y int, ch rune) *Cell {
	return &Cell{
		x:     x,
		y:     y,
		ch:    ch,
		style: DefaultStyle(),
		w:     runewidth.RuneWidth(ch),
	}
}

func (c *Cell) Draw() {
	SetCell(c.x, c.y, c.ch, c.style)
}

// CellBuffer - cell management
type CellBuffer struct {
	x     int
	y     int
	w     int
	h     int
	cells []*Cell
}

func NewCellBuffer(x, y, w, h int) *CellBuffer {
	b := new(CellBuffer)
	b.SetBounds(x, y, w, h)
	return b
}

func (cb *CellBuffer) SetBounds(x, y, w, h int) {
	cb.x = x
	cb.y = y
	cb.w = w
	cb.h = h
	cb.Clear()
}

func (cb *CellBuffer) Clear() {
	cb.cells = cb.cells[0:0]
	for y := 0; y <= cb.h; y++ {
		for x := 0; x <= cb.w; x++ {
			cb.cells = append(cb.cells, NewCell(cb.x+x, cb.y+y, 0))
		}
	}
}

func (cb *CellBuffer) GetCell(x, y int) *Cell {
	if y < 0 || y >= cb.h || x < 0 || x >= cb.w {
		return nil
	}
	return cb.cells[y*(cb.w+1)+x]
}
func (cb *CellBuffer) GetCh(x, y int) rune {
	if c := cb.GetCell(x, y); c != nil {
		return c.ch
	} else {
		return 0x00
	}
}
func (cb *CellBuffer) GetStyle(x, y int) Style {
	if c := cb.GetCell(x, y); c != nil {
		return c.style
	} else {
		return DefaultStyle()
	}
}

func (cb *CellBuffer) SetCell(x, y int, r rune, s Style) {
	if c := cb.GetCell(x, y); c != nil {
		c.ch = r
		c.style = s
	}
}
func (cb *CellBuffer) SetCh(x, y int, r rune) {
	if c := cb.GetCell(x, y); c != nil {
		c.ch = r
	}
}
func (cb *CellBuffer) SetStyle(x, y int, s Style) {
	if c := cb.GetCell(x, y); c != nil {
		c.style = s
	}
}

func (cb *CellBuffer) InsertCh(x, y int, r rune) {
	for i := cb.w - 1; i > x; i-- {
		c := cb.GetCell(i-1, y)
		if c != nil {
			cb.SetCell(i, y, c.ch, c.style)
		}
	}
	cb.SetCh(x, y, r)
}

func (cb *CellBuffer) Remove(x, y int) {
	for i := x; i < cb.w-1; i++ {
		c := cb.GetCell(i+1, y)
		if c != nil {
			cb.SetCell(i, y, c.ch, c.style)
		}
	}
	cb.SetCh(cb.w-1, y, 0x00)
}

func (cb *CellBuffer) Draw() {
	for _, c := range cb.cells {
		c.Draw()
	}
}

type TextBuffer struct {
	CellBuffer
	text string
}

func NewTextBuffer(x, y, w, h int) *TextBuffer {
	b := new(TextBuffer)
	b.SetBounds(x, y, w, h)
	return b
}

func (tb *TextBuffer) SetText(s string) {
	tb.Clear()
	tb.text = s
	lines := strings.Split(s, "\n")
	for y, line := range lines {
		runes := []rune(line)
		for x, r := range runes {
			tb.SetCh(x, y, r)
		}
		if y != len(lines)-1 {
			tb.SetCh(len(runes), y, 0x0A)
		}
	}
}

func (tb *TextBuffer) GetText() string {
	runes := []rune{}
	for _, c := range tb.cells {
		if c.ch != 0x00 {
			runes = append(runes, c.ch)
		}
	}
	return string(runes)
}
