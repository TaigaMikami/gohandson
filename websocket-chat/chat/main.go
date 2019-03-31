package main

import (
	"flag"
	"log"
	"net/http"
	"html/template"
	"os"
	"path/filepath"
)

type templateHandler struct {
		filename string
		templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.templ == nil {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	}
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	r := newRoom()
	r.tracer = tracer.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)


	// get the room going
	go r.run()

	// start the web server
	log.Println("Server is running, port:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
