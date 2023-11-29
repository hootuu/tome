package bk

import (
	"fmt"
	"github.com/hootuu/tome/vn"
)

type Invariable interface {
	GetType() Type
	GetVersion() Version
	GetVN() vn.ID
	GetTimestamp() int64
	GetSignature() []byte
}

type InvariableLead struct {
	Type       Type    `json:"type"`
	Version    Version `json:"version"`
	Invariable BID     `json:"invariable"`
}

func (lead *InvariableLead) S() string {
	return fmt.Sprintf("type=%s;version=%d;invariable=%s", lead.Type, lead.Version, lead.Invariable)
}

type InvariableTemplate struct {
	Type      Type    `json:"type"`
	Version   Version `json:"version"`
	Vn        vn.ID   `json:"vn"`
	Timestamp int64   `json:"timestamp"`
	Signature []byte  `json:"signature"`
}

func (i *InvariableTemplate) GetType() Type {
	return i.Type
}

func (i *InvariableTemplate) GetVersion() Version {
	return i.Version
}

func (i *InvariableTemplate) GetVN() vn.ID {
	return i.Vn
}

func (i *InvariableTemplate) GetTimestamp() int64 {
	return i.Timestamp
}

func (i *InvariableTemplate) GetSignature() []byte {
	return i.Signature
}
