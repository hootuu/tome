package bk

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"time"
)

const (
	InvariableVNType    Type = "hotu.vn"
	InvariableVNVersion      = DefaultVersion
)

type VNLead struct {
	Vn         vn.ID `json:"vn"`
	Invariable BID   `json:"invariable"`
}

func (lead *VNLead) S() string {
	return fmt.Sprintf("vn=%s;invariable=%s;", lead.Vn, lead.Invariable)
}

type InvariableVN struct {
	Type      Type    `json:"type"`
	Version   Version `json:"version"`
	Vn        vn.ID   `json:"vn"`
	Guardian  *Ki     `json:"guardian"`
	Timestamp int64   `json:"timestamp"`
	Signature []byte  `json:"signature"`
}

func NewInvariableVN(originBID BID, vnID vn.ID, guardian *ki.Ki, tag []Tag) (*InvariableVN, *errors.Error) {
	i := &InvariableVN{
		Type:      InvariableVNType,
		Version:   InvariableVNVersion,
		Vn:        vnID,
		Guardian:  KiOf(guardian),
		Timestamp: time.Now().UnixMilli(),
	}
	var err *errors.Error
	i.Signature, err = i.Sign(guardian.PRI)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (i *InvariableVN) getSignBuilder() *SignBuilder {
	return NewSignBuilder().
		Add("type", i.Type).
		Add("version", i.Version.S()).
		Add("vn", i.Vn.S()).
		Add("guardian", i.Guardian.ToString()).
		Add("timestamp", fmt.Sprintf("%d", i.Timestamp))
}

func (i *InvariableVN) Sign(pri ki.PRI) ([]byte, *errors.Error) {
	signBuilder := i.getSignBuilder()
	signStr, err := signBuilder.Sign(pri)
	if err != nil {
		return nil, err
	}
	return []byte(signStr), nil
}

func (i *InvariableVN) SignVerify() (bool, *errors.Error) {
	signBuilder := i.getSignBuilder()
	return signBuilder.Verify(ki.PUB(string(i.Guardian.Pub)), string(i.Signature))
}

func (i *InvariableVN) GetType() Type {
	return i.Type
}

func (i *InvariableVN) GetVersion() Version {
	return i.Version
}

func (i *InvariableVN) GetVN() vn.ID {
	return i.Vn
}

func (i *InvariableVN) GetSignature() []byte {
	return i.Signature
}

func (i *InvariableVN) GetTimestamp() int64 {
	return i.Timestamp
}
