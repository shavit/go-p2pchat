package main

import (
  "os"
)

func main(){
  if len(os.Args) <= 1 {
    // TODO: Loop through options here
    os.Exit(1)
  }

  println("done")
}
