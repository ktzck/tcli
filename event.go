package tcli

type EventType uint16

type Event interface {
	Type() string
}

// EventKey -
type EventKey struct {
	Event
	Key Key
	Ch  rune
}

func (ev *EventKey) Type() string {
	return "key"
}

// EventKey -
type EventMouse struct {
	Event
	Key Key
}

func (ev *EventMouse) Type() string {
	return "mouse"
}

// EventResize -
type EventResize struct {
	Event
	Width  int
	Height int
}

func (ev *EventResize) Type() string {
	return "resize"
}

// EventError -
type EventError struct{ Event }

func (ev *EventError) Type() string {
	return "error"
}

// EventQuit -
type EventQuit struct{ Event }

func (ev *EventQuit) Type() string {
	return "quit"
}

// EventNone -
type EventNone struct {
	Event
}

func (ev *EventNone) Type() string {
	return "none"
}

type EventHandler func(ev Event) bool
