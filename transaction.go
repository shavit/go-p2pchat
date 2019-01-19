package gop2p

type Transaction struct {
  amout int
  recipient string
  sender string
}

func NewTransaction() *Transaction {
  return new(Transaction)
}
