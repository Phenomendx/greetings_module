package greetings

import (
 "fmt"
 "rsc.io/quote"
 "errors"
 "math/rand"
 "time"
)

func Hello() {
    fmt.Println(quote.Go())
}


func init() {
  rand.Seed(time.Now().UnixNano())
}



func randomFormat() string {
  formats := []string {
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}


func Greeting(name string) (string, error) {

  if name  == ""  {
    return "", errors.New("empty name")
  }

   message := fmt.Sprintf(randomFormat(), name)
   return message, nil
}
