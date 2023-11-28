package kt

import (
	"encoding/json"
	"github.com/hootuu/tome/bk/bid"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type Invariable interface {
	GetType() Type
	GetVersion() Version
	GetVN() vn.ID
	GetSignature() *Signature
	SetSignature(signature *Signature)
	Signing() *Signing
}

func InvariableSign(inv Invariable, pri ki.PRI) *errors.Error {
	sign := NewSignature()
	err := sign.sign(inv, pri)
	if err != nil {
		return err
	}
	return nil
}

func InvariableVerify(inv Invariable) *errors.Error {
	return inv.GetSignature().verify(inv)
}

func InvariableMarshal(inv Invariable) ([]byte, *errors.Error) {
	bytes, nErr := json.Marshal(inv)
	if nErr != nil {
		return nil, errors.Verify("invalid invariable: " + nErr.Error())
	}
	return bytes, nil
}

type Lead struct {
	Type       Type    `json:"t"`
	Version    Version `json:"v"`
	Invariable bid.BID `json:"f"`
}

func (l *Lead) S() string {
	return l.Type.S() + ":" + l.Version.S() + "/" + l.Invariable.S()
}

type Template struct {
	Type      Type       `json:"t"`
	Version   Version    `json:"v"`
	Vn        vn.ID      `json:"vn"`
	Signature *Signature `json:"s"`
}

func InvariableOf(data []byte) (*Template, *errors.Error) {
	var tpl Template
	nErr := json.Unmarshal(data, &tpl)
	if nErr != nil {
		return nil, errors.Verify("invalid invariable: " + nErr.Error())
	}
	return &tpl, nil
}

func (t *Template) GetType() Type {
	return t.Type
}

func (t *Template) GetVersion() Version {
	return t.Version
}

func (t *Template) GetVN() vn.ID {
	return t.Vn
}

func (t *Template) GetSignature() *Signature {
	return t.Signature
}

func (t *Template) SetSignature(signature *Signature) {
	t.Signature = signature
}

func (t *Template) Signing() *Signing {
	return NewSigning()
}
