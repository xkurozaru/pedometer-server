package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

func NewDBConfig() (DBConfig, error) {
	dbConfig := DBConfig{}
	err := envconfig.Process("DB", &dbConfig)
	if err != nil {
		return DBConfig{}, err
	}

	return dbConfig, nil
}

type SupabaseConfig struct {
	APIURL    string `envconfig:"API_URL"`
	APIKey    string `envconfig:"API_KEY"`
	JWTSecret string `envconfig:"JWT_SECRET"`
}

func NewSupabaseConfig() (SupabaseConfig, error) {
	supabaseConfig := SupabaseConfig{}
	err := envconfig.Process("SUPABASE", &supabaseConfig)
	if err != nil {
		return SupabaseConfig{}, err
	}

	return supabaseConfig, nil
}
