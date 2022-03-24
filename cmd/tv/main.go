package main

import (
	"github.com/chrisp986/the_village/internal/server"
	utility "github.com/chrisp986/the_village/internal/util"
)

func main() {

	if !utility.IsSupportedOS() {
		panic("OS not supported!")
	}

	go server.Run()

	utility.ShowLogo()

	utility.MainMenu()

}
