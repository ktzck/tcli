package main

import (
	"strings"

	"github.com/ktzck/tcli"
)

var wordlist = []string{
	"1.hello",
	"2.world",
	"3.hoge",
	"4.fuga",
}
var dictionary = []string{
	"1. an expression of greeting\n every morning they exchanged polite hellos",
	"1. all of the living human inhabitants of the earth\n all the world loves\n a lover she always used `humankind' because `mankind' seemed to slight the women\n2. the concerns of this life as distinguished from heaven and the afterlife\n they consider the church to be independent of the world\n3. all of your experiences that determine how things appear to you\n his world was shattered\nwe live in different worlds\nfor them demons were as much a part of reality as trees were",
	"1. hoge hoge hoge hoge hoge",
	"1. fugafugafugafugafugafugafuga",
}

func main() {
	app := tcli.New()

	input := tcli.NewInput()
	display := tcli.NewText()
	menu := tcli.NewText()

	display.SetText("Please input left menu item number")

	menu.SetText(strings.Join(wordlist, "\n"))

	mainc := tcli.NewContainer(tcli.OrientationRow)
	mainc.AddView(menu, 0.2)
	mainc.AddView(display, 0.8)

	app.Root.AddView(mainc, 0.8)
	app.Root.AddView(input, 0.2)

	input.Handler = func(ev tcli.Event) bool {
		switch ev := ev.(type) {
		case *tcli.EventKey:
			switch ev.Key {
			case tcli.KeyEnter:
				found := false
				for i, menuitem := range wordlist {
					if input.GetText() == strings.Split(menuitem, ".")[0] {
						display.SetText(dictionary[i])
						found = true
					}
				}
				if !found {
					display.SetText("Not Found")
				}
				input.SetText("")
				return true
			}
		}
		return true
	}

	app.Run()
}
