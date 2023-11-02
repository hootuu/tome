package bk

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"time"
)

const (
	TieType    Type = "hotu.tie"
	TieVersion      = DefaultVersion
)

type Tie struct {
	Tie        string          `json:"tie"`
	Type       Type            `json:"type"`
	Version    Version         `json:"version"`
	Vn         vn.ID           `json:"vn"`
	Chain      Chain           `json:"chain"`
	Invariable *InvariableLead `json:"invariable"`
	Timestamp  int64           `json:"timestamp"`
	Signature  []byte          `json:"signature"`
}

func NewTie(chain Chain, invariable Invariable, vnGuardianPri ki.PRI) (*Tie, *errors.Error) {
	invariableBID, err := BuildBID(invariable)
	if err != nil {
		return nil, err
	}
	i := &Tie{
		Type:    TieType,
		Version: TieVersion,
		Vn:      invariable.GetVN(),
		Chain:   chain,
		Invariable: &InvariableLead{
			Type:       invariable.GetType(),
			Version:    invariable.GetVersion(),
			Invariable: invariableBID,
		},
		Timestamp: time.Now().UnixMilli(),
	}
	i.Signature, err = i.Sign(vnGuardianPri)
	if err != nil {
		return nil, err
	}
	i.Tie = string(i.Signature)
	return i, nil
}

func (i *Tie) getSignBuilder() *SignBuilder {
	return NewSignBuilder().
		Add("type", i.Type).
		Add("version", i.Version.S()).
		Add("vn", i.Vn.S()).
		Add("chain", i.Chain.S()).
		Add("invariable", i.Invariable.S()).
		Add("timestamp", fmt.Sprintf("%d", i.Timestamp))
}

func (i *Tie) Sign(pri ki.PRI) ([]byte, *errors.Error) {
	signBuilder := i.getSignBuilder()
	signStr, err := signBuilder.Sign(pri)
	if err != nil {
		return nil, err
	}
	return []byte(signStr), nil
}

func (i *Tie) SignVerify(vnGuardianPub ki.PUB) (bool, *errors.Error) {
	signBuilder := i.getSignBuilder()
	return signBuilder.Verify(vnGuardianPub, string(i.Signature))
}
