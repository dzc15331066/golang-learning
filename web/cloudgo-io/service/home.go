package service

import (
	"fmt"
	"github.com/unrolled/render"
	"net/http"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("method:", req.Method)
		if req.Method == "GET" {
			formatter.HTML(w, http.StatusOK, "index", nil)
		} else if req.Method == "POST" {
			err := req.ParseForm()
			if err != nil {
				panic(err)
			}
			mf := MyForm{req.Form.Get("id"), req.Form.Get("content")}
			formatter.HTML(w, http.StatusOK, "form", mf)
		}
	}
}
