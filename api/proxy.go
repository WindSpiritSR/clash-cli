package api

import (
	"clash-cli/model"
	"context"
	"fmt"
	"github.com/levigross/grequests"
)

func (c *Client) GetProxies() (*model.Proxies, error) {
	resp, err := grequests.Get(c.BaseURL+model.API_PATH_PROXIES, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + c.Secret},
	})
	if err != nil {
		return nil, err
	}
	rst := &model.Proxies{}
	if err := resp.JSON(rst); err != nil {
		return nil, err
	}
	return rst, nil
}

func (c *Client) UpdateProxy(group, proxy string) error {
	resp, err := grequests.Put(c.BaseURL+model.API_PATH_PROXIES+"/"+group, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + c.Secret},
		JSON: struct {
			Name string `json:"name"`
		}{Name: proxy},
	})
	if err != nil {
		return err
	}
	if !resp.Ok {
		return fmt.Errorf("%d: %s", resp.StatusCode, resp.String())
	}
	return nil
}

func (c *Client) GetLatency(ctx context.Context, proxy string) (*model.Latency, error) {
	resp, err := grequests.Get(c.BaseURL+model.API_PATH_PROXIES+"/"+proxy+model.API_PATH_PROXIES_LATENCY, &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + c.Secret},
		Context: ctx,
		Params: map[string]string{
			"url":     model.API_LATENCY_TEST_URL,
			"timeout": model.API_LATENCY_TEST_TIMEOUT,
		},
	})

	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, resp.String())
	}
	latency := &model.Latency{}
	if err := resp.JSON(latency); err != nil {
		return nil, err
	}

	return latency, nil
}
