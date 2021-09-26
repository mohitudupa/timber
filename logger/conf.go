package logger

import (
	"encoding/json"
	"io/ioutil"
)

const (
	defaultData = "./timber-logs"
	defaultPort = 36036
	defaultLogs = "default.log"
)

type Config struct {
	Data string   `json:"data"`
	Port int      `json:"port"`
	Logs []string `json:"logs"`
}

func NewConfig() *Config {
	conf := &Config{
		Data: defaultData,
		Port: defaultPort,
		Logs: []string{defaultLogs},
	}
	return conf
}

func (c *Config) Load() error {
	f, err := ioutil.ReadFile("./timberconf.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, c)
	if err != nil {
		return err
	}

	return nil
}
