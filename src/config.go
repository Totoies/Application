package Application

/*
Configure Different Application Components
*/
import (
	Controller "Application/src/Controllers"
	"embed"

	totoies "github.com/Totoies/Totoies"
)

/*
Config Static Directory
*/
//go:embed Static/*
var StaticDir embed.FS

var AllControllers = []interface{}{
	Controller.Home,
}

func RegisterRoutes() {
	totoies.Routes["/"] = Controller.LoadHome
}

func RegisterControllers() {
	totoies.Controllers["home"] = Controller.Home
}

func init() {
	RegisterRoutes()
	RegisterControllers()
}
