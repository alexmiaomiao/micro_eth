package consensus

import (
	"github.com/alexmiaomiao/micro_eth/core/types"
)

type ChainReader interface {
	CurrentHeader() *types.Header
}

type Engine interface {
	Verify(chain ChainReader, block *types.Block) bool
	Prepare(chain ChainReader, block *types.Block)
	Seal(chain ChainReader, block *types.Block)
}
