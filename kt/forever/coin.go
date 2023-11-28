package forever

import (
	"github.com/hootuu/tome/kt"
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/tome/vn"
)

const (
	CoinType    kt.Type = "HOTU.COIN"
	CoinVersion         = kt.DefaultVersion
)

type Coin struct {
	Type      kt.Type       `bson:"t" json:"t"`
	Version   kt.Version    `bson:"v" json:"v"`
	Vn        vn.ID         `bson:"vn" json:"vn"`
	Signature *kt.Signature `bson:"s" json:"s"`
	Code      uc.Code       `bson:"code" json:"code"`
	Wei       uc.WEI        `bson:"wei" json:"wei"`
}

func (c *Coin) GetType() kt.Type {
	return c.Type
}

func (c *Coin) GetVersion() kt.Version {
	return c.Version
}

func (c *Coin) GetVN() vn.ID {
	return "HOTU"
}

func (c *Coin) GetSignature() *kt.Signature {
	return c.Signature
}

func (c *Coin) SetSignature(signature *kt.Signature) {
	c.Signature = signature
}

func (c *Coin) Signing() *kt.Signing {
	return kt.NewSigning().
		Add("code", c.Code.S()).
		Add("wei", c.Wei.S())
}
