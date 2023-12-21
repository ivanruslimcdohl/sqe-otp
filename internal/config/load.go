package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type AppCfg struct {
	Port int `yaml:"PORT"`

	OTPDigits         int           `yaml:"OTP_DIGITS"`
	OTPExpirationTime time.Duration `yaml:"OTP_EXPIRATION_TIME"`
}

type DBCfg struct {
	DBHost string `yaml:"DB_HOST"`
	DBPort int    `yaml:"DB_PORT"`
	DBName string `yaml:"DB_NAME"`
	DBUser string `yaml:"DB_USER"`
}

type configs struct {
	App AppCfg
	DB  DBCfg
}

var cfg = configs{}

func Config() configs {
	return cfg
}

func Init() {
	loadCfg("app.yaml", &cfg.App)
	loadCfg("db.yaml", &cfg.DB)
}

// filename: full filename with extension
// support yaml and json only
func loadCfg(filename string, cfgSt interface{}) {
	if reflect.TypeOf(cfgSt).Kind() != reflect.Pointer {
		log.Panicln("cfgSt must be a pointer")
	}

	filePath := fmt.Sprintf("/etc/sqe/%s.%s", GetEnv(), filename)
	if !IsProd() {
		filePath = "./configs" + filePath
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Panicf("failed to open yaml cfg file, err: %v\n", err)
	}
	defer file.Close()

	fileExt := getFileExtension(filePath)
	switch fileExt {
	case "yaml":
		d := yaml.NewDecoder(file)
		if err := d.Decode(cfgSt); err != nil {
			log.Panicf("failed to decode yaml cfg file, err: %v\n", err)
		}
	case "json":
		d := json.NewDecoder(file)
		if err := d.Decode(cfgSt); err != nil {
			log.Panicf("failed to decode json cfg file, err: %v\n", err)
		}
	default:
		log.Panicf("unsupported file extension: %s\n", fileExt)
	}

}

func getFileExtension(filePath string) string {
	return strings.ToLower(filepath.Ext(filePath)[1:])
}
