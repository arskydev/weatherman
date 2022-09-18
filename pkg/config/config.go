package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	ConfPath    string `yaml:"CONF_PATH"` // I know it's not an intuitive way to define consts names this way, but it's the Go way.
	PassEnvName string `yaml:"PASS_ENV_NAME"`
	AppPort     string `yaml:"APP_PORT"`
}

func NewAppConfig(confPath string) (*AppConfig, error) {
	yamlFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return nil, err
	}

	conf := &AppConfig{}
	err = yaml.Unmarshal(yamlFile, conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}
