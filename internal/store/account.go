package store

type Account struct {
	id          int64
	name        string
	description string
	accountType AccountType
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Description() string {
	return a.description
}

func (a Account) Type() string {
	return a.accountType.Name
}
