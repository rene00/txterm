package store

import (
	"fmt"
	"math/big"
	"strings"
	"txterm/db/query"
)

type Transaction struct {
	query.Tx
}

func (t Transaction) Date() string {
	return t.Tx.DatePosted.Format("2006-01-02")
}

func (t Transaction) Memo() string {
	return t.Tx.Memo
}

func (t Transaction) Amount() string {
	return strings.TrimRight(strings.TrimRight(big.NewRat(t.Tx.AmountNum, t.Tx.AmountDen).FloatString(100), "0"), ".")
}

type TransactionError struct {
	code int
	msg  string
}

var (
	TransactionErrDuplicate = 1
)

func (e *TransactionError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.msg)
}

func (e *TransactionError) Code() int {
	return e.code
}

func TransactionDuplicateError() error {
	return &TransactionError{code: TransactionErrDuplicate, msg: "duplicate transaction"}
}
