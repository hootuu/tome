package bid

import (
	"github.com/hootuu/utils/errors"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/multiformats/go-multihash"
)

type BID string

const (
	NilBID     BID = ""
	GenesisBID BID = "$"
)

func (b BID) S() string {
	return string(b)
}

func Build(data interface{}) (BID, *errors.Error) {
	node, nErr := cbor.WrapObject(data, multihash.SHA2_256, -1)
	if nErr != nil {
		return NilBID, errors.Sys("wrap cbor node failed: " + nErr.Error())
	}
	return BID(node.Cid().String()), nil
}
