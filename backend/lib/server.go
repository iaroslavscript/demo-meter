package backend

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	ApiClient  *http.Client
	ServerMux  *http.ServeMux
	MetricsMux *http.ServeMux
}

func (s *Server) Routes() {
	s.ServerMux.HandleFunc("/", s.handleIndex())

	s.MetricsMux.HandleFunc("/", s.handleIndex())
	s.MetricsMux.HandleFunc("/metrics", s.handleMetrics())
}

func (s *Server) handleIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodHead:
		case http.MethodGet:
		default:

			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Printf("code: %d method: %s path: %s", http.StatusMethodNotAllowed, r.Method, r.URL.Path)
			return
		}

		if r.URL.Path == "/" {

			w.WriteHeader(http.StatusOK)
			log.Printf("code: %d method: %s path: %s", http.StatusOK, r.Method, r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		log.Printf("code: %d method: %s path: %s", http.StatusBadRequest, r.Method, r.URL.Path)
	}
}

func (s *Server) handleMetrics() http.HandlerFunc {

	msg := []byte("THE METRICS WILL BE HERE")

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(msg)
	}
}

func (s *Server) ApiCall(req *http.Request) (*http.Response, error) {

	return s.ApiClient.Do(req)
}

func NewServer() *Server {

	defaultTransport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	defaultClient := &http.Client{
		Transport: defaultTransport,
	}

	s := Server{
		ApiClient:  defaultClient,
		ServerMux:  http.NewServeMux(),
		MetricsMux: http.NewServeMux(),
	}

	return &s
}
