package gop2p

import (
  "testing"
)

func TestCreateHash(t *testing.T){
  var err error
  var b *Block
  var hash string

  b = new(Block)
  if b == nil {
    t.Error("Found nil while expecting a Block")
  }

  hash, err = b.Hash()
  if err != nil {
    t.Error(err)
  }

  if hash == "" {
    t.Error("Found empty string while expecting a hash")
  }
}
