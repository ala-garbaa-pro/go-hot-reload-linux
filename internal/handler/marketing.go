package handler

import (
	"net/http"

	"github.com/mavolin/go-htmx"

	"go-hot-reload-linux/internal/model"
	"go-hot-reload-linux/web/view/layout"
	"go-hot-reload-linux/web/view/marketing"
)

type Marketing struct {
}

func (h *Marketing) GetLandingPage(w http.ResponseWriter, r *http.Request) {
	user := model.User{ID: 1, Name: "Bob Loblaw"}
	p := marketing.LandingPage(user)

	if htmx.Request(r) != nil {
		p.Render(r.Context(), w)
	} else {
		b := layout.Base(p)
		b.Render(r.Context(), w)
	}
}

func (h *Marketing) GetAboutPage(w http.ResponseWriter, r *http.Request) {
	p := marketing.AboutPage()

	if htmx.Request(r) != nil {
		p.Render(r.Context(), w)
	} else {
		b := layout.Base(p)
		b.Render(r.Context(), w)
	}
}
