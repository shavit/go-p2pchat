package p2p

type Block struct {
  index float64
  proof float64
  timestamp float64
  previousHash string
}

func NewBlock() *Block {
  return new(Block)
}
