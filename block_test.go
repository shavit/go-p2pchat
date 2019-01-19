package p2p

import (
  "testing"
)

func testCreateNewBlock(t *testing.T){
  var b *Block
  b = NewBlock()

  if b == nil {
    t.Error("Found nil while expecting a Block")
  }
}
