package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ideahitme/k8s-load-test/controller"
)

func main() {
	go startServer()
	ctx := context.WithValue(context.Background(), "users", 10)
	controller.Run(ctx, "http://localhost:3001/")

	//start a server to test
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK!")
	})
	http.ListenAndServe(":3001", nil)
}
