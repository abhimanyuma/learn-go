package main

import (
  "fmt"
  "math/rand"
  "time"
  // "bytes"
  "math"
)
func sum(a int,b int) int {
	return (a+b)
}


func main(){
  rand.Seed( time.Now().UTC().UnixNano())
  fmt.Println("π = ",math.Pi," e = ",math.E)
  fmt.Println("My constant = ",(math.Pi*math.Pi)/(math.E+math.E))
  fmt.Println("My favourite number is",rand.Intn(43));
  fmt.Printf("Now you have %d problems\n",
  	          sum(2,3));
}