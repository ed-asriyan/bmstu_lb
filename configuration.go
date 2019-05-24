package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const TOKEN_PATH = "/tmp/.bmstu_lb"
const CONFIG_PATH = "bmstu_lb.json"

const DEFAULT_PERMISSIONS = 0660

type Configuration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token string

func loadConfiguration() (Configuration, error) {
	configFile, err := os.Open(CONFIG_PATH)
	defer configFile.Close()
	if err != nil {
		return Configuration{}, err
	}

	jsonParser := json.NewDecoder(configFile)

	var result Configuration
	jsonParser.Decode(&result)
	return result, nil
}

func createEmptyConfigurationFile() error {
	configuration := Configuration{"", ""}

	configJson, err := json.MarshalIndent(configuration, "", "   ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(CONFIG_PATH, configJson, DEFAULT_PERMISSIONS)
}

func loadToken() (Token, error) {
	token, err := ioutil.ReadFile(TOKEN_PATH)
	if err != nil {
		return "", err
	} else {
		return Token(token), nil
	}
}

func saveToken(token Token) error {
	return ioutil.WriteFile(TOKEN_PATH, []byte(token), DEFAULT_PERMISSIONS)
}

func deleteToken() error {
	return os.Remove(TOKEN_PATH)
}
