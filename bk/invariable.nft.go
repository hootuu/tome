package bk

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/uc/nft"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"strings"
	"time"
)

const (
	InvariableNFTType    Type = "hotu.nft"
	InvariableNFTVersion      = DefaultVersion
)

type InvariableNFT struct {
	Type        Type         `json:"type"`
	Version     Version      `json:"version"`
	Vn          *VNLead      `json:"vn"`
	Category    nft.Category `json:"category"`
	Token       nft.Token    `json:"token"`
	Tag         []nft.Tag    `json:"tag"`
	Link        nft.Link     `json:"link"`
	Title       string       `json:"title"`
	Description string       `json:"desc"`
	Meta        nft.Meta     `json:"meta"`
	Timestamp   int64        `json:"timestamp"`
	Signature   []byte       `json:"signature"`
}

func NewInvariableNFT(
	vnLead *VNLead,
	nftM *nft.NFT,
	vnGuardianPri ki.PRI,
) (*InvariableNFT, *errors.Error) {
	if err := nftM.Verify(); err != nil {
		return nil, err
	}
	i := &InvariableNFT{
		Type:        InvariableNFTType,
		Version:     InvariableNFTVersion,
		Vn:          vnLead,
		Category:    nftM.Category,
		Token:       nftM.Token,
		Tag:         nftM.Tag,
		Link:        nftM.Link,
		Title:       nftM.Title,
		Description: nftM.Description,
		Meta:        nftM.Meta,
		Timestamp:   time.Now().UnixMilli(),
		Signature:   nil,
	}
	var err *errors.Error
	i.Signature, err = i.Sign(vnGuardianPri)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (i *InvariableNFT) getSignBuilder() *SignBuilder {
	metaBytes, _ := nft.MetaToData(i.Meta)
	metaStr := ""
	if metaBytes != nil {
		metaStr = string(metaBytes)
	}
	return NewSignBuilder().
		Add("type", i.Type.S()).
		Add("version", i.Version.S()).
		Add("vn", i.Vn.S()).
		Add("category", i.Category.S()).
		Add("token", i.Token.S()).
		Add("tag", strings.Join(i.Tag, ",")).
		Add("link", i.Link).
		Add("title", i.Title).
		Add("desc", i.Description).
		Add("meta", metaStr).
		Add("timestamp", fmt.Sprintf("%d", i.Timestamp))
}

func (i *InvariableNFT) Sign(pri ki.PRI) ([]byte, *errors.Error) {
	signBuilder := i.getSignBuilder()
	signStr, err := signBuilder.Sign(pri)
	if err != nil {
		return nil, err
	}
	return []byte(signStr), nil
}

func (i *InvariableNFT) SignVerify(vnGuardianPub ki.PUB) (bool, *errors.Error) {
	signBuilder := i.getSignBuilder()
	return signBuilder.Verify(vnGuardianPub, string(i.Signature))
}

func (i *InvariableNFT) GetType() Type {
	return i.Type
}

func (i *InvariableNFT) GetVersion() Version {
	return i.Version
}

func (i *InvariableNFT) GetVN() vn.ID {
	return i.Vn.Vn
}

func (i *InvariableNFT) GetSignature() []byte {
	return i.Signature
}

func (i *InvariableNFT) GetTimestamp() int64 {
	return i.Timestamp
}
