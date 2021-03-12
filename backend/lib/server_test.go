package backend

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestServerHandleIndex(t *testing.T) {
	tests := map[string]struct {
		inputMethod string
		inputPath   string
		wantCode    int
	}{
		"method head path /":    {inputMethod: "HEAD", inputPath: "/", wantCode: http.StatusOK},
		"method get path /":     {inputMethod: "GET", inputPath: "/", wantCode: http.StatusOK},
		"method post path /":    {inputMethod: "POST", inputPath: "/", wantCode: http.StatusMethodNotAllowed},
		"method put path /":     {inputMethod: "PUT", inputPath: "/", wantCode: http.StatusMethodNotAllowed},
		"method delete path /":  {inputMethod: "DELETE", inputPath: "/", wantCode: http.StatusMethodNotAllowed},
		"method connect path /": {inputMethod: "CONNECT", inputPath: "/", wantCode: http.StatusMethodNotAllowed},
		"method options path /": {inputMethod: "OPTIONS", inputPath: "/", wantCode: http.StatusMethodNotAllowed},
		"method trace path /":   {inputMethod: "TRACE", inputPath: "/", wantCode: http.StatusMethodNotAllowed},
		"method patch path /":   {inputMethod: "PATCH", inputPath: "/", wantCode: http.StatusMethodNotAllowed},

		"method get path /hello": {inputMethod: "GET", inputPath: "/hello", wantCode: http.StatusBadRequest},
	}

	log.SetOutput(ioutil.Discard)
	s := NewServer()
	s.Routes()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			req := httptest.NewRequest(tc.inputMethod, tc.inputPath, nil)
			w := httptest.NewRecorder()
			s.ServerMux.ServeHTTP(w, req)

			got := w.Result().StatusCode
			if !reflect.DeepEqual(tc.wantCode, got) {
				t.Errorf("expected: %v, got: %v", tc.wantCode, got)
			}
		})
	}
}

func TestApiClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	log.SetOutput(ioutil.Discard)
	s := NewServer()
	s.Routes()

	var req *http.Request
	var resp *http.Response
	var err error
	var body []byte

	req, err = http.NewRequest("GET", ts.URL, nil)

	if err != nil {
		t.Errorf("error creating request %v", err)
	}

	resp, err = s.ApiClient.Do(req)

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	got := string(body)
	want := "Hello, client\n"
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestApiCall(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "{\"message\": \"Hello World!\", \"x\": \"y\"}")
	}))
	defer ts.Close()

	log.SetOutput(ioutil.Discard)
	s := NewServer()
	s.Routes()

	var req *http.Request
	var err error

	type testObj struct {
		Message string `json:"message"`
	}

	req, err = http.NewRequest("GET", ts.URL, nil)

	if err != nil {
		t.Errorf("error creating request %v", err)
	}

	obj := testObj{}
	err = s.ApiCall(req, &obj)

	got := obj.Message
	want := "Hello World!"
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

/*
func TestRequestQuestion(t *testing.T) {
	dataID := "48d75c359ce4"
	entrypoint := "https://api.mentimeter.com"

	log.SetOutput(ioutil.Discard)
	s := NewServer()
	s.ApiEntrypoint = entrypoint
	s.Routes()

	data, err := s.requestQuestion(dataID)
	if err != nil {
		t.Errorf("error creating request %v", err)
	}

	t.Errorf("Got %v", data.Question)
}

func TestRequestQuestionResult(t *testing.T) {
	dataID := "48d75c359ce4"
	entrypoint := "https://api.mentimeter.com"

	log.SetOutput(ioutil.Discard)
	s := NewServer()
	s.ApiEntrypoint = entrypoint
	s.Routes()

	data, err := s.requestResult(dataID)
	if err != nil {
		t.Errorf("error creating request %v", err)
	}

	t.Errorf("Got %v", data.Results)
}
*/
