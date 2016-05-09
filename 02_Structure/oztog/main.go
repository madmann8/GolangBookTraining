package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	for _,oz := range os.Args[1:]{
		t, err:= strconv.ParseFloat(oz,64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s ounces to grams:%v\n",oz ,t*28.3495)
	}
}
