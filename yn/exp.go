package yn

import (
	"encoding/json"
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Currency string

const (
	NilCurrency Currency = ""
)

func CurrencyVerify(str string) *errors.Error {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{3,60}$", str)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid currency: %s", str))
	}
	return nil
}

func CurrencyOf(str string) (Currency, *errors.Error) {
	if err := CurrencyVerify(str); err != nil {
		return NilCurrency, err
	}
	return Currency(strings.ToUpper(str)), nil
}

type ExpenseItem struct {
	Amount   int64    `bson:"amount" json:"amount"`
	Currency Currency `bson:"currency" json:"currency"`
}

type Expense struct {
	Amount   int64                   `bson:"amount" json:"amount"`
	Currency Currency                `bson:"currency" json:"currency"`
	Ex       map[string]*ExpenseItem `bson:"ex,omitempty" json:"ex,omitempty"`
}

func NewExpense(currencyStr string, amount int64) (*Expense, *errors.Error) {
	currency, err := CurrencyOf(currencyStr)
	if err != nil {
		return nil, err
	}
	if amount < 0 {
		return nil, errors.Verify("amount < 0")
	}
	return &Expense{
		Amount:   amount,
		Currency: currency,
		Ex:       nil,
	}, nil
}

func (exp *Expense) S() string {
	jsonByte, _ := json.Marshal(exp)
	return string(jsonByte)
}

func (exp *Expense) PutEx(code string, exExp *ExpenseItem) {
	if exExp == nil {
		return
	}
	if exp.Ex == nil {
		exp.Ex = make(map[string]*ExpenseItem)
	}
	exp.Ex[code] = exExp
}
