package account

import "github.com/offlinebrain/cqrs-bank-example/command-app/base"

const (
	OpenAccountCommandName   = "OpenAccountCommand"
	DepositFundsCommandName  = "DepositFundsCommand"
	WithdrawFundsCommandName = "WithdrawFundsCommand"
	CloseAccountCommandName  = "CloseAccountCommand"
)

type OpenAccountCommand struct {
	base.CommandBase
	HolderName     string
	AccountType    string
	OpeningBalance uint64
}

func (o OpenAccountCommand) GetName() string {
	return OpenAccountCommandName
}

type DepositFundsCommand struct {
	base.CommandBase
	Amount uint64
}

func (d DepositFundsCommand) GetName() string {
	return DepositFundsCommandName
}

type WithdrawFundsCommand struct {
	base.CommandBase
	Amount uint64
}

func (w WithdrawFundsCommand) GetName() string {
	return WithdrawFundsCommandName
}

type CloseAccountCommand struct {
	base.CommandBase
}

func (c CloseAccountCommand) GetName() string {
	return CloseAccountCommandName
}
