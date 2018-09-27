package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
			<html>
			<head>
			<title>Hello, World</title>
			</head>
			<body>
			<h1>
			<b>Hello, World!</b>
			</h1>
			</body>
			</html>`,
	)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}