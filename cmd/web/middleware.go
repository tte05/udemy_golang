package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func hitLogger(next http.Handler) http.Handler{
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	log.Println("HIT .... road, Jack ")
	next.ServeHTTP(w,r)
	})
}

func NoSurf(next http.Handler) http.Handler{
	csrfHandler :=nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler

}

func SessionLoad(next http.Handler)http.Handler{
	return session.LoadAndSave(next)
}