package account

import (
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
)

type CommandHandler struct {
	storage infrastructure.AggregateRepository[*Aggregate]
}

func NewCommandHandler(storage infrastructure.AggregateRepository[*Aggregate]) *CommandHandler {
	return &CommandHandler{storage: storage}
}

func (c *CommandHandler) handleOpenAccountCommand(cmd base.Command) error {
	acc := OpenAccount(cmd.(OpenAccountCommand))
	return c.storage.Save(acc)
}

func (c *CommandHandler) handleDepositFundsCommand(cmd base.Command) error {
	command := cmd.(DepositFundsCommand)
	acc, _ := c.storage.Get(command.Id)
	err := acc.DepositFunds(command.Amount)
	if err != nil {
		return err
	}
	return c.storage.Save(acc)
}

func (c *CommandHandler) handleWithdrawFundsCommand(cmd base.Command) error {
	command := cmd.(WithdrawFundsCommand)
	acc, _ := c.storage.Get(command.Id)
	err := acc.WithdrawFunds(command.Amount)
	if err != nil {
		return err
	}
	return c.storage.Save(acc)
}

func (c *CommandHandler) handleCloseAccountCommand(cmd base.Command) error {
	command := cmd.(CloseAccountCommand)
	acc, _ := c.storage.Get(command.Id)
	err := acc.Close()
	if err != nil {
		return err
	}
	return c.storage.Save(acc)
}

func (c *CommandHandler) Register(dispatcher infrastructure.CommandDispatcher) {
	_ = dispatcher.Register(OpenAccountCommandName, c.handleOpenAccountCommand)
	_ = dispatcher.Register(DepositFundsCommandName, c.handleDepositFundsCommand)
	_ = dispatcher.Register(WithdrawFundsCommandName, c.handleWithdrawFundsCommand)
	_ = dispatcher.Register(CloseAccountCommandName, c.handleCloseAccountCommand)
}
