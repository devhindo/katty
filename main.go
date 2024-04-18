package main

import (
	"net/http"
	"time"
	"fmt"
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

	go func() {
		for {
			_, err := http.Get("https://katty-mqkx.onrender.com/")
			if err != nil {
				// Handle error
				fmt.Println(err)
			}
			// Wait for 14 minutes before sending the next request
			time.Sleep(14 * time.Minute)
		}
	}()

	select {}
}
