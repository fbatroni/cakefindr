package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	text := "Hello world!"
	if len(os.Getenv("RESPONSE_TEXT")) > 0 {
		text = os.Getenv("RESPONSE_TEXT")
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		w.Write([]byte("PONG"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		w.Write([]byte(text))
	})

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
