package ngh

import (
	"encoding/json"
	"os"
)

type Config struct {
	ProjectName string       `json:"name"`
	OutDir      string       `json:"out"`
	Deployments []Deployment `json:"deployments"`
}

func ParseConfig(path string) (cfg *Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	cfg = new(Config)
	json.NewDecoder(file).Decode(cfg)
	return
}
