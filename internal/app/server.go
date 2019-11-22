package app

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func serve(addr string) error {
	http.HandleFunc("/", print)
	fmt.Printf("Serving at %s\n", addr)

	return http.ListenAndServe(addr, nil)
}

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
