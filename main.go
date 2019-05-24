package main

import (
	"fmt"
)


func waitAndLogoutAndDelete(token Token) {
	fmt.Println("Press enter to logout...")
	fmt.Scanln()
	logOut(token)
	fmt.Println("You have been disconnected.")
	deleteToken()
}

func main() {
	token, err := loadToken()
	if err == nil {
		fmt.Println("Loaded token:", token)
		waitAndLogoutAndDelete(token)
		return
	}

	configuration, err := loadConfiguration()
	if err != nil || configuration.Username == "" || configuration.Password == "" {
		createEmptyConfigurationFile()
		fmt.Println("Please, fill", CONFIG_PATH, "configuration file.")
		return
	}

	token, err = logIn(configuration.Username, configuration.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Your token:", token)
	saveToken(token)

	waitAndLogoutAndDelete(token)
}
