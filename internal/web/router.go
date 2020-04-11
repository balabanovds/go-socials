package web

import (
	"balabanovds/go-social/cfg"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Start() {
	fmt.Printf("Serving at http://%s:%s\n", cfg.App.Host, cfg.App.Port)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", index)
	r.Get("/err", err)
	r.Get("/login", login)
	r.Get("/logout", logout)
	r.Get("/signup", signup)
	r.Post("/authenticate", authenticate)

	r.Route("/users", func(r chi.Router) {
		r.Post("/new", createUser)
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, cfg.App.Static))
	FileServer(r, "/static", filesDir)

	addr := fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
