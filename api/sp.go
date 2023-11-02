package api

import (
	"github.com/hootuu/tome/bk"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type SPApi interface {
	Create(vnID vn.ID, link []sp.ID, vnGuardianPri ki.PRI) (*sp.SP, bk.BID, bk.BID, *errors.Error)
}
