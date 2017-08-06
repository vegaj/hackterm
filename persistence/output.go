package persistence

import (
  "fmt"
  "github.com/vegaj/hackterm/choices"
  rnd "math/rand"
  tim "time"
  "io/ioutil"
)

var LINE_LENGTH int = 32
var random *rnd.Rand = rnd.New(rnd.NewSource(tim.Now().Unix()))


var FILENAME string = "coredump.hack"
var chars string = "<>[]{}!#~/&+-':;.,"

func nextCharacter() byte {
  return chars[random.Intn(len(chars))]
}

func generateSizedJunkString(length int) string {
  var content []byte = make([]byte, length)

  for i := range content {
      content[i] = nextCharacter()
  }

  return string(content)
}

func generateLeftJunk(length int) string {
  if (length <= 0) {return ""}
  length = random.Intn(length)
  return generateSizedJunkString(length)
}

func generateRightJunk(length int) string {
  length = LINE_LENGTH - length
  if (length <= 0) {return ""}
  return generateSizedJunkString(length)
}

func generateLine(key string) string{
  var left string = generateLeftJunk(LINE_LENGTH - len(key))
  var right string = generateRightJunk(len(left) + len(key))
  return left + key + right
}

func generateForEachKey(keys []string) string {
  if keys == nil || len(keys) == 0 {
    panic("No keys assigned")
  }
  res := generateLine(keys[0])

  var i int = 1
  for i < len(keys) {
    res = fmt.Sprintf("%s\n%s", res, generateLine(keys[i]))
    i++
  }

  return res
}

func PrintInto(filename string, data string) error {
  return ioutil.WriteFile(filename, []byte(data), 0644)
}

func GenerateOutputFile(data *choices.PosibilitySet){
  var content string
  content = generateForEachKey(data.Set())
  err := PrintInto(FILENAME, content)
  if (err != nil) {
    panic(err)
  }
}
