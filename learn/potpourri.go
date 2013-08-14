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

const (
	Big	= 1<<100
)



//Var to define variables //Now adding initial values
var alpha,beta,gamma int = 42,0,21

func Sqrt(x float64) float64 {
	z := float64(1)
	const times = 20
	for i:=0 ; i<times; i++ {
		z=z-(z*z-x)/(2*z)
	}
	return z
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // can't use v here, though
    return lim
}



//Shows named return in GO
func extendedGCD (a,b int) (int,int) {
	x,lastx := 0,1
	y,lasty := 1,0
	q := 0
	//In Go for is while
	for b>0 {
		q = a/b
		a,b = b, a%b
		x, lastx = (lastx - q*x), x
        y, lasty = (lasty - q*y), y 
	}
	return lastx,lasty
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
	// pi = pi + e Can't use coz pi is const
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
    //Fibonacci is quite easy in go
    a,b := 1,1
    for a<100 {
    	a,b = a+b,a
    	fmt.Println(a,b)
    }

    fmt.Println(extendedGCD(42,35))

    fmt.Println(pow(3, 2, 10),pow(3, 3, 20))

    fmt.Println(Sqrt(2))

    fmt.Println(needFloat(Big))

}