package config

/*
Perpose - Hold the settings of the application
*/
// import _ "embed"

// Enviourment of the application which can be either Dev/Prod
const (
	prod = false
	dev  = true
)

var (
	Enviourment = dev
)

// // Our Server Config Json
// //go:embed ..\\..\\app.settings.json
// var ApplicationSettings string
