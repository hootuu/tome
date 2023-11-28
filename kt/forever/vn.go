package forever

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
)

type VN struct {
	ID  vn.ID  `bson:"id" json:"id"`
	ADR ki.ADR `bson:"adr" json:"adr"`
	PUB ki.PUB `bson:"pub" json:"pub"`
}
