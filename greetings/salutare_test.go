package greetings
 import (
    "testing"
    "regexp"
)

func TestGreetingsName(t *testing.T) {
  name:="Gladys"
  want := regexp.MustCompile(`\b`+name+`\b`)
  msg, err := Greeting("Gladys")

  if !want.MatchString(msg) || err != nil {
      t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
  }
}

func TestGreetingsEmpty(t *testing.T) {
    msg, err := Greeting("")
    if msg != "" || err == nil {
      t.Fatalf(`Hello("") = %q, %v, want ""`, msg, err)
    }
}

func TestGreetingsEmptyReturnError(t *testing.T) {
  _, err := Greeting("")
  if err == nil {
    t.Fatalf(`No error has been signaled for empty passed parameter`)
  }
}
