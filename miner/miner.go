package miner

import (
	"github.com/alexmiaomiao/micro_eth/consensus"
	"github.com/alexmiaomiao/micro_eth/consensus/pow"
	"github.com/alexmiaomiao/micro_eth/core"
	"github.com/alexmiaomiao/micro_eth/core/types"
)

type Miner struct {
	coinbase string
	bc       *core.Blockchain
	engine   consensus.Engine
}

func New(coinbase string, bc *core.Blockchain) *Miner {
	m := &Miner{
		coinbase: coinbase,
		bc:       bc,
		engine:   pow.NewProofOfWork(),
	}
	return m
}

func (self *Miner) Start(data string) {
	block := types.NewBlock(self.coinbase + "||"  + data)
	self.engine.Prepare(self.bc, block)
	self.engine.Seal(self.bc, block)
	self.bc.AddBlock(block)
}
