package bk

import (
	"fmt"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type KnotIDX int64

const (
	GenesisKnotIDX KnotIDX = 0
)

func (knotIdx KnotIDX) S() string {
	return fmt.Sprintf("%d", knotIdx)
}

func (knotIdx KnotIDX) Next() KnotIDX {
	return knotIdx + 1
}

func (knotIdx KnotIDX) Previous() KnotIDX {
	return knotIdx - 1
}

const (
	KnotType    Type = "knot"
	KnotVersion      = DefaultVersion
)

func BuildGenesis() Invariable {
	return &InvariableTemplate{
		Type:      "",
		Version:   0,
		Vn:        "",
		Timestamp: 0,
		Signature: nil,
	}
}

type Knot struct {
	Type       Type            `json:"type"`
	Version    Version         `json:"version"`
	Vn         vn.ID           `json:"vn"`
	Chain      Chain           `json:"chain"`
	Knot       KnotIDX         `json:"knot"`
	Invariable *InvariableLead `json:"invariable"`
	Previous   BID             `json:"previous"`
	Timestamp  int64           `json:"timestamp"`
	Signature  []byte          `json:"signature"`
}

func NewKnot(rope Rope, tie *Tie) (*Knot, *errors.Error) {
	link := &Knot{
		Type:       KnotType,
		Version:    KnotVersion,
		Chain:      rope.GetChain(),
		Knot:       rope.GetTailKnotIDX().Next(),
		Invariable: tie.Invariable,
		Previous:   rope.GetTailKnotBID(),
		Timestamp:  tie.Timestamp,
		Signature:  tie.Signature,
	}
	return link, nil
}

func (k *Knot) GetType() Type {
	return k.Type
}

func (k *Knot) GetVersion() Version {
	return k.Version
}

func (k *Knot) GetVN() vn.ID {
	return k.Vn
}

func (k *Knot) GetChain() Chain {
	return k.Chain
}

func (k *Knot) GetPrevious() BID {
	return k.Previous
}

func (k *Knot) GetTimestamp() int64 {
	return k.Timestamp
}

func (k *Knot) GetSignature() []byte {
	return k.Signature
}
