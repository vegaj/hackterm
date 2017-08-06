package cracking

import (
  "fmt"
  ch "github.com/vegaj/hackterm/choices"
)

const (
  MAX_ATTEMPTS int = 4
  DIFF_LENGTH int = -2
  PASSWORD_CORRECT int = -1
)

var (
  remainingAttempts int
  pset *ch.PosibilitySet
)

func compareCharacters(a,b string) int {
  var similarity int = 0

  for i := range a {
    if a[i] == b[i]{
      similarity++
    }
  }
  return similarity
}

func attempt(str string) int {
  if str == pset.Correct() {
    return PASSWORD_CORRECT
  }

  if len(str) != len(pset.Correct()){
    return DIFF_LENGTH
  }

  return compareCharacters(pset.Correct(), str)
}

func loginLoop() bool {
  var passwordAttempt string
  var similarity int
  for remainingAttempts > 0 {
    fmt.Printf("%d attempts remainings. Insert password:", remainingAttempts)
    _ , err := fmt.Scanf("%s\n", &passwordAttempt)
    if err != nil {
      panic(err)
    }
    similarity = attempt(passwordAttempt)
    switch similarity {
    case PASSWORD_CORRECT:
      return true
    case DIFF_LENGTH:
      similarity = 0
      break
    }
    fmt.Println("Attempt failed:",similarity,"characters in common.")
    remainingAttempts--
  }
  return false
}

func StartHacking (posibilities *ch.PosibilitySet){
  remainingAttempts = MAX_ATTEMPTS
  pset = posibilities

  fmt.Printf("Terminal secured with VaultSec.\n")
  var unlocked bool = loginLoop()
  if !unlocked {
    fmt.Println("Terminal locked: Too many attempts. Try again in a few.")
  } else {
    fmt.Println("\n\t--- Success. Access granted.--- ")
  }
}
