package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!</h1>")
}

func main() {
	// Custom http server
	s := &http.Server{
		Addr:           ":8080\hello",
		Handler:        http.HandlerFunc(myHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: ", err.Error())
	}
}
