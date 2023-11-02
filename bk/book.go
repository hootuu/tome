package bk

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
)

type Type string

const (
	GenesisType Type = "genesis"
	ChainType   Type = "chain"
)

func (t Type) S() string {
	return t
}

func TypeVerify(t Type) *errors.Error {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_.]{1,100}$`, t)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid bk.type: %s", t))
	}
	return nil
}

type Version int64

const (
	DefaultVersion Version = 1
)

func (v Version) S() string {
	return fmt.Sprintf("%d", v)
}

type Tag = string

type Linkable interface {
	GetType() Type
	GetVersion() Version
	GetChain() Chain
	GetPrevious() BID
	GetTimestamp() int64
	GetSignature() []byte
}

type Book interface {
	RegisterDataType(tpl interface{})
	Inscribe(data Invariable) (BID, *errors.Error)
}
