package main

import (
  	"fmt"
  	"math/rand"
  	"time"
  	// "bytes"
  	"math"
  	"math/cmplx"
)
var (
	ToBe 	bool 		= false
	MaxInt 	uint64 		= 1<<64 - 1
	z      	complex128 	= cmplx.Sqrt(-5 + 12i)
)

//Var to define variables //Now adding initial values
var alpha,beta,gamma int = 42,0,21



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
	//Omitting var by using
	iscorrect,iswrong,isascii := true,false,"Yes!" 

	const pi = math.Pi
	const e  = math.E
	rand.Seed( time.Now().UTC().UnixNano())
	fmt.Println("π = ",pi," e = ",e)
	fmt.Println("My constant = ",(math.Pi*math.Pi)/(math.E+math.E))
	fmt.Println("My favourite number is",rand.Intn(43));
	fmt.Printf("Now you have %d problems\n",
				sum(2,3));
	fmt.Println(swap("Halla","Bazinga"))
	fmt.Println(extendedGCD(2,3))
	fmt.Println(alpha,beta,gamma,iswrong,iscorrect,isascii)
	const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
}