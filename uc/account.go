package uc

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/yn"
)

type AccountOwner = ki.ADR

const (
	ISSUER AccountOwner = "$"
)

type AccountLead struct {
	Owner AccountOwner `bson:"owner" json:"owner"`
	Coin  Code         `bson:"coin" json:"coin"`
}

type Account struct {
	Owner   AccountOwner `bson:"owner" json:"owner"`
	Coin    Code         `bson:"code" json:"code"`
	Balance *Amount      `bson:"balance" json:"balance"`
}

type AlterID string

func (a AlterID) S() string {
	return string(a)
}

type AlterType string

const (
	INC AlterType = "+"
	DEC AlterType = "-"
)

type Alter struct {
	ID           AlterID      `bson:"id" json:"id"`
	AlterType    AlterType    `bson:"alter_type" json:"alter_type"`
	AccountOwner AccountOwner `bson:"account_owner" json:"account_owner"`
	Side         AccountOwner `bson:"side" json:"side"`
	Coin         Code         `bson:"coin" json:"coin"`
	Amount       *Amount      `bson:"amount" json:"amount"`
	Yin          yn.YID       `bson:"yin" json:"yin"`
}
