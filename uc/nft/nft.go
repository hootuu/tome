package nft

import (
	"encoding/json"
	"fmt"
	"github.com/hootuu/tome/bk/bid"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"regexp"
)

type NID = bid.BID

type Link = string
type Tag = string

func TagVerify(t Tag) *errors.Error {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._]{3,256}$`, t)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid nft.Tag(^[a-zA-Z0-9._]{3,256}$): %s", t))
	}
	return nil
}

func TagArrVerify(arr []Tag) *errors.Error {
	if len(arr) == 0 {
		return nil
	}
	for _, t := range arr {
		if err := TagVerify(t); err != nil {
			return err
		}
	}
	return nil
}

type Category string

func (c Category) S() string {
	return string(c)
}

func CategoryVerify(cStr string) *errors.Error {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._]{3,256}$`, cStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid nft.Category: %s", cStr))
	}
	return nil
}

type Token string

func (t Token) S() string {
	return string(t)
}

func TokenVerify(cStr string) *errors.Error {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._]{3,256}$`, cStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid nft.CategoyToken(^[a-zA-Z0-9._]{3,256}$): %s", cStr))
	}
	return nil
}

type Meta map[string]string

func NewMeta() Meta {
	return make(map[string]string)
}

func (c Meta) Put(code string, val string) *errors.Error {
	codeLen := len(code)
	if codeLen == 0 || codeLen > 60 {
		return errors.Verify("length of code must be in (0, 60]")
	}
	valLen := len(val)
	if valLen == 0 || valLen > 360 {
		return errors.Verify("length of value must be in (0, 360]")
	}
	c[code] = val
	return nil
}

func MetaToData(c Meta) ([]byte, *errors.Error) {
	if len(c) == 0 {
		c = NewMeta()
	}
	b, nErr := json.Marshal(c)
	if nErr != nil {
		return nil, errors.Verify("invalid meta")
	}
	if len(b) > 10240 {
		return nil, errors.Verify("invalid meta, too big.")
	}
	return b, nil
}

type NFT struct {
	VN          vn.ID    `bson:"vn" json:"vn"`
	Category    Category `bson:"category" json:"category"`
	Token       Token    `bson:"token" json:"token"`
	Tag         []Tag    `bson:"tag" json:"tag"`
	Link        Link     `bson:"link" json:"link,omitempty"`
	Title       string   `bson:"title" json:"title,omitempty"`
	Description string   `bson:"desc" json:"desc,omitempty"`
	Meta        Meta     `bson:"meta" json:"meta,omitempty"`
}

type Wrap struct {
	NID NID  `bson:"nid" json:"nid"`
	NFT *NFT `bson:"nft" json:"nft"`
}

func (m *NFT) Verify() *errors.Error {
	if err := vn.IDVerify(m.VN.S()); err != nil {
		return err
	}
	if err := CategoryVerify(m.Category.S()); err != nil {
		return err
	}
	if err := TokenVerify(m.Token.S()); err != nil {
		return err
	}
	if len(m.Tag) == 0 {
		m.Tag = []Tag{}
	} else {
		for _, t := range m.Tag {
			if err := TagVerify(t); err != nil {
				return err
			}
		}
	}
	if len(m.Meta) == 0 {
		m.Meta = NewMeta()
	}
	return nil
}
