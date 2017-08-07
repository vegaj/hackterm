package main

import (
  "fmt"
  "os"
  "github.com/vegaj/hackterm/choices"
  "github.com/vegaj/hackterm/persistence"
  "github.com/vegaj/hackterm/cracking"
)

var (
  filename string
)

func reportPanic() {
  if r := recover(); r != nil {
    fmt.Printf("ERROR: <%s>.\n", r)
  }
}

func main(){
  defer reportPanic()
  args := os.Args[1:]

  if len(args) < 1 {
    filename = persistence.RequestFilename()
  } else {
    filename = args[1]
  }
  args = persistence.ObtainSelectionList(filename)
  fmt.Println(args)

  var pos *choices.PosibilitySet = choices.New(args)

  persistence.GenerateOutputFile(pos)

  cracking.StartHacking(pos)
}
