package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	port :=
		flag.Int("port", 8080, "Port to start the server")

	flag.Parse()

	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
