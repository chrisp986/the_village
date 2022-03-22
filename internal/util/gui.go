package utility

import (
	"fmt"
	"os"
)

const logo = `
			   __	__                  _  __   __
			  / /  / /    __    __   __(_)/ /  / /  ____    ____  ___
			 / __//	/___/ _ \   | | / / // /  / /  / __ \  / __ \/ _ \
			/ /_ /  _  /  __/   | |/ / // /__/ /__/ /_/ /_/ /_/ /  __/
			\__//_/	/_/\___/    |___/_//____/____/_______/___  /\___/ 
			                                              __/ /
			                                             /___/
____________________________________________________________________________________________________

`

func ShowLogo() {
	ClearScreen()
	fmt.Print(logo)
	// time.Sleep(3 * time.Second)
	// ClearScreen()
}

func MainMenu() {

	// fmt.Println("Type your name: ")

	pages := []string{"(1) Create player", "(2) Do stuff", "(99) Quit"}

	fmt.Println("____Main Menu____")
	fmt.Println()

	for _, page := range pages {
		fmt.Println(page)
	}
	for {

		choice := userInput("Pick your choice")

		switch choice {
		case "1":
			name := userInput("Name")
			password := userInput("Password")

			fmt.Printf("Name: %s Password: %s\n", name, password)
			break

		case "2":
			fmt.Println("testtest")

		case "99":
			fmt.Println("Quitting The Village!")
			os.Exit(3)
		}
	}

}

func userInput(stmt string) string {

	fmt.Print(stmt + ": ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "error!"
	}
	// fmt.Printf("%T\n", input)
	return fmt.Sprintf("%s", input)

}
