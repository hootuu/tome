package bk

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/tome/yn"
	"github.com/hootuu/utils/errors"
	"strings"
	"time"
)

const (
	InvariableYinType    Type = "yin"
	InvariableYinVersion      = DefaultVersion
)

type InvariableYin struct {
	Type      Type        `json:"type"`
	Version   Version     `json:"version"`
	Vn        *VNLead     `json:"vn"`
	Sp        *SPLead     `json:"sp"`
	Tag       []Tag       `json:"tag"`
	Who       *yn.Who     `json:"who"`
	When      yn.When     `json:"when"`
	Act       yn.Act      `json:"act"`
	What      *yn.What    `json:"what"`
	Exp       *yn.Expense `json:"exp"`
	Title     yn.Title    `json:"title"`
	Data      []byte      `json:"data"`
	Timestamp int64       `json:"timestamp"`
	Signature []byte      `json:"signature"`
}

func NewInvariableYin(yin *yn.Yin, vnLead *VNLead, spLead *SPLead, vnGuardianPri ki.PRI) (*InvariableYin, *errors.Error) {
	if err := yin.Verify(); err != nil {
		return nil, err
	}
	i := &InvariableYin{
		Type:      InvariableYinType,
		Version:   InvariableYinVersion,
		Vn:        vnLead,
		Sp:        spLead,
		Tag:       yin.Tag,
		Who:       yin.Who,
		When:      yin.When,
		Act:       yin.Act,
		What:      yin.What,
		Exp:       yin.Exp,
		Title:     yin.Title,
		Data:      nil,
		Timestamp: time.Now().UnixMilli(),
		Signature: nil,
	}
	var err *errors.Error
	i.Data, err = yn.CtxToData(yin.Ctx)
	if err != nil {
		return nil, err
	}

	i.Signature, err = i.Sign(vnGuardianPri)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (i *InvariableYin) getSignBuilder() *SignBuilder {
	return NewSignBuilder().
		Add("type", i.Type.S()).
		Add("version", i.Version.S()).
		Add("vn", i.Vn.S()).
		Add("sp", i.Sp.S()).
		Add("tag", strings.Join(i.Tag, ",")).
		Add("who", i.Who.S()).
		Add("when", i.When.S()).
		Add("act", i.Act.S()).
		Add("what", i.What.S()).
		Add("exp", i.Exp.S()).
		Add("title", i.Title.S()).
		Add("timestamp", fmt.Sprintf("%d", i.Timestamp)).
		Add("data", string(i.Data))
}

func (i *InvariableYin) Sign(pri ki.PRI) ([]byte, *errors.Error) {
	signBuilder := i.getSignBuilder()
	signStr, err := signBuilder.Sign(pri)
	if err != nil {
		return nil, err
	}
	return []byte(signStr), nil
}

func (i *InvariableYin) SignVerify(originPub ki.PUB) (bool, *errors.Error) {
	signBuilder := i.getSignBuilder()
	return signBuilder.Verify(originPub, string(i.Signature))
}

func (i *InvariableYin) GetType() Type {
	return i.Type
}

func (i *InvariableYin) GetVersion() Version {
	return i.Version
}

func (i *InvariableYin) GetVN() vn.ID {
	return i.Vn.Vn
}

func (i *InvariableYin) GetSignature() []byte {
	return i.Signature
}

func (i *InvariableYin) GetTimestamp() int64 {
	return i.Timestamp
}
