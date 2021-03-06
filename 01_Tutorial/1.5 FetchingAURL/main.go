package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://")!= true {
			url= "http://"+ url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stterr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if strings.HasPrefix(url, "http://")!= true {
			url= "http:// "+ url
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Println(resp.Status)
	}
}