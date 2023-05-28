package store

import (
	"fmt"
	"txterm/internal/db/query"
)

type Transaction struct {
	query.Tx
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
