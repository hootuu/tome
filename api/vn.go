package api

import (
	"github.com/hootuu/tome/bk"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type VNApi interface {
	PreCreate(vnID vn.ID) (vnM *vn.VN, vnBID bk.BID, err *errors.Error)
	ConfirmCreate(vnBID bk.BID, guardianPri ki.PRI) (vnLinkBID bk.BID, err *errors.Error)
}
