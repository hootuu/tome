package bk

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/utils/errors"
	"time"
)

const (
	InvariableTxType    Type = "hotu.tx"
	InvariableTxVersion      = DefaultVersion
)

type InvariableTx struct {
	Type      Type       `json:"type"`
	Version   Version    `json:"version"`
	Tx        string     `json:"tx"`
	Coin      *CoinLead  `json:"coin"`
	From      ki.ADR     `json:"from"`
	To        ki.ADR     `json:"to"`
	Amount    *uc.Amount `json:"amount"`
	Timestamp int64      `json:"timestamp"`
	Signature []byte     `json:"signature"`
}

func NewInvariableTx(
	coin *CoinLead,
	from ki.ADR,
	to ki.ADR,
	amount *uc.Amount,
	guardianPri ki.PRI,
) (*InvariableTx, *errors.Error) {
	i := &InvariableTx{
		Type:      InvariableTxType,
		Version:   InvariableTxVersion,
		Tx:        "",
		Coin:      coin,
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now().UnixMilli(),
		Signature: nil,
	}
	var err *errors.Error
	i.Signature, err = i.Sign(guardianPri)
	if err != nil {
		return nil, err
	}
	i.Tx = string(i.Signature)
	return i, nil
}

func (i *InvariableTx) getSignBuilder() *SignBuilder {
	return NewSignBuilder().
		Add("type", i.Type).
		Add("version", i.Version.S()).
		Add("tx", i.Tx).
		Add("coin", i.Coin.S()).
		Add("from", i.From.S()).
		Add("to", i.To.S()).
		Add("amount", i.Amount.S()).
		Add("timestamp", fmt.Sprintf("%d", i.Timestamp))
}

func (i *InvariableTx) Sign(pri ki.PRI) ([]byte, *errors.Error) {
	signBuilder := i.getSignBuilder()
	signStr, err := signBuilder.Sign(pri)
	if err != nil {
		return nil, err
	}
	return []byte(signStr), nil
}

func (i *InvariableTx) SignVerify(guardianPub ki.PUB) (bool, *errors.Error) {
	signBuilder := i.getSignBuilder()
	return signBuilder.Verify(guardianPub, string(i.Signature))
}

func (i *InvariableTx) GetType() Type {
	return i.Type
}

func (i *InvariableTx) GetVersion() Version {
	return i.Version
}

func (i *InvariableTx) GetSignature() []byte {
	return i.Signature
}

func (i *InvariableTx) GetTimestamp() int64 {
	return i.Timestamp
}
