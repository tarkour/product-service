package conf

import (
	"log"

	"github.com/tarkour/product-service/pkg/config"
)

func ReadConfig() (*config.Config, error) {

	cfg, err := config.LoadConfig("./internal/config")
	if err != nil {
		log.Fatal("Config error: ", err)
	}

	return cfg, nil
}
