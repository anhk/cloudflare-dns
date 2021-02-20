package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	KeyWord string `json:"keyWord"`
	InfName string `json:"infName"`
	CfKey   string `json:"cfkey"`
	ZoneId  string `json:"zoneId"`
}

func load(path string) (*Config, error) {
	var cfg Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}
