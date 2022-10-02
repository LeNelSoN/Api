package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	var (
		route      string
		controller []string
	)
	route = r.URL.Path
	route = strings.Trim(route, "/")
	route = strings.ToLower(route)

	controller = strings.Split(route, "/")
	fmt.Println(controller)
	// fmt.Fprintf(w,controller[0])
}

