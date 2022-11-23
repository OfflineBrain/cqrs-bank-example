package infrastructure

import "github.com/offlinebrain/cqrs-bank-example/command-app/base"

type EventPublisher interface {
	Publish(topic string, event base.EventModel) error
}

type MockEventPublisher struct {
}

func (m *MockEventPublisher) Publish(_ string, _ base.EventModel) error {
	return nil
}

func NewMockEventPublisher() *MockEventPublisher {
	return &MockEventPublisher{}
}
