package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/leadery/go-shit/model"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := model.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func main() {
	fmt.Printf("Hello %s/%s\n", runtime.GOOS, runtime.GOARCH)
	log.Println("Starting...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
