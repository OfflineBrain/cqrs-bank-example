package account

import (
	"github.com/offlinebrain/cqrs-bank-example/command-app/base"
)

const (
	OpenAccountV1         = "openAccountV1"
	CloseAccountV1        = "closeAccountV1"
	DepositToAccountV1    = "depositToAccountV1"
	WithdrawFromAccountV1 = "withdrawFromAccountV1"
	ClearEvent            = "clear"
)

func Topics(topic string) map[string]string {
	return map[string]string{
		OpenAccountV1:         topic,
		CloseAccountV1:        topic,
		DepositToAccountV1:    topic,
		WithdrawFromAccountV1: topic,
		ClearEvent:            topic,
	}
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

type Clear struct {
}

func NewClearEvent(id string) *base.Event {
	return base.NewEvent(id, ClearEvent, Clear{})
}
