package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", CalcEndpoint)
	http.ListenAndServe(":8088", nil)
}
