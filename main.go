package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type server struct {
	Name string `toml:"Name"`
	Host string `toml:"Host"`
	Port string `toml:"Port"`
}
type client struct {
	Name string `toml:"Name"`
	Host string `toml:"Host"`
	Port string `toml:"Port"`
}

type connectfig struct {
	ServerMap map[string]server `toml:"server"`
	ClientMap map[string]client `toml:"client"`
}

func main() {
	configPath := "tomltest/config.toml"

	var conf connectfig
	if _, err := toml.Decode(configPath, &conf); err != nil {
		// handle error
	}

	fmt.Printf("%s", conf)

}
