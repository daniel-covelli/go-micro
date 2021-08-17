package main

import (
	"log"
	"net/http"
	"os"
	"working/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe("127.0.0.1:9090", sm)
}
