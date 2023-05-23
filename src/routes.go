package Application

import (
	"net/http"

	controllers "Application/src/Controllers"
)

/*
Perpose - Have the details of all the roots
*/

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/": controllers.Home.Load,
}
