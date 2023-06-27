package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type routes struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}
type config struct {
	AppName      string `yaml:"appname"`
	PreFork      bool   `yaml:"prefork"`
	ServerHeader string `yaml:"serverheader"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
}
type database struct {
	Driver string `yaml:"driver"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	DB     string `yaml:"db"`
}
type serverConfig struct {
	Server struct {
		Config config   `yaml:"config"`
		Routes []routes `yaml:"routes"`
	} `yaml:"server"`
	Database database `yaml:"database"`
}

var (
	_cfg serverConfig
)

// Read Yaml Config File and unmarshal into the _cfg variable
func readYaml() (serverConfig, error) {
	_cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Server->readYaml I : %+v\n", err)
		return serverConfig{}, err
	}
	_cfgFile := path.Join(_cwd, "config", "config.yaml")
	_, err = os.Stat(_cfgFile)
	if err != nil {
		log.Printf("Server->readYaml II : %+v\n", err)
		return serverConfig{}, err
	}

	_cfgByteContent, err := os.ReadFile(_cfgFile)
	if err != nil {
		log.Printf("Server->readYaml III : %+v\n", err)
		return serverConfig{}, err
	}
	err = yaml.Unmarshal(_cfgByteContent, &_cfg)
	if err != nil {
		log.Printf("Server->readYaml IV : %+v\n", err)
	}
	log.Printf("Server->readYaml V : % +v\n", _cfg)
	return _cfg, nil
}
func readConfig() fiber.Config {
	_, err := readYaml()
	if err != nil {
		log.Printf("Server->readConfig I : %+v\n", err)
	}
	fmt.Printf("Server->readConfig II : %+v\n", _cfg)
	cfg := new(fiber.Config)
	cfg.AppName = _cfg.Server.Config.AppName
	cfg.Prefork = _cfg.Server.Config.PreFork
	cfg.ServerHeader = _cfg.Server.Config.ServerHeader
	cfg.EnablePrintRoutes = true
	return *cfg
}
func readDBConfig() map[string]interface{} {
	var db map[string]interface{}
	in := database{
		Driver: _cfg.Database.Driver,
		Host:   _cfg.Database.Host,
		Port:   _cfg.Database.Port,
		DB:     _cfg.Database.DB,
	}
	inInterface, _ := json.Marshal(in)
	_ = json.Unmarshal(inInterface, &db)
	return db
}

// func (app *fiber.App) registerRoutes() {
// 	app.Get('/', func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{
// 			"index": true,
// 		})
// 	})
// }
