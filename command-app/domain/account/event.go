package account

import (
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
)

const (
	OpenAccountV1         = "openAccountV1"
	CloseAccountV1        = "closeAccountV1"
	DepositToAccountV1    = "depositToAccountV1"
	WithdrawFromAccountV1 = "withdrawFromAccountV1"
)

const (
	TopicAccountTransactions = "account-transactions"
)

var Topics = map[string]string{
	OpenAccountV1:         TopicAccountTransactions,
	CloseAccountV1:        TopicAccountTransactions,
	DepositToAccountV1:    TopicAccountTransactions,
	WithdrawFromAccountV1: TopicAccountTransactions,
}

type OpenV1 struct {
	HolderName string `json:"holder_name"`
}

func NewOpenEventV1(id string, v1 OpenV1) *base.Event {
	return base.NewEvent(id, OpenAccountV1, v1)
}

type CloseV1 struct {
}

func NewCloseEventV1(id string, v1 CloseV1) *base.Event {
	return base.NewEvent(id, CloseAccountV1, v1)
}

type DepositV1 struct {
	Amount uint64 `json:"amount"`
}

func NewDepositEventV1(id string, v1 DepositV1) *base.Event {
	return base.NewEvent(id, DepositToAccountV1, v1)
}

type WithdrawV1 struct {
	Amount uint64 `json:"amount"`
}

func NewWithdrawEventV1(id string, v1 WithdrawV1) *base.Event {
	return base.NewEvent(id, WithdrawFromAccountV1, v1)
}
