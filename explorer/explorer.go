package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/jonggulee/gbis/bus"
)

const (
	port        int    = 4000
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle   string
	Buses       []bus.Bus
	StationName string
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", bus.GetArrivalBus(), "위례중앙중학교"}
	templates.ExecuteTemplate(rw, "home", data)
}

func Start() {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
