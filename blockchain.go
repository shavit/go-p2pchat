package gop2p

type Blockchain struct {
  chain []*Block
  currentTransactions []*Transaction
}

func NewBlockchain() (bc *Blockchain) {
  bc = new(Blockchain)
  bc.AddBlock(10, "1")

  return bc
}

func (bc *Blockchain) AddBlock(proof float64, previousHash string) (err error) {
  block := &Block{
    proof: 10,
    previousHash: "1",
  }
  bc.chain = append(bc.chain, block)

  return err
}
