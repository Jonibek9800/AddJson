package json

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Client struct {
	Name string
	Suename string
	Age string
	Sex string
	ClientId string
	Date     string
}

type Settings struct {
	Clients []Client
}

const settingsFilename = "settings.json"

func JsonFunk() {
	rawDataIn, err := ioutil.ReadFile(settingsFilename)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	var settings Settings
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}

	newClient := Client{
		Name:     "Jonibek",
		Suename:  "Mahmudov",
		Age:      "22",
		Sex: "man",
		ClientId: "1",
		Date:     "07.12.2020",
	}

	settings.Clients = append(settings.Clients, newClient)

	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(settingsFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}
}
