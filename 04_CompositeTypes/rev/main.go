package main

import "fmt"

func main(){
	s:= [...]int{1,2,3}
	rev(s[:])
	fmt.Println(s)
}

func rev(input []int) []int {
	var j int
	l:=len(input)-1
	for i:=0; i<l;i++ {
		j= input[len(input)-i]
		input=append(input, j)
		i++
	}
	return input[:l]
}
