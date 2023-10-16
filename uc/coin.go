package uc

import (
	"github.com/hootuu/tome/vn"
)

type Coin struct {
	Issuer      vn.ID   `bson:"issuer" json:"issuer"`
	Code        Code    `bson:"code" json:"code"`
	Wei         WEI     `bson:"wei" json:"wei"`
	Issuance    *Amount `bson:"issuance" json:"issuance"`
	Balance     *Amount `bson:"balance" json:"balance"`
	Circulation *Amount `bson:"circulation" json:"circulation"`
}
