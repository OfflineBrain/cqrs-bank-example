package account

import (
	"context"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
	"github.com/offlinebrain/cqrs-bank-example/command-app/infrastructure"
)

type CommandHandler struct {
	storage infrastructure.AggregateRepository[*Aggregate]
}

func NewCommandHandler(storage infrastructure.AggregateRepository[*Aggregate]) *CommandHandler {
	return &CommandHandler{storage: storage}
}

func (c *CommandHandler) handleOpenAccountCommand(ctx context.Context, cmd base.Command) error {
	acc := OpenAccount(ctx, cmd.(OpenAccountCommand))
	return c.storage.Save(ctx, acc)
}

func (c *CommandHandler) handleDepositFundsCommand(ctx context.Context, cmd base.Command) error {
	command := cmd.(DepositFundsCommand)
	acc, _ := c.storage.Get(command.AccountId)
	err := acc.DepositFunds(ctx, command.Amount)
	if err != nil {
		return err
	}
	return c.storage.Save(ctx, acc)
}

func (c *CommandHandler) handleWithdrawFundsCommand(ctx context.Context, cmd base.Command) error {
	command := cmd.(WithdrawFundsCommand)
	acc, _ := c.storage.Get(command.AccountId)
	err := acc.WithdrawFunds(ctx, command.Amount)
	if err != nil {
		return err
	}
	return c.storage.Save(ctx, acc)
}

func (c *CommandHandler) handleCloseAccountCommand(ctx context.Context, cmd base.Command) error {
	command := cmd.(CloseAccountCommand)
	acc, _ := c.storage.Get(command.AccountId)
	err := acc.Close(ctx)
	if err != nil {
		return err
	}
	return c.storage.Save(ctx, acc)
}

func (c *CommandHandler) Register(dispatcher infrastructure.CommandDispatcher) {
	_ = dispatcher.Register(OpenAccountCommandName, c.handleOpenAccountCommand)
	_ = dispatcher.Register(DepositFundsCommandName, c.handleDepositFundsCommand)
	_ = dispatcher.Register(WithdrawFundsCommandName, c.handleWithdrawFundsCommand)
	_ = dispatcher.Register(CloseAccountCommandName, c.handleCloseAccountCommand)
}
