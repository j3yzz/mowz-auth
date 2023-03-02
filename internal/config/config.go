package config

import (
	"github.com/j3yzz/mowz/internal/db"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

type Config struct {
	Database db.Config `koanf:"database"`
}

func New() Config {
	var instance Config

	k := koanf.New(".")

	if err := k.Load(file.Provider("configs/config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	return instance
}
