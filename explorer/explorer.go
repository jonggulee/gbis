package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

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
	// NowTime     time.Time
	NowTime string
}

func getNowTime() string {
	now := time.Now()
	// kst, _ := time.LoadLocation("Asia/Seoul")
	// kstTime := now.In(kst)
	a := fmt.Sprintln(now.Format("1994-03-01 00:00:00 KST"))
	return a
}

func home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	kstTime := getNowTime()
	fmt.Println(kstTime)
	data := homeData{"Home", bus.GetArrivalBus(), "위례중앙중학교", kstTime}
	templates.ExecuteTemplate(rw, "home", data)
}

func health(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func Start() {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/health", health)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
