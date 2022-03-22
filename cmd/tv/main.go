package main

import (
	utility "github.com/chrisp986/the_village/internal/util"
)

func main() {

	if !utility.IsSupportedOS() {
		panic("OS not supported!")
	}

	utility.ShowLogo()

	utility.MainMenu()

}
