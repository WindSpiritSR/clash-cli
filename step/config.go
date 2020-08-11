package step

import (
	"clash-cli/api"
	"clash-cli/model"
	"clash-cli/storage"
	"errors"
	"github.com/manifoldco/promptui"
	"log"
	"net/url"
	"regexp"
)

type Config struct {
	*api.Client
	LastStep Step
}

func (c Config) Run() error {
	prompt := promptui.Select{
		Label: model.PROMPT_CONFIG_LABEL,
		Items: []string{model.PROMPT_CONFIG_ITEM_URL, model.PROMPT_CONFIG_ITEM_SECRET},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	if err = c.setConfig(result); err != nil {
		log.Fatalln(err)
	}
	return c.LastStep.Run()
}

func (c Config) setConfig(confKey string) error {
	db, err := storage.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	switch confKey {
	case model.PROMPT_CONFIG_ITEM_URL:
		curUrl, err := db.GetUrl()
		if err != nil {
			return err
		}
		prompt := promptui.Prompt{
			Label: confKey,
			Validate: func(s string) error {
				u, err := url.Parse(s)
				if err != nil {
					return err
				}
				r, _ := regexp.Compile("^(https://|http://)\\S+\\w$")
				if !r.MatchString(u.String()) {
					return errors.New(model.WARNING_UNKNOWN_URL_TYPE)
				}
				return nil
			},
			Default:   curUrl,
			AllowEdit: true,
		}
		apiUrl, err := prompt.Run()
		if err != nil {
			return err
		}
		c.Client.BaseURL = apiUrl

		err = db.SetUrl(apiUrl)
		if err != nil {
			return err
		}
	case model.PROMPT_CONFIG_ITEM_SECRET:
		curSecret, err := db.GetSecret()
		if err != nil {
			return err
		}
		prompt := promptui.Prompt{
			Label:     confKey,
			Default:   curSecret,
			AllowEdit: true,
		}
		apiSecret, err := prompt.Run()
		if err != nil {
			return err
		}
		c.Client.Secret = apiSecret

		if err = db.SetSecret(apiSecret); err != nil {
			return err
		}
	}

	return nil
}
