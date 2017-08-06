package choices

import "math/rand"
import "time"
import "strings"

import "fmt"

type PosibilitySet struct {
  posibilities []string
  passwordIndex int
}

var randomizer *rand.Rand = rand.New(rand.NewSource(time.Now().Unix()))

func  New(list []string) *PosibilitySet {
  var p PosibilitySet
  if len(list) == 0 {
    list = make([]string, 0)
    p = PosibilitySet {posibilities: list, passwordIndex: -1}
  } else {
    fmt.Println("len",len(list))
    p = PosibilitySet {posibilities: list, passwordIndex:  randomizer.Intn(len(list))}
  }
  return &p
}

func (p *PosibilitySet) String() string {
  if len(p.Set()) == 0 {
    return "Empty struct: No posibilities."
  }
  return "<" + strings.Join(p.posibilities, ", ")+ " with correct: " + p.posibilities[p.passwordIndex] + ">"
}

func (p *PosibilitySet) Set() []string {
  return p.posibilities
}

func (p *PosibilitySet) Correct() string {
  return p.posibilities[p.passwordIndex]
}

func (p *PosibilitySet) SetPosibilities(list []string){
  if list == nil || len(list) == 0 {
    return
  }

  p.posibilities = list
  p.passwordIndex = randomizer.Intn(len(list))
}
