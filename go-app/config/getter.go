package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Get retrieves the config file (JSON with comments).
//
// params:
// - 0: Environment config file prefix (<env>_config.jsonc). Default: "dev".
func Get(params ...string) Config {
	config := Config{}

	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}

	fileName := fmt.Sprintf("./%s_config.yaml", env)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	return config
}
