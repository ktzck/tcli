package tcli

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

var quit = make(chan int)

type App struct {
	currentView View
	Root        *Container
	Handler     EventHandler
	QuitKey     Key
}

func New() *App {
	app := &App{}
	app.Root = NewContainer(OrientationColumn)
	app.currentView = app.Root
	app.QuitKey = KeyCtrlC
	return app
}

func (a *App) Draw() {
	a.Root.Clear()
	a.Root.Draw()
	termbox.Flush()
}
func (a *App) Clear() {
	a.Root.Clear()
	termbox.Flush()
}

func (a *App) Run() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	termbox.SetOutputMode(termbox.Output256)
	a.Root.Resize(termbox.Size())

	go func() {
	eventloop:
		for {
			a.Draw()
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				tev := &EventKey{}
				tev.Key = Key(ev.Key)
				tev.Ch = ev.Ch
				if tev.Key == a.QuitKey {
					fmt.Println("ktzck/tcli app quit")
					if a.Handler != nil {
						if a.Handler(&EventQuit{}) {
							break eventloop
						}
					} else {
						break eventloop
					}
				}
				if a.Handler != nil {
					if a.Handler(tev) {
						a.currentView.Handle(tev)
					}
				} else {
					a.currentView.Handle(tev)
				}
			case termbox.EventResize:
				tev := &EventResize{}
				w, h := termbox.Size()
				tev.Width = w
				tev.Height = h
				a.Root.Resize(w, h)
				a.Root.Draw()
				if a.Handler != nil {
					if a.Handler(tev) {
						a.currentView.Handle(tev)
					}
				} else {
					a.currentView.Handle(tev)
				}
			case termbox.EventMouse:
				tev := &EventMouse{}
				tev.Key = Key(ev.Key)
				if a.Handler != nil {
					if a.Handler(tev) {
						a.currentView.Handle(tev)
					}
				} else {
					a.currentView.Handle(tev)
				}
			}
		}
	}()

	for {
		select {
		case <-quit:
			termbox.Close()
			return nil
		}
	}

}

func (a *App) Close() {
	close(quit)
}
