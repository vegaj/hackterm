package main

import (
  "fmt"
  "os"
  "github.com/vegaj/hackterm/choices"
  "github.com/vegaj/hackterm/persistence"
  "github.com/vegaj/hackterm/cracking"
  )

func reportPanic() {
  if r := recover(); r != nil {
    fmt.Printf("ERROR: <%s>.\n", r)
  }
}

func main(){
  defer reportPanic()
  args := os.Args[1:]

  var pos *choices.PosibilitySet = choices.New(args)

  persistence.GenerateOutputFile(pos)

  //Hacking
  cracking.StartHacking(pos)
}
