package main

import (
	"fmt"

	"github.com/harshrastogiexe/lib/go/proxy"
)

func main() {
	r, _ := proxy.Magnates.Encyclopedia.GetResourcesList()
	for _, v := range r {
		fmt.Println(v.Name)
	}

}
