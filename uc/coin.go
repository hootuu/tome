package uc

import (
	"github.com/hootuu/tome/vn"
)

type Coin struct {
	Issuer      vn.ID `bson:"issuer" json:"issuer"`
	Code        Code  `bson:"code" json:"code"`
	Issuance    int64 `bson:"issuance" json:"issuance"`
	Balance     int64 `bson:"balance" json:"balance"`
	Circulation int64 `bson:"circulation" json:"circulation"`
}
