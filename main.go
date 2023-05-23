package main

import (
	"fmt"
	"net/http"

	totoies "github.com/Totoies/Totoies"
)

func main() {

	totoies.App.AddRoutes(totoies.Routes{
		"/": func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, World!")
		},
	})

	totoies.App.Buid()

}
