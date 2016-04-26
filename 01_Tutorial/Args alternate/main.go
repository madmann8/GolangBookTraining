// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
	for n,m := range os.Args {
		fmt.Println(n,m)
	}
}

//log time on this later