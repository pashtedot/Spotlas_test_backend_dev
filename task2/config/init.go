package config

import (
	"github.com/spf13/viper"
)

// init initialize default config params
func init() {
	// environment - could be "local", "prod", "dev"
	viper.SetDefault("env", "dev")

	viper.SetDefault("db.address", "postgres://spots_db_user:spots_db_pass@localhost:45678/spots_db?sslmode=disable")

	// grpc
	viper.SetDefault("grpc.address", "0.0.0.0:9090")
	viper.SetDefault("grpc.timeout", "60s")

	// http
	//viper.SetDefault("http.address", "0.0.0.0:8080")
	//viper.SetDefault("http.timeout", "60s")
}
