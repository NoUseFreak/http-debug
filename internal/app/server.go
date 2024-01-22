package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func print(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	s := string(requestDump)
	fmt.Println("---")
	fmt.Println(s)
	fmt.Fprint(w, s)
}

type request struct {
	Host   string
	Method string
	URL    string
	Proto  string

	Headers map[string][]string

	Body string
}

func NewRequest(r *http.Request) (*request, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return &request{
		Host:    r.Host,
		Method:  r.Method,
		URL:     r.URL.String(),
		Proto:   r.Proto,
		Headers: r.Header,
		Body:    string(bodyBytes),
	}, nil
}

func printJson(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	s := string(requestDump)
	fmt.Fprint(w, s)

	req, err := NewRequest(r)
	if err != nil {
		fmt.Println(err)

		return
	}

	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
