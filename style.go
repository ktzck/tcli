package tcli

import "github.com/nsf/termbox-go"

type Color termbox.Attribute

const (
	ColorDefault Color = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

type Style struct {
	Bgcolor   Color
	Fgcolor   Color
	Bold      bool
	Underline bool
}

func DefaultStyle() Style {
	return Style{
		Bgcolor:   ColorBlack,
		Fgcolor:   ColorWhite,
		Bold:      false,
		Underline: false,
	}
}

func CursorStyle() Style {
	return Style{
		Bgcolor:   240,
		Fgcolor:   ColorWhite,
		Bold:      true,
		Underline: false,
	}
}
