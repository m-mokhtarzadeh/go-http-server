package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is Version 3.\n")
}
func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /healthz request\n")
        w.WriteHeader(http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/healthz", getHealth)

	server_port := os.Getenv("SERVER_PORT")
	server_listen := ":" + server_port
	err := http.ListenAndServe(server_listen, nil)
	  if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
