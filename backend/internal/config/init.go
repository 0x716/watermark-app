package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var GlobalConfig *AppConfig

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config error: %v", err)
	}

	var cfg AppConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Load config error: %v", err)
	}

	GlobalConfig = &cfg

	os.Mkdir(cfg.Image.ImageDir, os.ModePerm)
	os.Mkdir(cfg.Image.OutputDir, os.ModePerm)
	os.Mkdir(cfg.Watermark.WatermarkDir, os.ModePerm)

	log.Println(GlobalConfig)

	return nil
}
