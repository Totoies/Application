package main

import (
	Application "Application/src"

	totoies "github.com/Totoies/Totoies"
)

func main() {

	// Configure teh Routing in src routes.go
	totoies.InitStaticFolder(Application.StaticDir)
	totoies.Buid()
}
