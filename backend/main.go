package main

import (
	"log"
	"net/http"

	backend "github.com/iaroslavscript/demo-meter/backend/lib"
)

func listenAndServe(s *backend.Server) error {

	done := make(chan error)

	go func() {
		log.Printf("Start serving at 2112")
		done <- http.ListenAndServe("127.0.0.1:2112", s.MetricsMux)
	}()

	go func() {
		log.Printf("Start serving at 8080")
		done <- http.ListenAndServe("127.0.0.1:8080", s.ServerMux)
	}()

	return <-done
}

func main() {
	log.Printf("INFO version:%s commit:%s buildtime:%s\n", BuildVersion, BuildCommit, BuildTime)

	s := backend.NewServer()
	s.Routes()

	log.Fatal(listenAndServe(s))
}
