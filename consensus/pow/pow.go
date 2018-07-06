package pow

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/alexmiaomiao/micro_eth/common"
	"github.com/alexmiaomiao/micro_eth/consensus"
	"github.com/alexmiaomiao/micro_eth/core/types"
)

var (
	maxNonce   = math.MaxInt64
	difficulty = 20
)

type ProofOfWork struct {
	target *big.Int
}

func NewProofOfWork() *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	return &ProofOfWork{target}
}

func prepareData(block *types.Block, nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			block.Header.ParentHash[:],
			block.Data,
			common.IntToHex(block.Header.Timestamp),
			common.IntToHex(block.Header.Number),
			common.IntToHex(int64(block.Header.Difficulty)),
			common.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (self *ProofOfWork) Prepare(chain consensus.ChainReader, block *types.Block) {
	h := chain.CurrentHeader()
	block.Header.Timestamp = time.Now().Unix()
	block.Header.Difficulty = difficulty
	block.Header.ParentHash = h.SelfHash
	block.Header.Number = h.Number + 1
}

func (self *ProofOfWork) Seal(chain consensus.ChainReader, block *types.Block) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", block.Data)
	for nonce < maxNonce {
		data := prepareData(block, nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(self.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n")

	block.Header.Nonce = nonce
	block.Header.SelfHash = hash
}

func (self *ProofOfWork) Verify(chain consensus.ChainReader, block *types.Block) bool {
	var hashInt big.Int

	data := prepareData(block, block.Header.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(self.target) == -1

	h := chain.CurrentHeader()
	isValid = isValid && (block.Header.ParentHash == h.SelfHash) && (h.Number + 1 == block.Header.Number)

	return isValid
}
