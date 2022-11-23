package account

import (
	"encoding/json"
	"errors"
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
)

const AggregateType = "Account"

type Aggregate struct {
	base.AggregateRoot
	active  bool
	balance uint64
	holder  string
}

func (a *Aggregate) GetType() string {
	return AggregateType
}

func (a *Aggregate) GetActive() bool {
	return a.active
}

func (a *Aggregate) GetBalance() uint64 {
	return a.balance
}

func OpenAccount(command OpenAccountCommand) *Aggregate {
	var a = &Aggregate{
		AggregateRoot: base.AggregateRoot{
			Id:      command.Id,
			Version: -1,
			Changes: []base.Event{},
		},
		active:  true,
		balance: command.OpeningBalance,
		holder:  command.HolderName,
	}
	a.Raise(*NewOpenEventV1(command.Id, OpenV1{
		command.HolderName,
		command.OpeningBalance,
	}))
	return a
}

func (a *Aggregate) DepositFunds(amount uint64) error {
	if !a.active {
		return errors.New("account is not active")
	}

	a.Raise(*NewDepositEventV1(a.Id, DepositV1{Amount: amount}))

	return nil
}

func (a *Aggregate) WithdrawFunds(amount uint64) error {
	if !a.active {
		return errors.New("account is not active")
	}
	if a.balance < amount {
		return errors.New("insufficient funds")
	}

	a.Raise(*NewWithdrawEventV1(a.Id, WithdrawV1{Amount: amount}))

	return nil
}

func (a *Aggregate) Close() error {
	if !a.active {
		return errors.New("account is not active")
	}
	a.Raise(*NewCloseEventV1(a.Id, CloseV1{}))
	return nil
}

func (a *Aggregate) Apply(event base.Event, isNewEvent bool) {
	switch event.Type {
	case CloseAccountV1:
		a.active = false
		break
	case DepositToAccountV1:
		a.applyDepositV1(event)
		break
	case WithdrawFromAccountV1:
		a.applyWithdrawV1(event)
		break
	case OpenAccountV1:
		a.applyOpenV1(event)
		break
	default:
	}

	if isNewEvent {
		a.Changes = append(a.Changes, event)
	}
}

func (a *Aggregate) applyOpenV1(event base.Event) {
	var data OpenV1
	_ = json.Unmarshal(event.Data, &data)
	a.active = true
	a.balance = data.Balance
	a.holder = data.HolderName
}

func (a *Aggregate) applyDepositV1(event base.Event) {
	var data DepositV1
	_ = json.Unmarshal(event.Data, &data)
	a.balance += data.Amount
}

func (a *Aggregate) applyWithdrawV1(event base.Event) {
	var data WithdrawV1
	_ = json.Unmarshal(event.Data, &data)
	a.balance -= data.Amount
}

func (a *Aggregate) Raise(event base.Event) {
	a.Apply(event, true)
}

func (a *Aggregate) Replay(events []base.Event) {
	for _, event := range events {
		a.Apply(event, false)
	}
}
