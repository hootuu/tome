package uc

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/yn"
)

type AccountID = ki.ADR

const (
	ISSUER AccountID = "$"
)

type Account struct {
	ID      AccountID `bson:"id" json:"id"`
	Coin    Code      `bson:"code" json:"code"`
	Balance int64     `bson:"balance" json:"balance"`
}

type AlterID string

func (a AlterID) S() string {
	return string(a)
}

type Alter struct {
	ID        AlterID   `bson:"id" json:"id"`
	AccountID AccountID `bson:"account_id" json:"account_id"`
	Side      AccountID `bson:"side" json:"side"`
	Coin      Code      `bson:"coin" json:"coin"`
	Amount    int64     `bson:"amount" json:"amount"`
	Yin       yn.YID    `bson:"yin" json:"yin"`
}
