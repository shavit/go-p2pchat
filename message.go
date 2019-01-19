package gop2p


// MessageType represents a message type between the peers:
//  1. Join
//  2. Leave
//  3. Copy the routing table
//  4. Find the closest
type MessageType byte

const (
  _ MessageType = iota
  JOIN
  LEAVE
  COPY_ROUTES
  FIND_CLOSEST
)

// Message contains information about a message between clients in the network
type Message struct {
  msgType MessageType
}
