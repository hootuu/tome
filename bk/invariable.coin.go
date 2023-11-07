package bk

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"strconv"
	"time"
)

const (
	InvariableCoinType    Type = "hotu.coin"
	InvariableCoinVersion      = DefaultVersion
)

type CoinLead struct {
	Coin       uc.Code `json:"coin"`
	Invariable BID     `json:"invariable"`
}

func (lead *CoinLead) S() string {
	return fmt.Sprintf("coin=%s;invariable=%s", lead.Coin, lead.Invariable)
}

type InvariableCoin struct {
	Type      Type       `json:"type"`
	Version   Version    `json:"version"`
	Vn        *VNLead    `json:"vn"`
	Coin      *CoinLead  `bson:"coin" json:"coin"`
	Wei       uc.WEI     `bson:"wei" json:"wei"`
	Infinite  bool       `bson:"infinite" json:"infinite"`
	Issuance  *uc.Amount `bson:"issuance" json:"issuance"`
	Timestamp int64      `json:"timestamp"`
	Signature []byte     `json:"signature"`
}

func NewInvariableCoin(
	vnLead *VNLead,
	coin *CoinLead,
	wei uc.WEI,
	infinite bool,
	issuance *uc.Amount,
	guardian *ki.Ki,
) (*InvariableCoin, *errors.Error) {
	i := &InvariableCoin{
		Type:      InvariableCoinType,
		Version:   InvariableCoinVersion,
		Vn:        vnLead,
		Coin:      coin,
		Wei:       wei,
		Infinite:  infinite,
		Issuance:  issuance,
		Timestamp: time.Now().UnixMilli(),
		Signature: nil,
	}
	var err *errors.Error
	i.Signature, err = i.Sign(guardian.PRI)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (i *InvariableCoin) getSignBuilder() *SignBuilder {
	return NewSignBuilder().
		Add("type", i.Type.S()).
		Add("version", i.Version.S()).
		Add("vn", i.Vn.S()).
		Add("coin", i.Coin.S()).
		Add("wei", i.Wei.S()).
		Add("infinite", strconv.FormatBool(i.Infinite)).
		Add("issuance", i.Issuance.S()).
		Add("timestamp", fmt.Sprintf("%d", i.Timestamp))
}

func (i *InvariableCoin) Sign(pri ki.PRI) ([]byte, *errors.Error) {
	signBuilder := i.getSignBuilder()
	signStr, err := signBuilder.Sign(pri)
	if err != nil {
		return nil, err
	}
	return []byte(signStr), nil
}

func (i *InvariableCoin) SignVerify(guardianPub ki.PUB) (bool, *errors.Error) {
	signBuilder := i.getSignBuilder()
	return signBuilder.Verify(guardianPub, string(i.Signature))
}

func (i *InvariableCoin) GetType() Type {
	return i.Type
}

func (i *InvariableCoin) GetVersion() Version {
	return i.Version
}

func (i *InvariableCoin) GetVN() vn.ID {
	return i.Vn.Vn
}

func (i *InvariableCoin) GetSignature() []byte {
	return i.Signature
}

func (i *InvariableCoin) GetTimestamp() int64 {
	return i.Timestamp
}
