package domain

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Account struct {
	ID            string
	Class         AccountClass
	Type          string
	OwnerID       string
	Balance       decimal.Decimal
	AllowNegative bool
	Status        Status
}

func (a *Account) CanDebit(amount decimal.Decimal) error {
	if a.Status != StatusActive {
		return errors.New("account is locked")
	}

	if !a.AllowNegative && a.Balance.Sub(amount).IsNegative() {
		return errors.New("insufficient funds")
	}

	return nil
}

func (a *Account) ApplyPosting(direction Direction, amount decimal.Decimal) {
	if direction == DirectionCredit {
		a.Balance = a.Balance.Add(amount)
	} else {
		a.Balance = a.Balance.Sub(amount)
	}
}

type AccountClass string

const (
	ClassAsset     AccountClass = "ASSET"
	ClassLiability AccountClass = "LIABILITY"
	ClassExpense   AccountClass = "EXPENSE"
	ClassRevenue   AccountClass = "REVENUE"
)

func (c AccountClass) IsNormalCredit() bool {
	return c == ClassLiability || c == ClassRevenue
}

type Status string

const (
	StatusActive Status = "ACTIVE"
	StatusLocked Status = "LOCKED"
	StatusClosed Status = "CLOSED"
)
