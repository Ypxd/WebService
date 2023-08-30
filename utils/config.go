package utils

import (
	"github.com/Ypxd/WebService/internal/models"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

var (
	once    sync.Once
	config  *models.Config
	cfgPath = "configuration/config.yaml"
)

func readConfig(st interface{}, cfgPath string) {
	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, fi.Size())
	_, err = f.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, st)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func GetConfig() *models.Config {
	once.Do(func() {
		var conf models.Config
		readConfig(&conf, cfgPath)
		config = &conf
	})

	if config == nil {
		log.Fatal("nil config")
	}

	return config
}
