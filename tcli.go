/*
 * tcli - terminal cui light interface
 */
package tcli

import (
	"github.com/nsf/termbox-go"
)

func SetCell(x, y int, r rune, s Style) {
	fg := s.Fgcolor
	if s.Bold {
		fg = fg | 1<<9
	}
	if s.Underline {
		fg = fg | 1<<10
	}
	termbox.SetCell(x, y, r, termbox.Attribute(fg), termbox.Attribute(s.Bgcolor))
}

func tGetCell(x, y int) *Cell {
	w, _ := termbox.Size()
	buf := termbox.CellBuffer()
	c := buf[w*y+x]
	s := Style{}
	s.Bgcolor = Color(c.Bg)
	s.Fgcolor = Color(c.Fg & ^termbox.AttrBold & ^termbox.AttrUnderline)
	if (c.Fg & ^termbox.AttrBold) > 0 {
		s.Bold = true
	}
	if (c.Fg & ^termbox.AttrUnderline) > 0 {
		s.Underline = true
	}
	nc := NewCell(x, y, c.Ch)
	nc.style = s
	return nc
}

func tSetCh(x, y int, r rune) {
	w, _ := termbox.Size()
	buf := termbox.CellBuffer()
	c := &buf[w*y+x]
	c.Ch = r
}

func tSetStyle(x, y int, s Style) {
	w, _ := termbox.Size()
	buf := termbox.CellBuffer()
	c := &buf[w*y+x]
	c.Bg = termbox.Attribute(s.Bgcolor)
	c.Fg = termbox.Attribute(s.Fgcolor)
	if s.Bold {
		c.Fg = c.Fg | termbox.AttrBold
	}
	if s.Bold {
		c.Fg = c.Fg | termbox.AttrUnderline
	}
}

func tSetCell(x, y int, r rune, s Style) {
	tSetCh(x, y, r)
	tSetStyle(x, y, s)
}
