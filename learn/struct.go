package main

import "fmt"
import "code.google.com/p/go-tour/pic"

type Coordinate struct{
	lat,long float64
	description string
}

func Pic(dx, dy int) [][]uint8 {
    
    arr := make ([][]uint8,dy)
    for y:=0;y<dy;y++ {
    	line := make ([]uint8,dx)
    	for x:= 0;x<dx;x++ {
    		line[x]=x*y
    	}
    	arr[y]=line
    }
    return arr
        
}


func main(){
	pic.Show(Pic)
}