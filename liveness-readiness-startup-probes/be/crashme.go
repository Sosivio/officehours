package main

import (
	"fmt"
	"net/http"
)

func crashme(server http.Server) {
	// var pslice string
	// _, err := fmt.Println(pslice[10])
	// if err != nil {
	r := 0
	_, err := fmt.Println(100 / r)
	if err != nil {
		fmt.Println(err)
		server.Close()

	}
	// 	panic(err)
	// }
}
