package main

import (
  "os"
  "github.com/shavit/gop2p"
)

func main(){
  var err error

  if len(os.Args) <= 1 {
    // TODO: Loop through options here
    os.Exit(1)
  }

  var srv = gop2p.NewServer()
  println(srv)
  if err = srv.Start(); err != nil {
    panic(err)
  }

}
