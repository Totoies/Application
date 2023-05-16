package util

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/Totoies/Totoies/src/Config"
)

/*
	Out Main Application Object

// If a user want to pass the application settings directly they can provide it in
// configData: {"AppName" : "name"}
// AppConfig.AppName = "App Name"
*/
type app struct{}

// New utillication Created
var Application = app{}

/*
This function Will Build our server configurations
Arguments --------------
ServerConfig struct type Config holding configuration values of the server
*/
func (a *app) ConfigureServer() {
	// Define the server routes
	for route, function := range config.Routes {
		http.HandleFunc(route, function)
	}

	// // Create a new file system from the embedded static directory
	// staticFS, err := fs.Sub(settings.ServerSettings.StaticDir, "static")
	// if err != nil {
	// 	fmt.Println("Failed to get static directory:", err.Error())
	// 	return
	// }

	// // Create a file server for the static directory
	// fileServer := http.FileServer(http.FS(staticFS))

	// // Serve the static files from the /static path prefix
	// http.Handle("/static/", http.StripPrefix("/static/", fileServer))
}

/*
This function will Start the server
Arguments --------------
ServerConfig struct type Config holding configuration values of the server
*/
func (a *app) Run() {
	// Start the server
	fmt.Printf("Server starting on http://%s:%s", config.ServerSettings.ServerIP, config.ServerSettings.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerSettings.ServerIP+":"+config.ServerSettings.ServerPort, nil))
}
