package api

import (
	"github.com/hootuu/tome/bk"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type CoinApi interface {
	Issue(
		vnID vn.ID,
		coin uc.Code,
		wei uc.WEI,
		infinite bool,
		issuance int64,
		vnGuardianPri ki.PRI,
	) (*uc.Coin, bk.BID, *errors.Error)

	Airdrop(
		coin uc.Code,
		to ki.ADR,
		amount *uc.Amount,
		vnGuardianPri ki.PRI,
	) (*uc.Alter, bk.BID, *errors.Error)
}
