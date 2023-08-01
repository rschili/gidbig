package cfg

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Config struct with all parameters
type Config struct {
	Token       string `yaml:"token"`
	Shard       string `yaml:"shard"`
	ShardCount  string `yaml:"shardcount"`
	Owner       string `yaml:"owner"`
	Port        int    `yaml:"port"`
	RedirectURL string `yaml:"redirecturl"`
	Ci          int    `yaml:"ci"`
	Cs          string `yaml:"cs"`
}

const s = "config.yaml"

// LoadConfigFile config.yaml and creates a Config struct
func LoadConfigFile() *Config {
	return loadFile(s)
}

func loadFile(cf string) *Config {
	config := &Config{}
	configFile, err := os.Open(cf)
	if err != nil {
		log.Warningln("Could not load config file.", err)
	}
	defer configFile.Close()

	d := yaml.NewDecoder(configFile)

	if err := d.Decode(&config); err != nil {
		log.Error("could not decode config: ", err)
	}

	return config
}
