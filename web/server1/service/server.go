package service

import (
	"fmt"
	"net/http"
)

type MyServer struct {
	mux *http.ServeMux
}

// set the running address
func (sr *MyServer) Run(addr string) {
	fmt.Println("listening at port "+addr)
	http.ListenAndServe(addr, sr)
	
}

func (sr *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if sr.mux == nil {
		sr.mux = http.DefaultServeMux
	}
	sr.mux.ServeHTTP(w, r)
	return
}

// NewServer configures and returns a Server.
func NewServer() *MyServer {
	sr := &MyServer{}
	sr.mux = http.NewServeMux()
	sr.mux.HandleFunc("/", sayhelloName)
	return sr
}

// define a handle function
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"]
	if len(name) == 0 {
		fmt.Fprintf(w, "Hello\n")
	} else {
		fmt.Fprintf(w, "Hello "+name[0]+"\n")
	}

}
