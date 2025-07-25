package main

import (
	"fmt"
	"io"
	"net/http"
)

var msgCounter int64

func log(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	msgCounter++
	fmt.Println("received messages:", msgCounter)
	err = req.Body.Close()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", log)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}
