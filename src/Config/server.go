package config

import "embed"

/*
Perpose - Hold the settings of the server
*/

/*
This struct will hold the config of surver
*/
type serversettings struct {
	ServerIP   string
	ServerPort string
	StaticDir  *embed.FS
}

// Configureing Static Folder - If you with to change change the bellow Path
// //go:embed ..\\Application\\Static\\
// var staticDir embed.FS

var ServerSettings = serversettings{
	ServerIP:   "localhost",
	ServerPort: "8080",
	// StaticDir:  &staticDir,
}
