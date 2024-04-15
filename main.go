package main

import (
	"net/http"

	"github.com/devhindo/katty/katty"
)

func main() {
	go func() {
		katty.Run()
	}()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from katty!"))
		})
		http.ListenAndServe(":8080", nil)
	}()

	select {}
}
