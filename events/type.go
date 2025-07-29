package events

type Fetcher interface {
	Fetche(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Event struct {
	Type Type
	Text string
}

type Type int

const (
	Unknown Type = iota
	Message
)
