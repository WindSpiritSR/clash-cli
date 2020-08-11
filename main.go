package main

import (
	"clash-cli/api"
	"clash-cli/model"
	"clash-cli/step"
	"clash-cli/storage"
	"log"
)

func main() {
	checkConf()

	root := step.Root{
		Client: getClientData(),
	}
	if err := root.Run(); err != nil {
		log.Fatalln(err)
	}
}

func checkConf() {
	db, err := storage.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	if err = db.CheckKey(model.DB_KEY_URL, model.DEFAULT_API_URL); err != nil {
		log.Fatalln(err)
	}
	if err = db.CheckKey(model.DB_KEY_SECRET, model.DEFAULT_API_SECRET); err != nil {
		log.Fatalln(err)
	}
}

func getClientData() *api.Client {
	db, err := storage.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	baseUrl, err := db.GetUrl()
	if err != nil {
		log.Fatalln(err)
	}

	secret, err := db.GetSecret()
	if err != nil {
		log.Fatalln(err)
	}

	client := &api.Client{
		BaseURL: baseUrl,
		Secret:  secret,
	}
	return client
}
