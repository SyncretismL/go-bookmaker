package main

import (
	"bookmaker/internal/config"
	"bookmaker/internal/sport"
	"bookmaker/pkg/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	logger logger.Logger
	cfg    config.Config
	client *http.Client
	sports sport.Sports
}

// NewHandler ...
func newClient(newLogger logger.Logger, cfg config.Config, sports sport.Sports) *Client {
	return &Client{
		logger: newLogger,
		cfg:    cfg,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
		sports: sports,
	}
}
func (c *Client) getSport(adress, name string) (*sport.Lines, error) {
	line := &sport.Lines{
		Lines: map[string]string{},
	}

	resp, err := c.client.Get(adress)
	if err != nil {
		return &sport.Lines{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &sport.Lines{}, err
	}

	err = json.Unmarshal(body, &line)
	if err != nil {
		return &sport.Lines{}, err
	}

	return line, nil
}

func (c *Client) worker(n, adress, name string) {
	var s sport.Sport
	for {
		line, err := c.getSport(adress, name)
		if err != nil {
			c.logger.Debugf(err.Error())
		}

		s.Sport = name

		coef, err := strconv.ParseFloat(line.Lines[strings.ToUpper(name)], 64)
		if err != nil {
			c.logger.Debugf(err.Error())
		}
		s.Coefficient = coef

		err = c.sports.Upsert(&s)
		if err != nil {
			c.logger.Debugf(err.Error())
		}

		dur, err := time.ParseDuration(n)
		if err != nil {
			c.logger.Debugf(err.Error())
		}

		time.Sleep(dur)

	}
}
