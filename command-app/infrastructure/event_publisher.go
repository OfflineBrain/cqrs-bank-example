package infrastructure

import (
	"context"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
)

type EventPublisher interface {
	Publish(ctx context.Context, topic string, event base.EventModel) error
}

type MockEventPublisher struct {
}

func (m *MockEventPublisher) Publish(_ context.Context, _ string, _ base.EventModel) error {
	return nil
}

func NewMockEventPublisher() *MockEventPublisher {
	return &MockEventPublisher{}
}
