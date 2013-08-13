package main

import (
  "fmt"
  "math/rand"
  "time"
  // "bytes"
  "math"
)

func main(){
  rand.Seed( time.Now().UTC().UnixNano())
  fmt.Println("My favourite number is",rand.Intn(43));
  fmt.Println("Now you have %g problems",math.Nextafter(2,3));
}