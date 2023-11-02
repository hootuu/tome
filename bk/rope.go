package bk

type Rope interface {
	GetChain() Chain
	GetTailKnotIDX() KnotIDX
	GetTailKnotBID() BID
}
