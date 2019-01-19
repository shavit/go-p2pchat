package p2p

import (
  "crypto/sha256"
)

type Block struct {
  index float64
  proof float64
  timestamp float64
  transactions []*Transaction
  previousHash string
}

// Hash creates a hash from a Block
func (b *Block) Hash() (hash string, err error) {
  h := sha256.New()
  h.Write([]byte("TODO: Replace this"))
  bArray := sha256.Sum256([]byte("some data to replace"))
  hash = string(bArray[0:len(bArray)])

  return hash, err
}
