package events

type EventListener interface {
	Listen(...string) (<-chan Event, <-chan error)
}
