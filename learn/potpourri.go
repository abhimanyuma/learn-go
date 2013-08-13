package main

import (
  "fmt"
  "math/rand"
  "time"
  // "bytes"
  "math"
)
//Var to define variables //Now adding initial values
var alpha,beta,gamma int = 42,0,21
var iscorrect,iswrong,isascii = true,false,"Yes!" 


//Shows named return in GO
func extendedGCD (a,b int) (x,y int) {
	x = 0
	y = 1
	lastx := 1
	lasty := 0
	x = x-lastx
	y = y-lasty
	return 
}

//Shows multiple return in Go, well actually
//there is a single return with type (..,..)
func swap (left,right string) (string,string){
	return right,left
}

//Simple func, look how they turned the function
//definition around, 
func sum(a,b int) int {
	return (a+b)
}


func main(){
  	rand.Seed( time.Now().UTC().UnixNano())
  	fmt.Println("π = ",math.Pi," e = ",math.E)
  	fmt.Println("My constant = ",(math.Pi*math.Pi)/(math.E+math.E))
  	fmt.Println("My favourite number is",rand.Intn(43));
  	fmt.Printf("Now you have %d problems\n",
  	          	sum(2,3));
  	fmt.Println(swap("Halla","Bazinga"))
  	fmt.Println(extendedGCD(2,3))
  	fmt.Println(alpha,beta,gamma,iswrong,iscorrect,isascii)
}