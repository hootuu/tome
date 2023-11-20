package bk

import (
	"github.com/hootuu/tome/bk/bkhelper"
	"github.com/hootuu/utils/errors"
)

type BID string

const (
	NilBID     BID = ""
	GenesisBID BID = "$"
)

func (b BID) S() string {
	return string(b)
}

func BuildBID(data Invariable) (BID, *errors.Error) {
	bidStr, err := bkhelper.BuildBID(data)
	if err != nil {
		return NilBID, err
	}
	return BID(bidStr), nil
}
