package base

import "context"

type Command interface {
	GetName() string
}

type CommandHandler func(ctx context.Context, cmd Command) error

type CommandDispatcher interface {
	Register(cmd string, handler CommandHandler) error
	Send(ctx context.Context, cmd Command) error
}
