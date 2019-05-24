package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	configuration, err := loadConfiguration()
	if err != nil || configuration.Username == "" || configuration.Password == "" {
		createEmptyConfigurationFile()
		fmt.Println("Please, fill", CONFIG_PATH, "configuration file.")
		return
	}

	token, err := loadToken()
	if err == nil {
		fmt.Println("Loaded token:", token)
	}

	isRunning := true
	go func() {
		isNetworkReachable := false
		for isRunning {
			newNetworkStatus := checkNetwork()
			if newNetworkStatus {
				if !isNetworkReachable {
					fmt.Println("Network is reachable.")
					fmt.Println("You will be automatically reconnected after bmstu_lb timeout.")
					fmt.Println()
					fmt.Print("Press enter to ")
					if token == NullToken {
						fmt.Print("exit the program")
					} else {
						fmt.Print("logout")
					}
					fmt.Println("...")

					fmt.Println()
				}
				isNetworkReachable = newNetworkStatus

				time.Sleep(15 * time.Second)
				continue
			}

			fmt.Println("Network is unreachable.")
			token = NullToken

			fmt.Println("Authorizing in bmstu_lb...")
			token, err = logIn(configuration.Username, configuration.Password)
			if err != nil {
				fmt.Println("Can not authorize:")
				fmt.Println(err)
				fmt.Println()
				continue
			}

			fmt.Println("Done. Your token:", token)
			saveToken(token)
			fmt.Println()
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Scanln()
	isRunning = false
	if token != NullToken {
		logOut(token)
		fmt.Println("You have been disconnected.")
	}
	deleteToken()
	os.Exit(0)
}
