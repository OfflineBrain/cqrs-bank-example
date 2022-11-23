package base

type Aggregate interface {
	Apply(event Event, isNewEvent bool)
	Raise(event Event)
	Replay(events []Event)
	GetId() string
	GetVersion() int64
	GetChanges() []Event
	GetType() string
	MarkCommitted()
}

type AggregateRoot struct {
	Id      string
	Version int64
	Changes []Event
}

func (a *AggregateRoot) GetId() string {
	return a.Id
}

func (a *AggregateRoot) GetVersion() int64 {
	return a.Version
}

func (a *AggregateRoot) GetChanges() []Event {
	return a.Changes
}

func (a *AggregateRoot) MarkCommitted() {
	a.Changes = []Event{}
}
