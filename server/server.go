package server

import (
	"fmt"
	"net/http"
	ui "starter/web/views"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Run(PORT string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := ui.Root()
		component.Render(r.Context(), w)
	})

	fmt.Printf("Server is running on port: %s", PORT)
	err := http.ListenAndServe(PORT, r)
	return err
}
