package infrastructure

import (
	"context"
	"errors"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	l "github.com/offlinebrain/cqrs-bank-example/command-app/log"
	"github.com/sirupsen/logrus"
)

type CommandDispatcher struct {
	handlers map[string]base.CommandHandler
}

func NewCommandDispatcher() *CommandDispatcher {
	handlers := make(map[string]base.CommandHandler, 1)
	return &CommandDispatcher{
		handlers: handlers,
	}
}

func (c *CommandDispatcher) Register(cmd string, handler base.CommandHandler) error {
	if val := c.handlers[cmd]; val != nil {
		return errors.New("dispatcher: handler for this command already registered")
	}
	c.handlers[cmd] = handler
	return nil
}

func (c *CommandDispatcher) Send(ctx context.Context, cmd base.Command) error {
	log := l.Logger.WithFields(logrus.Fields{
		TraceIdKey: ctx.Value(TraceIdKey),
	})
	log.Infof("dispatcher: received command(%v)", cmd)
	if handler := c.handlers[cmd.GetName()]; handler != nil {
		return handler(ctx, cmd)
	}
	return errors.New("dispatcher: handler for this command not found")
}
