package main

import "github.com/harshrastogi/eagle/cmd/server/route"

func main() {
	router := route.New()

	router.Run(":8080")
}
