package handlers

import (
	"net/http"

	"github.com/tte05/udemy_golang/pkg/config"
	"github.com/tte05/udemy_golang/pkg/models"
	"github.com/tte05/udemy_golang/pkg/render"
)


type Repository struct{
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers (r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home-page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	sidekickMap := make(map[string]string)
	sidekickMap["morty"] = "Ooh, wee!"
	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	sidekickMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about-page.html", &models.TemplateData{
		StringMap: sidekickMap,
	})
}