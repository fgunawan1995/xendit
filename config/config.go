package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fgunawan1995/xendit/util"
)

type Config struct {
	Marvel MarvelAPIConfig `json:"marvel"`
	Server ServerConfig    `json:"server"`
	Cache  CacheConfig     `json:"cache"`
}

type CacheConfig struct {
	DefaultExpirationInMinutes int64 `json:"default_expiration_in_minutes"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type MarvelAPIConfig struct {
	Host                  string `json:"host"`
	PublicKey             string `json:"public_key"`
	PrivateKey            string `json:"private_key"`
	EndpointGetCharacters string `json:"get_character"`
	TimeoutInMilliSecond  int64  `json:"timeout_in_millisecond"`
}

func GetConfig(path string) *Config {
	var byteValue []byte
	var conf *Config
	jsonFile, err := os.Open(fmt.Sprintf("%s%s.json", path, util.GetEnv()))
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(byteValue), &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
