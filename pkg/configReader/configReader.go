// Read configuration file from configs/cfg/config.yml
package configReader

import (
	"io/ioutil"
	"os"

	"github.com/kooberetis/keeper/pkg/utils"
	"gopkg.in/yaml.v2"
)

// Configuration file path
var configPath string = os.Getenv("PWD") + "/configs/cfg/config.yaml"

// Configuration structure from yaml file
type config struct {
	Global struct {
		LogLevel string `yaml:"logLevel"`
		Timeout  string `yaml:"timeout"`
	} `yaml:"global"`
	Note []struct {
		Alias       string `yaml:"alias"`
		ExpiredDate string `yaml:"expiredDate"`
	} `yaml:"note"`
	Url []struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Description string `yaml:"description"`
		CertInfo    []struct {
			CN      string
			Expired string
		}
	} `yaml:"url"`
}

// Configuration file reader from path: /configs/cfg/config.yml
func ReadConfig() config {
	cfg := config{}
	yamlFile, errReadFile := ioutil.ReadFile(configPath)
	if errReadFile != nil {
		utils.WarningLog.Printf("Yaml content incorrect. Verify %v file exist", configPath)
	}
	errUnmarshal := yaml.Unmarshal(yamlFile, &cfg)
	if errUnmarshal != nil {
		utils.WarningLog.Println("Can't load yaml content")
	} else {
		if cfg.Global.LogLevel == "debug" {
			utils.InfoLog.Println("Yaml content successfuly load")
			utils.InfoLog.Printf("Configuration file: %+v\n", cfg)
		}
	}
	return cfg
}
