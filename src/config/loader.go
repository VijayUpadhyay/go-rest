package config

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

type (
	Config struct {
		LogDetails
		DbConnectionDetails
		AppVersion string `mapstructure:"APP_VERSION"`
		ServerPort int    `mapstructure:"SERVER_PORT"`
		TestOsVar  string `mapstructure:"TEST_OS_VAR"`
	}

	LogDetails struct {
		LogLevel string `mapstructure:"LOG_LEVEL"`
		LogPath  string `mapstructure:"LOG_PATH"`
	}

	DbConnectionDetails struct {
		DBUser   string `mapstructure:"DB_USER"`
		DBPass   string `mapstructure:"DB_PASS"`
		DBHost   string `mapstructure:"DB_HOST"`
		DBPort   string `mapstructure:"DB_PORT"`
		DBDriver string `mapstructure:"DB_DRIVER"`
	}
)

func LoadConfig() (config Config, err error) {
	// Read file path
	viper.AddConfigPath(path.Join(os.Getenv("APP_HOME")))
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func UpdateConfig() {

}
