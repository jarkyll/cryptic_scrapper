package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sync"
)

var conf *Configuration

// Singleton approach
var dial sync.Once

// Load returns Configuration struct
func Load(env string) {
	_, filePath, _, _ := runtime.Caller(0)
	configFile := filePath[:len(filePath)-9]
	bytes, err := ioutil.ReadFile(configFile +
		"env" + string(filepath.Separator) + "env." + env + ".yaml")
	if err != nil {
		_ = fmt.Errorf("error reading config file, %s", err)
	}
	dial.Do(func() {
		conf = new(Configuration)
		if err := yaml.Unmarshal(bytes, conf); err != nil {
			_ = fmt.Errorf("unable to decode into struct, %v", err)
		}
	})
}

func GetConfig() *Configuration {
	return conf
}


type Configuration struct {
	Server *Server
	DB     *Database
	Logger *Logger
}

// database configuration
type Database struct {
	Path     string
	Database string
}

// server configuration
type Server struct {
	Port int
}
// logger configuration
type Logger struct {
	Path string
}