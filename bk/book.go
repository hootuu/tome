package bk

import (
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type BID string

const (
	NilBID BID = ""
)

func (b BID) S() string {
	return string(b)
}

type Type string

func (t Type) S() string {
	return string(t)
}

type Invariable interface {
	GetType() Type
	GetVN() vn.ID
	GetSP() sp.ID
}

type Chain string

func (c Chain) S() string {
	return string(c)
}

type Linkable interface {
	Invariable
	GetChain() Chain
	GetInvariable() Invariable
	GetPrevious() BID
}

type Book interface {
	RegisterDataType(tpl interface{})
	Inscribe(data Invariable) (BID, *errors.Error)
}
