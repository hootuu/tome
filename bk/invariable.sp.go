package bk

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"time"
)

const (
	InvariableSPType    Type = "hotu.sp"
	InvariableSPVersion      = DefaultVersion
)

type SPLead struct {
	Sp         sp.ID `json:"sp"`
	Invariable BID   `json:"invariable"`
}

func (lead *SPLead) S() string {
	return fmt.Sprintf("sp=%s;invariable=%s;", lead.Sp, lead.Invariable)
}

type InvariableSP struct {
	Type      Type    `json:"type"`
	Version   Version `json:"version"`
	Vn        *VNLead `json:"vn"`
	Sp        sp.ID   `json:"sp"`
	Guardian  *Ki     `json:"guardian"`
	Timestamp int64   `json:"timestamp"`
	Signature []byte  `json:"signature"`
}

func NewInvariableSP(vnLead *VNLead, spID sp.ID, guardian *ki.Ki, tag []Tag) (*InvariableSP, *errors.Error) {
	i := &InvariableSP{
		Type:      InvariableSPType,
		Version:   InvariableSPVersion,
		Vn:        vnLead,
		Sp:        spID,
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

func (i *InvariableSP) getSignBuilder() *SignBuilder {
	return NewSignBuilder().
		Add("type", i.Type).
		Add("version", i.Version.S()).
		Add("vn", i.Vn.S()).
		Add("sp", i.Sp.S()).
		Add("guardian", i.Guardian.ToString()).
		Add("timestamp", fmt.Sprintf("%d", i.Timestamp))
}

func (i *InvariableSP) Sign(pri ki.PRI) ([]byte, *errors.Error) {
	signBuilder := i.getSignBuilder()
	signStr, err := signBuilder.Sign(pri)
	if err != nil {
		return nil, err
	}
	return []byte(signStr), nil
}

func (i *InvariableSP) SignVerify(originPub ki.PUB) (bool, *errors.Error) {
	signBuilder := i.getSignBuilder()
	return signBuilder.Verify(originPub, string(i.Signature))
}

func (i *InvariableSP) GetType() Type {
	return i.Type
}

func (i *InvariableSP) GetVersion() Version {
	return i.Version
}

func (i *InvariableSP) GetVN() vn.ID {
	return i.Vn.Vn
}

func (i *InvariableSP) GetSP() sp.ID {
	return i.Sp
}

func (i *InvariableSP) GetSignature() []byte {
	return i.Signature
}

func (i *InvariableSP) GetTimestamp() int64 {
	return i.Timestamp
}
