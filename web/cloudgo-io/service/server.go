package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"html/template"
	"net/http"
	"os"
)

type MyForm struct {
	ID      string
	Content string
}

var tmpl = template.New("root")

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {

	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root

			//fmt.Println(root)
		}
	}
	mx.Handle("/", homeHandler(formatter))
	mx.Handle("/time", timeHandler(formatter))
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(webRoot+"/assets/"))))
	mx.Handle("/api/time", apiTimeHandler(formatter))
	mx.Handle("/api/unknown", NotImplementedHandler())
}
