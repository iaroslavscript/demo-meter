package main

import (
	"fmt"
	"log"
	"net/http"

	backend "github.com/iaroslavscript/demo-meter/backend/lib"
)

func listenAndServe(s *backend.Server, cfg *backend.Config) error {

	done := make(chan error)

	go func() {
		addr := fmt.Sprintf("%s:%d", cfg.MetricsBindAddr, cfg.MetricsBindPort)
		log.Printf("Start serving metrict at %s", addr)
		done <- http.ListenAndServe(addr, s.MetricsMux)
	}()

	go func() {

		addr := fmt.Sprintf("%s:%d", cfg.ServerBindAddr, cfg.ServerBindPort)
		log.Printf("Start serving at %s", addr)
		done <- http.ListenAndServe(addr, s.ServerMux)
	}()

	return <-done
}

func main() {
	log.Printf("INFO version:%s commit:%s buildtime:%s\n", BuildVersion, BuildCommit, BuildTime)

	cfg := backend.NewConfig()
	cfg.PopulateFromEnv()

	s := backend.NewServer()
	s.ApiEntrypoint = cfg.ApiEntrypoint
	s.Routes()

	log.Fatal(listenAndServe(s, cfg))
}
