package bk

import "github.com/hootuu/tome/ki"

type Ki struct {
	Addr ki.ADR `json:"addr"`
	Pub  []byte `json:"pub"`
}

func KiOf(guardian *ki.Ki) *Ki {
	return &Ki{
		Addr: guardian.ADR,
		Pub:  []byte(guardian.PUB),
	}
}

func (ki *Ki) ToString() string {
	return ki.Addr.S() + ":" + string(ki.Pub)
}
