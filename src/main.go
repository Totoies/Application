package main

import (
	"os"

	util "github.com/Totoies/Totoies/src/Util"
)

var CurrentPath, err = os.Getwd()

func main() {

	util.Application.ConfigureServer()
	util.Application.Run()

}
