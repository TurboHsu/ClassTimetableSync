package config

import (
	"github.com/pelletier/go-toml/v2"
	"io"
	"log"
	"os"
)

// Declear config
var Config ConfigStruct

// ReadConfig from toml file
func ReadConfig() {
	f, err := os.OpenFile("config.toml", os.O_RDONLY, 0666)
	if err != nil {
		//create config file
		body, _ := toml.Marshal(Config)
		_ = os.WriteFile("config.toml", body, 0666)
		log.Panicln("Config file not found, create a new one.")
	}
	body, _ := io.ReadAll(f)
	_ = toml.Unmarshal(body, &Config)

}
