package service

import (
	"github.com/unrolled/render"
	"net/http"
	"time"
)

func apiTimeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ServerTime string `json:"server_time"`
		}{time.Now().String()})
	}
}
