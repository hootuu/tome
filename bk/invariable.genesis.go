package bk

import (
	"github.com/hootuu/tome/vn"
)

const (
	InvariableGenesisType    Type = "hotu.chain.genesis"
	InvariableGenesisVersion      = DefaultVersion
)

type InvariableGenesis struct {
	Type    Type    `json:"type"`
	Version Version `json:"version"`
	Chain   Chain   `json:"chain"`
	Vn      vn.ID   `json:"vn"`
}

func NewInvariableGenesis(vnID vn.ID, chain Chain) *InvariableGenesis {
	i := &InvariableGenesis{
		Type:    InvariableVNType,
		Version: InvariableVNVersion,
		Chain:   chain,
		Vn:      vnID,
	}
	return i
}

func (i *InvariableGenesis) GetType() Type {
	return i.Type
}

func (i *InvariableGenesis) GetVersion() Version {
	return i.Version
}

func (i *InvariableGenesis) GetVN() vn.ID {
	return i.Vn
}

func (i *InvariableGenesis) GetTimestamp() int64 {
	return 0
}

func (i *InvariableGenesis) GetSignature() []byte {
	return []byte{0, 0, 0, 0, 0, 0, 0, 0}
}
