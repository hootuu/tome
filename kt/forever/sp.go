package forever

import (
	"github.com/hootuu/tome/bk/bid"
	"github.com/hootuu/tome/sp"
)

type SP struct {
	VN bid.BID `bson:"vn" json:"vn"`
	ID sp.ID   `bson:"id" json:"id"`
}
