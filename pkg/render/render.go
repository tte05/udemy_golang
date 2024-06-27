package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tte05/udemy_golang/pkg/config"
	"github.com/tte05/udemy_golang/pkg/models"
)

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
    return td
}

var app *config.AppConfig
func NewTemplates(a *config.AppConfig){
    app = a
}

// RenderTemplate serves as a wrapper and renders
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
    var tc map[string]*template.Template
    if app.UseCache{
        tc = app.TemplateCache
    }else{
        tc, _ = CreateTemplateCache()
    }

	// Get the right template from the cache.
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("Template %s not in cache", ok)
	}

	// Store result in a buffer and double-check if it is a valid value.
	buf := new(bytes.Buffer)

    td = AddDefaultData(td)
    
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// Render that template.
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache creates a cache of templates.
func CreateTemplateCache() (map[string]*template.Template, error) {
	theCache := map[string]*template.Template{}

	// Get all available files *-page.html from the ./templates folder.
	pages, err := filepath.Glob("./templates/*-page.html")
	if err != nil {
		return theCache, err
	}

	// Range through the slice of *-page.html files.
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return theCache, err
		}

		matches, err := filepath.Glob("./templates/*-layout.html")
		if err != nil {
			return theCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*-layout.html")
			if err != nil {
				return theCache, err
			}
		}

		theCache[name] = ts
	}

	return theCache, nil
}
