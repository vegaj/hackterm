package persistence

import (
  "fmt"
  "os"
  "strings"
)

var (
  filename string
)

func RequestFilename() string {
  fmt.Printf("Please, insert the path for the file to load:\n")
  var msg string
  fmt.Scanf("%s\n", &msg)
  return msg
}

func ObtainSelectionList(filename string) []string {
  var content string = ""
  content = readTextFile(filename)
  return strings.Split(content, "\n")
}

func readTextFile(filename string) string {

  f, err := os.Open(filename)
  check(err)
  defer f.Close()

  finfo, err := os.Stat(filename)
  check(err)
  size := finfo.Size()
  content := make([]byte, size)

  _ , err = f.Read(content)
  check(err)

  return string(content)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func handlePanic() {
  if r := recover(); r != nil {
    fmt.Println(r)
  }
}
