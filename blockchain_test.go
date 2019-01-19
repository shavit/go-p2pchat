package gop2p

import (
  "testing"
)

func TestCreateNewBlockchain(t *testing.T) {
  var bc *Blockchain

  bc = NewBlockchain()
  if bc == nil {
    t.Error("Found nil while expecting Blockchain")
  }
}
