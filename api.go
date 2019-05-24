package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func logIn(username, password string) (Token, error) {
	response, err := http.PostForm("https://lbpfs.bmstu.ru:8003/index.php?zone=bmstu_lb", url.Values{
		"auth_user": {string(username)},
		"auth_pass": {string(password)},
		"redirurl":  {"/"},
		"accept":    {"Continue"},
	})
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	token := Token(string(body)[962:978])
	return token, nil
}

func logOut(token Token) error {
	_, err := http.PostForm("https://lbpfs.bmstu.ru:8003/", url.Values{
		"logout_id": {string(token)},
		"zone":      {"bmstu_lb"},
		"logOut":    {"Logout"},
	})
	if err != nil {
		return err
	}
	return nil
}