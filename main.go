package main

import (
	"fmt"
	"strconv"
	"time"
)

func log(s ...any) {
	fmt.Print(time.Now().Format("01-02-2006 15:04:05 | "))
	fmt.Println(s...)
}

func main() {
	log("Starting...")
	wasNetworkReachable := false
	var lastReconnect int64 = 0
	for {
		redirectUrl := checkNetwork()
		networkStatus := redirectUrl == ""

		if !networkStatus {
			if wasNetworkReachable {
				log("Network has become unreachable after " + strconv.FormatInt((time.Now().UnixMilli()-lastReconnect)/1000, 10) + " seconds.")
			}

			log("Authorizing in Nandos...")
			err := logIn(redirectUrl)
			if err != nil {
				log("Can not authorize: ", err)
				log()
				time.Sleep(10 * time.Second)
			}

			log()
			time.Sleep(5 * time.Second)
		} else {
			// network state changed
			if !wasNetworkReachable {
				log("No Nandos restrictions were detected.")
				log("You will be automatically reconnected after Nandos timeout.")
				log()
				log("Press Ctrl+C to exit the program...")
				log()

				lastReconnect = time.Now().UnixMilli()
			}

			time.Sleep(15 * time.Second)
		}
		wasNetworkReachable = networkStatus
	}
}
