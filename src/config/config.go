package config

import "go.uber.org/zap"

type Config struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

var Cfg Config

func LoadConfig() {
	Cfg.DbHost = LoadEnv("DB_HOST")
	Cfg.DbPort = LoadEnv("DB_PORT")
	Cfg.DbUser = LoadEnv("DB_USER")
	Cfg.DbPass = LoadEnv("DB_PASS")
	Cfg.DbName = LoadEnv("DB_NAME")
	LogInfo("config loaded", zap.Any("config", "success"))
}
