package gop2p

import (
  "testing"
)

func TestCreateNewTransaction(t *testing.T){
  var tr *Transaction
  tr = NewTransaction()

  if tr == nil {
    t.Error("Found nil while expecting a Transaction")
  }
}
