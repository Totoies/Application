package controlers

import (
	_ "embed"
	"net/http"
	"text/template"
)

// Private Variables
type home struct {
	WelcomeTxt string
}

// Publicaly assesible variables
var Home = home{
	WelcomeTxt: "Wlecome All",
}

// Load our View During Building the application
// load home.html
//
//go:embed ..\\Views\\home.html
var html string

var homeTemplate, err = template.New("template").Parse(html)

/*
Load() - Function load our controller and other settings
*/
func (HomeController *home) Load(w http.ResponseWriter, r *http.Request) {

	if err != nil {
		// fmt.Fprintf(w, html)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	} else {
		homeTemplate.Execute(w, Home)
	}
}
