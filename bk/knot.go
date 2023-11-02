package bk

import (
	"fmt"
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

type Knot struct {
	Type       Type            `json:"type"`
	Version    Version         `json:"version"`
	Chain      Chain           `json:"chain"`
	Knot       KnotIDX         `json:"knot"`
	Invariable *InvariableLead `json:"invariable"`
	Previous   BID             `json:"previous"`
	Timestamp  int64           `json:"timestamp"`
	Signature  []byte          `json:"signature"`
}

func NewKnot(rope Rope, invariable Invariable) (*Knot, *errors.Error) {
	invariableBID, err := BuildBID(invariable)
	if err != nil {
		return nil, err
	}
	link := &Knot{
		Type:    KnotType,
		Version: KnotVersion,
		Chain:   rope.GetChain(),
		Knot:    rope.GetTailKnot().Next(),
		Invariable: &InvariableLead{
			Type:       invariable.GetType(),
			Version:    invariable.GetVersion(),
			Invariable: invariableBID,
		},
		Previous:  rope.GetTailInvariable(),
		Timestamp: invariable.GetTimestamp(),
		Signature: invariable.GetSignature(),
	}
	return link, nil
}

func (k *Knot) GetType() Type {
	return k.Type
}

func (k *Knot) GetVersion() Version {
	return k.Version
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
