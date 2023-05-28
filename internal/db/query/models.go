// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package query

import (
	"database/sql"
	"time"
)

type Import struct {
	ID               int64         `json:"id"`
	DateCreated      time.Time     `json:"date_created"`
	Filename         string        `json:"filename"`
	BalanceAmountNum sql.NullInt64 `json:"balance_amount_num"`
	BalanceAmountDen sql.NullInt64 `json:"balance_amount_den"`
	DateAsOf         sql.NullTime  `json:"date_as_of"`
}

type Tx struct {
	ID          int64     `json:"id"`
	DateCreated time.Time `json:"date_created"`
	DatePosted  time.Time `json:"date_posted"`
	Memo        string    `json:"memo"`
	AmountNum   int64     `json:"amount_num"`
	AmountDen   int64     `json:"amount_den"`
	ImportID    int64     `json:"import_id"`
}
