package events

type EventEmitter interface {
	Emit(Event) error
}
