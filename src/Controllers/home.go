package Controller

import (
	"net/http"

	totoies "github.com/Totoies/Totoies"
)

var Home = totoies.CreateController(totoies.VViews{
	"Home": "Static/Views/Home/home.html",
})

/*
LoadHome() - Function load our controller and other settings
*/
func LoadHome(w http.ResponseWriter, r *http.Request) {

	totoies.Exec(w, &Home.Templates, "Home", totoies.VData{
		"WelcomeTxt": "Welcome All",
		"Title":      "First Application",
	})
}
