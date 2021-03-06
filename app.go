package tcli

import (
	"github.com/nsf/termbox-go"
)

type App struct {
	currentView View
	Root        *Container
	Handler     EventHandler
	QuitKey     Key
	QuitChannel chan int
}

func New() *App {
	app := &App{}
	app.Root = NewContainer(OrientationColumn)
	app.currentView = app.Root
	app.QuitKey = KeyCtrlC
	app.QuitChannel = make(chan int)
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
		for {
			a.Draw()
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				tev := &EventKey{}
				tev.Key = Key(ev.Key)
				tev.Ch = ev.Ch
				if tev.Key == a.QuitKey {
					a.Close()
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
				tev.Width, tev.Height = termbox.Size()
				a.Root.Resize(tev.Width, tev.Height)
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
			case termbox.EventInterrupt:
				tev := &EventQuit{}
				if a.Handler != nil {
					if a.Handler(tev) {
						close(a.QuitChannel)
					}
				} else {
					close(a.QuitChannel)
				}
			}
		}
	}()

	select {
	case <-a.QuitChannel:
		termbox.Close()
		return nil
	}
}

func (a *App) Close() {
	termbox.Interrupt()
}
