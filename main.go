package main

import (
	"fmt"
	"io"
	"net/http"
)

func log(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	req.Body.Close()
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", log)
	http.ListenAndServe(":8090", nil)
}
