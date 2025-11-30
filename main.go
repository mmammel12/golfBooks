package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	fmt.Println("Server starting on :1889")
	err := http.ListenAndServe(":1889", mux)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Only match exact path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := home().Render(context.Background(), w)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
