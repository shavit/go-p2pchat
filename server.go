package p2p

import (
  "errors"
  "io"
  "log"
  "net"
  "os"
  "sync"
)

// server contains the routing table
type server struct {
  ln net.Listener // Network listener
  done chan error

  conns map[uint32]*edge // Live connections

  *sync.RWMutex
}

// edge is a live connection to a server
type edge struct {
  conn *io.ReadWriteCloser
  msg chan Message // Message queue
}

// NewServer initialize a new empty server
func NewServer() (s *server) {
  return &server{
    conns: make(map[uint32]*edge),
    RWMutex: new(sync.RWMutex),
  }
}

// Start starts the server to accept incoming connections, announce and
//  publish messages.
func (s *server) Start() (err error){
  var conn net.Conn
  var addr string = "127.0.0.1:2400"

  s.ln, err = net.Listen("tcp", addr)
  if err != nil {
    return errors.New("Error listening to " + addr)
  }

  // Listen to the exit signal
  go func() {
    // Block and wait for a signal to exit the program
    s.Close(<-s.done) // It will wait for the channel before executing
    log.Println("\n\nSending exit notificaiton to all clients")
  }()

  // TODO: Check connections availability
  // ...

  println("Listening on", addr)
  for {
    conn, err = s.ln.Accept()
    if err != nil {
      log.Println(err)
      return errors.New("Error listening on port 2400")
    }

    // TODO: Handle the connection
    log.Println("New connection", conn)
  }

  return err
}

// Close sends a signal to the server, to exit gracefully.
//  The server will announce to the peers that it going to shutdown.
func (s *server) Close(err error) {
  log.Println("Closing the server.\n", err)

  // TODO: Announce closing to all the clients
  // ..

  // Close the network connection
  err = s.ln.Close()
  log.Println("Closing the network connection.\n", err)

  // Exit the program
  os.Exit(0)
}
