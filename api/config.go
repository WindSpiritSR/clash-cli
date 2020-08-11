package api

import (
	"fmt"

	"clash-cli/model"
	T "github.com/Dreamacro/clash/tunnel"
	"github.com/levigross/grequests"
)

type Client struct {
	BaseURL string
	Secret  string
}

func (c *Client) GetConfigs() (*model.Config, error) {
	resp, err := grequests.Get(c.BaseURL+model.API_PATH_CONFIGS, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + c.Secret},
	})
	if err != nil {
		return nil, err
	}
	rst := &model.Config{}
	if err := resp.JSON(rst); err != nil {
		return nil, err
	}
	return rst, nil
}

func (c *Client) UpdateMode(mode T.TunnelMode) error {
	resp, err := grequests.Patch(c.BaseURL+model.API_PATH_CONFIGS, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + c.Secret},
		JSON: &model.Config{
			Mode: &mode,
		},
	})
	if err != nil {
		return err
	}
	if !resp.Ok {
		return fmt.Errorf("%d: %s", resp.StatusCode, resp.String())
	}
	return nil
}
