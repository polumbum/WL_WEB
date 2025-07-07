package service

import (
	"encoding/json"
	"os"
)

type LimitsConfig struct {
	MinAge int `json:"minAge"`
}

type Limitations struct {
	Limitations LimitsConfig `json:"limitations"`
}

func LoadLimits(filename string) (Limitations, error) {
	var config Limitations
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
