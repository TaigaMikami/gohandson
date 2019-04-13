package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/TaigaMikami/gohandson/websocket-chat/trace"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
		filename string
		templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if t.templ == nil {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	}
	data := map[string]interface{} {
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "auth",
		Value: "",
		Path: "/",
		MaxAge: -1,
	})
	w.Header()["Location"] = []string{"/chat"}
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	var is_trace = flag.String("is_trace", "true", "トレース")
	flag.Parse()
	err := godotenv.Load("envfiles/development.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gomniauth.SetSecurityKey("98dfbg7iu2nb4uywevihjw4tuiyub34noilk")
	gomniauth.WithProviders(
		github.New(os.Getenv("GITHUB_ID"), os.Getenv("GITHUB_SECRET"), "http://localhost:8080/auth/callback/github"),
		google.New(os.Getenv("GOOGLE_ID"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/callback/google"),
	)
	r := newRoom(UseGravatar)

	if *is_trace == "true" {
		r.tracer = trace.New(os.Stdout)
	}
	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", logout)
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Server is running, port:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
