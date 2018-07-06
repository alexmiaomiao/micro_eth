package types

type Header struct {
	Timestamp  int64
	Number     int64
	SelfHash   [32]byte
	ParentHash [32]byte
	Difficulty int
	Nonce      int
}

type Block struct {
	Header *Header
	Data   []byte
}

func NewBlock(data string) *Block {
	return &Block{&Header{}, []byte(data)}
}

func NewGenesisBlock() *Block {
	h := Header{
		Timestamp:  int64(0),
		Number:     int64(0),
		SelfHash:   [32]byte{},
		ParentHash: [32]byte{},
		Difficulty: 0,
		Nonce:      0,
	}

	b := NewBlock("Genesis Block")
	b.Header = &h

	return b
}
