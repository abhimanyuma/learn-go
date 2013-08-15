package main

import "fmt"
import "strings"
//import "code.google.com/p/go-tour/pic"

type Coordinate struct{
	lat,long float64
	description string
}

func WordCount(s string) map[string]int {
	var words []string = strings.Fields(s)
	wordCount=make(map[string]int)
	for _,word := range words {
		c,ok = wordCount["word"]
		if ok {
			wordCount["word"]+=1
		} else {
			wordCount["word"]=1
		} 
	}
	return wordCount
}

var wikiEntry="George Lansbury, PC (22 February 1859 – 7 May 1940) was a British politician and social reformer who led the Labour Party from 1932 to 1935. Apart from a brief period of ministerial office during the Labour government of 1929–31, he spent his political life campaigning against established authority and vested interests, his main causes being the promotion of social justice, women's rights and world disarmament. Originally a radical Liberal, Lansbury converted to socialism in the early 1890s, and thereafter served his local community in the East End of London in numerous elective offices. His activities were underpinned by his Christian beliefs which, except for a short period of doubt, sustained him through his life. Elected to parliament in 1910, he resigned his seat in 1912 to campaign for women's suffrage, and was briefly imprisoned after publicly supporting militant action."

var places = map[string]Coordinate{
		"red-fort" : {28.655833, 77.240278,"A 17th century fort constructed by Mughals"},
		"taj-mahal": {27.174167, 78.042222,"Shaj Jahan constructed memorial in memory of his wife Mumtaz"},
		"kozhikode": {11.25, 75.77,"The city of spices"},
	}

func main(){
	
	colorcodes := map[string]string{"red":"#FF0000","blue":"#00FF00","green":"#0000FF"}

	fmt.Println(places["red-fort"])
	fmt.Println(colorcodes)
	fmt.Println(WordCount(wikiEntry))

}