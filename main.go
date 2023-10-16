package main

import (
	"log"
	"net/http"
	"os"
	"gomicro/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	handl := handlers.hello(l)

	sm := http.NewServeMux()
	sm.Handle("/", handl)


	http.ListenAndServe(":9090", nil)
}