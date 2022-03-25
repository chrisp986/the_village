package utility

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

	"golang.org/x/term"
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
const serverUrl string = "http://localhost:8001"

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
			name, password, email, err := credentials()
			if err != nil {
				fmt.Println(err)
			}

			// name := userInput("Choose your name: ")
			createPlayer(name, password, email)
			// password := userInput("Password")

			// break

		case "2":
			fmt.Println("testtest")

		case "99":
			fmt.Println("Quitting The Village!")
			time.Sleep(2 * time.Second)
			os.Exit(0)
		}
	}

}

func userInput(stmt string) string {

	fmt.Print(stmt)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "error!"
	}
	return fmt.Sprintf("%s", input)

}

type Player struct {
	PlayerID    int32  `json:"player_id"`
	PlayerName  string `json:"player_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PlayerScore int64  `json:"player_score"`
	Active      bool   `json:"active"`
	Connected   bool   `json:"connected"`
	Created     string `json:"created"`
}

func createPlayer(name string, password string, email string) {

	values := Player{
		PlayerID:    5,
		PlayerName:  name,
		Password:    password,
		Email:       email,
		PlayerScore: 0,
		Active:      true,
		Connected:   false,
		Created:     time.Now().Format("2006-01-02 15:04:05"),
	}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(serverUrl+"/v1/players", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func credentials() (username string, password string, email string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err = reader.ReadString('\n')
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Enter Email: ")
	email, err = reader.ReadString('\n')
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", "", err
	}
	return username, string(bytePassword[:]), email, nil
}
