package bkhelper

import (
	"github.com/hootuu/utils/errors"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/multiformats/go-multihash"
)

func BuildBID(data interface{}) (string, *errors.Error) {
	node, nErr := cbor.WrapObject(data, multihash.SHA2_256, -1)
	if nErr != nil {
		return "", errors.Sys("wrap cbor node failed: " + nErr.Error())
	}
	return node.Cid().String(), nil
}
