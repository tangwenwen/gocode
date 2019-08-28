package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Client struct {
	CADDRESS string `yaml:"CADDRESS"`
}
type Server struct {
	SADDRESS string `yaml:"SADDRESS"`
}
type Config struct {
	ClientConf Client `yaml:"Client"`
	ServerConf Server `yaml:"Server"`
	Name       string `yaml:"SiteName"`
}

func GetConfig() Config {
	var setting Config
	config, err := ioutil.ReadFile("C:/Users/123/Desktop/untitled1/src/demo/tcpServerDemo/conf/conf.yaml")
	if err != nil {
		fmt.Println("error", err)
	}
	yaml.Unmarshal(config, &setting)
	return setting
}
