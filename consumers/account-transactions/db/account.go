package db

type Account struct {
	Id         string
	HolderName string
	Balance    uint64
	Active     bool
}
