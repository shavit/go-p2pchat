package p2p

import (
  "errors"
  "fmt"
  "io"
  "log"
  "net"
  "os"
  "os/signal"
  "sync"
  "syscall"
)

// server contains the routing table
type server struct {
  done chan error
  conns map[uint32]*edge // Live connections
  ip net.IP
  ln net.Listener // Network listener
  port uint16
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
    port: 2400,
    RWMutex: new(sync.RWMutex),
  }
}

// Start starts the server to accept incoming connections, announce and
//  publish messages.
func (s *server) Start() (err error){
  var conn net.Conn
  var addr string

  // Exit gracefully
  var ch chan os.Signal = make(chan os.Signal)
  go signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

  var srvErr chan error = make(chan error, 1)
  go func() {
    addr = fmt.Sprintf("%s:%d", "127.0.0.1", s.port)
    s.Lock()
    s.ln, err = net.Listen("tcp", addr)
    s.Unlock()
    if err != nil {
      srvErr<- errors.New("Error listening to " + addr)
    }

    conn, err = s.ln.Accept()
    if err != nil {
      log.Println(err)
      srvErr<- errors.New("Error listening on port 2400")
    }

    println("Listening on", addr)
  }()

  // TODO: Check connections availability
  // ...

  L:
  for {
    select {
    case err = <-s.done:
      s.Close(err)
      break L
    case <- ch:
      s.Close(err)
      break L
    }

    // TODO: Handle the connection
    log.Println("New connection", conn)
  }

  return err
}

// Close sends a signal to the server, to exit gracefully.
//  The server will announce to the peers that it going to shutdown.
func (s *server) Close(err error) {
  if err != nil {
    log.Println("Exit with error", err)
  }

  s.Lock()
  defer s.Unlock()
  // TODO: Announce closing to all the clients
  // ..

  // Close the network connection
  err = s.ln.Close()
  log.Println("Closing the network connection.\n", err)

  // Exit the program
  os.Exit(0)
}
