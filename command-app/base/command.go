package base

type CommandBase struct {
	Id string
}

type Command interface {
	GetName() string
}

type CommandHandler func(cmd Command) error

type CommandDispatcher interface {
	Register(cmd string, handler CommandHandler) error
	Send(cmd Command) error
}
