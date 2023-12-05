package main

import (
	"my-go-project/web/vulns"
	"net/http"
)

func main() {
	http.HandleFunc("/xxe", vulns.Xxe)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
