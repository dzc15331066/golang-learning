package service

import (
	"github.com/unrolled/render"
	"net/http"
)

func timeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "time", nil)
	}
}
