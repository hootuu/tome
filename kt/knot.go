package kt

import (
	"github.com/hootuu/tome/bk/bid"
	"github.com/hootuu/tome/vn"
)

type KID = bid.BID

const (
	KnotType    Type    = "HOTU.KNOT"
	KnotVersion Version = DefaultVersion
)

type Knot struct {
	Type       Type       `json:"t"`
	Version    Version    `json:"v"`
	Vn         vn.ID      `json:"vn"`
	Chain      Chain      `json:"c"`
	Height     Height     `json:"h"`
	Invariable *Lead      `json:"i"`
	Previous   KID        `json:"p"`
	Signature  *Signature `json:"s"`
}

func (k *Knot) GetType() Type {
	return KnotType
}

func (k *Knot) GetVersion() Version {
	return KnotVersion
}

func (k *Knot) GetVN() vn.ID {
	return k.Vn
}

func (k *Knot) Signing() *Signing {
	return NewSigning().
		Add("chain", k.Chain.S()).
		Add("height", k.Height.S()).
		Add("invariable", k.Invariable.S()).
		Add("previous", k.Previous.S())
}

func (k *Knot) GetSignature() *Signature {
	return k.Signature
}

func (k *Knot) Sign(signature *Signature) {
	k.Signature = signature
}
