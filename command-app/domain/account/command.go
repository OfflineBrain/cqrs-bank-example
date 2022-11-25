package account

const (
	OpenAccountCommandName   = "OpenAccountCommand"
	DepositFundsCommandName  = "DepositFundsCommand"
	WithdrawFundsCommandName = "WithdrawFundsCommand"
	CloseAccountCommandName  = "CloseAccountCommand"
)

type OpenAccountCommand struct {
	AccountId      string
	HolderName     string
	AccountType    string
	OpeningBalance uint64
}

func (o OpenAccountCommand) GetName() string {
	return OpenAccountCommandName
}

type DepositFundsCommand struct {
	AccountId string
	Amount    uint64
}

func (d DepositFundsCommand) GetName() string {
	return DepositFundsCommandName
}

type WithdrawFundsCommand struct {
	AccountId string
	Amount    uint64
}

func (w WithdrawFundsCommand) GetName() string {
	return WithdrawFundsCommandName
}

type CloseAccountCommand struct {
	AccountId string
}

func (c CloseAccountCommand) GetName() string {
	return CloseAccountCommandName
}
