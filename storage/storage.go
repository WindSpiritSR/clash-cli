package storage

import (
	"clash-cli/model"
	"log"
	"os"
	"path"

	"github.com/asdine/storm"
)

type Client struct {
	db *storm.DB
}

func Open() (*Client, error) {
	var db *storm.DB
	home, err := os.UserHomeDir()
	if err != nil {
		db, err = storm.Open(model.DB_FILE_NAME)
	} else {
		cfgPath := path.Join(home, ".config", "clash")
		if err := os.MkdirAll(cfgPath, os.ModePerm); err != nil {
			return nil, err
		}
		db, err = storm.Open(path.Join(cfgPath, model.DB_FILE_NAME))
	}
	if err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}

func (c *Client) SetUrl(clashApiUrl string) error {
	return c.db.Set(model.DB_BUCKET_NAME, model.DB_KEY_URL, clashApiUrl)
}

func (c *Client) GetUrl() (string, error) {
	var clashApiUrl string
	if err := c.db.Get(model.DB_BUCKET_NAME, model.DB_KEY_URL, &clashApiUrl); err != nil {
		return "", err
	}
	return clashApiUrl, nil
}

func (c *Client) SetSecret(clashApiSecret string) error {
	return c.db.Set(model.DB_BUCKET_NAME, model.DB_KEY_SECRET, clashApiSecret)
}

func (c *Client) GetSecret() (string, error) {
	var clashApiSecret string
	if err := c.db.Get(model.DB_BUCKET_NAME, model.DB_KEY_SECRET, &clashApiSecret); err != nil {
		return "", err
	}
	return clashApiSecret, nil
}

func (c *Client) CheckKey(key string, defaultValue string) error {
	var tempValue string
	if err := c.db.Get(model.DB_BUCKET_NAME, key, &tempValue); err != nil {
		if err.Error() == model.DB_ERROR_KEYNOTFOUND {
			if err = c.db.Set(model.DB_BUCKET_NAME, key, defaultValue); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (c *Client) Close() {
	if err := c.db.Close(); err != nil {
		log.Println(err)
	}
}
