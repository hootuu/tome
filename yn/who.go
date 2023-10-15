package yn

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/utils/errors"
)

type Who struct {
	Who ki.ADR `bson:"who" json:"who"`
	Ref *Ref   `bson:"ref" json:"ref"`
}

func NewWho(whoStr string, refCode string, refRef string) (*Who, *errors.Error) {
	who, err := ki.ADROf(whoStr)
	if err != nil {
		return nil, err
	}
	ref, err := RefOf(refCode, refRef)
	if err != nil {
		return nil, err
	}
	return &Who{
		Who: who,
		Ref: ref,
	}, nil
}
