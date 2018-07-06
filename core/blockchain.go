package core

import (
	"fmt"
	"github.com/alexmiaomiao/micro_eth/consensus"
	"github.com/alexmiaomiao/micro_eth/consensus/pow"
	"github.com/alexmiaomiao/micro_eth/core/types"
)

type Blockchain struct {
	blocks []*types.Block
	engine consensus.Engine
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*types.Block{types.NewGenesisBlock()},
		engine: pow.NewProofOfWork(),
	}
}

func (bc *Blockchain) Blocks() []*types.Block {
	return bc.blocks
}

func (bc *Blockchain) AddBlock(block *types.Block) error {
	if bc.engine.Verify(bc, block) {
		bc.blocks = append(bc.blocks, block)
		return nil
	}
	return fmt.Errorf("bad block")
}

func (bc *Blockchain) CurrentHeader() *types.Header {
	return bc.blocks[len(bc.blocks)-1].Header
}
