package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tte05/udemy_golang/pkg/config"
	"github.com/tte05/udemy_golang/pkg/handlers"
	"github.com/tte05/udemy_golang/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {
	//dont forget to change to true in Production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteDefaultMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("error creating template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting server at port %s\n", portNumber)
	
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
