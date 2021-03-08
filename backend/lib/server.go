package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"
)

type Server struct {
	ApiClient  *http.Client
	ServerMux  *http.ServeMux
	MetricsMux *http.ServeMux
}

func (s *Server) Routes() {

	s.ServerMux.HandleFunc("/", s.methodHeadGet(s.handleIndexOr()))

	s.MetricsMux.HandleFunc("/", s.methodHeadGet(s.handleIndex()))
	s.MetricsMux.HandleFunc("/metrics", s.methodHeadGet(s.handleMetrics()))
}

func (s *Server) methodHeadGet(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodHead:
		case http.MethodGet:
		default:

			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Printf("code: %d method: %s path: %s", http.StatusMethodNotAllowed, r.Method, r.URL.Path)
			return
		}

		h(w, r)
	}
}

func (s *Server) handleIndexOr() http.HandlerFunc {

	dataID := regexp.MustCompile(`^/([0-9a-f]+)$`)

	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" {

			w.WriteHeader(http.StatusOK)
			log.Printf("code: %d method: %s path: %s", http.StatusOK, r.Method, r.URL.Path)
			return
		}

		if dataID.MatchString(r.URL.Path) {

			s.requestData(dataID.FindStringSubmatch(r.URL.Path)[1], w, r)
			log.Printf("code: %d method: %s path: %s", http.StatusOK, r.Method, r.URL.Path)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		log.Printf("code: %d method: %s path: %s", http.StatusBadRequest, r.Method, r.URL.Path)
	}
}

func (s *Server) handleIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

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

func (s *Server) requestData(dataID string, w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var question *Question
	var result *QuestionResult
	var question_err error
	var result_err error

	wg.Add(1)
	go func() {
		question, question_err = s.requestQuestion(dataID)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		result, result_err = s.requestResult(dataID)
		wg.Done()
	}()

	wg.Wait()

	if (question_err != nil) || (result_err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("code: %d method: %s path: %s", http.StatusInternalServerError, r.Method, r.URL.Path)
		return
	}

	answer := QuestionDesc{
		Question: question.Question,
		Results:  result.Results,
	}

	buf, err := json.Marshal(answer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("code: %d method: %s path: %s error: %v", http.StatusInternalServerError, r.Method, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
	log.Printf("code: %d method: %s path: %s", http.StatusOK, r.Method, r.URL.Path)
}

func (s *Server) requestQuestion(dataID string) (*Question, error) {

	var req *http.Request
	var err error
	data := Question{}
	entrypoint := "https://api.mentimeter.com"

	url := fmt.Sprintf("%s/questions/%s", entrypoint, dataID)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return &data, err
	}

	if err = s.ApiCall(req, &data); err != nil {
		log.Printf("error creating request %v", err)
		return &data, err
	}

	return &data, nil
}

func (s *Server) requestResult(dataID string) (*QuestionResult, error) {

	var req *http.Request
	var err error

	entrypoint := "https://api.mentimeter.com"
	url := fmt.Sprintf("%s/questions/%s/result", entrypoint, dataID)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}

	data := QuestionResult{}
	if err = s.ApiCall(req, &data); err != nil {
		log.Printf("error creating request %v", err)
		return nil, err
	}

	return &data, nil
}

func (s *Server) ApiCall(req *http.Request, obj interface{}) error {

	var resp *http.Response
	var err error
	var body []byte

	resp, err = s.ApiClient.Do(req)

	if err != nil {
		log.Printf("error creating request %v", err)
		return err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("error creating request %v", err)
		return err
	}

	if err = json.Unmarshal(body, obj); err != nil {
		log.Printf("error creating request %v", err)
		return err
	}

	return nil
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
