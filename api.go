package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func checkNetwork() (bool, error) {
	result := true

	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	// if user is not authorized in bmstu_lb any request should ne redirected to lbpfs.bmstu.ru:8003
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		result = false
		return nil
	}
	_, err := client.Get("http://bmstu.ru")

	if err != nil {
		return false, err
	} else {
		return result, nil
	}
}

func logIn(username, password string) (Token, error) {
	response, err := http.PostForm("https://lbpfs.bmstu.ru:8003/index.php?zone=bmstu_lb", url.Values{
		"auth_user": {username},
		"auth_pass": {password},
		"redirurl":  {"/"},
		"accept":    {"Continue"},
	})
	if err != nil {
		return NullToken, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return NullToken, err
	}
	defer response.Body.Close()
	bodyStr := string(body)

	if !strings.Contains(bodyStr, "Logout") {
		return NullToken, errors.New("username\\password is invalid or another device has been already connected")
	}

	token := Token(bodyStr[962:978])
	return token, nil
}

func logOut(token Token) error {
	response, err := http.PostForm("https://lbpfs.bmstu.ru:8003/", url.Values{
		"logout_id": {string(token)},
		"zone":      {"bmstu_lb"},
		"logOut":    {"Logout"},
	})
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	bodyStr := string(body)

	if !strings.Contains(bodyStr, "You have been disconnected.") {
		return errors.New("can not disconnect using this token: " + string(token))
	}

	return nil
}
