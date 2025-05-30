package generator

import "sync/atomic"

type Generator struct {
	quoteIDCounter uint64
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) GetQuoteID() uint64 {
	atomic.AddUint64(&g.quoteIDCounter, 1)
	return g.quoteIDCounter
}
