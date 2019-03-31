package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/TaigaMikami/gohandson/websocket-chat/trace"
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
	var is_trace = flag.String("is_trace", "true", "トレース")
	flag.Parse()
	r := newRoom()

	if *is_trace == "true" {
		r.tracer = trace.New(os.Stdout)
	}
	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)


	// get the room going
	go r.run()

	// start the web server
	log.Println("Server is running, port:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
