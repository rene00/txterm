package modext

import (
	"math/big"
	"strings"
	"txterm/db/model"
)

type Transaction struct {
	Transaction model.Transaction
}

func NewTransaction(transaction model.Transaction) Transaction {
	return Transaction{Transaction: transaction}
}

func (t Transaction) Increase() string {
	i := &big.Rat{}
	for _, split := range t.Transaction.R.Splits {
		if split.ValueNum == 0 {
			continue
		}
		i.Sub(i, big.NewRat(split.ValueNum, split.ValueDenom))
	}
	return strings.TrimRight(strings.TrimRight(i.FloatString(100), "0"), ".")
}

func (t Transaction) Decrease() string {
	i := &big.Rat{}
	for _, split := range t.Transaction.R.Splits {
		if split.ValueNum == 0 {
			continue
		}
		i.Add(i, big.NewRat(split.ValueNum, split.ValueDenom))
	}
	return strings.TrimRight(strings.TrimRight(i.FloatString(100), "0"), ".")
}

// SourceAccount returns the source account of the transaction as a string.
func (t Transaction) SourceAccount() string {
	sourceAccount := ""
	for _, split := range t.Transaction.R.Splits {
		if split.R != nil {
			if split.R.Account != nil {
				if split.R.Account.R.AccountType != nil {
					if strings.ToLower(split.R.Account.R.AccountType.Name) == "liability" {
						sourceAccount = split.R.Account.Name
						break
					}
				}
			}
		}
	}
	return sourceAccount
}
