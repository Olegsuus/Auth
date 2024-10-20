package config

import (
	"github.com/spf13/viper"
	"os"
)

type MongoSetting struct {
	URI        string `mapstructure:"uri" yaml:"uri"`
	DBNAME     string `mapstructure:"database" yaml:"database"`
	Timeout    int    `mapstructure:"timeout" yaml:"timeout"`
	Collection struct {
		Name string `mapstructure:"collection_name" yaml:"collection_name"`
	} `mapstructure:"collection" yaml:"collection"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret" yaml:"secret"`
	Expiry string `mapstructure:"expiry" yaml:"expiry"`
}

type ServerConfig struct {
	Port string `mapstructure:"port" yaml:"port"`
}

type Config struct {
	Mongo  MongoSetting `mapstructure:"mongo" yaml:"mongo"`
	JWT    JWTConfig    `mapstructure:"jwt" yaml:"jwt"`
	Server ServerConfig `mapstructure:"server" yaml:"server"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Путь к конфигурационному файлу

	// Установка значений по умолчанию или чтение из переменных окружения
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Альтернативно можно переопределить значения из переменных окружения
	config.Mongo.URI = os.Getenv("MONGO_URI")
	config.JWT.Secret = os.Getenv("JWT_SECRET")
	config.JWT.Expiry = os.Getenv("JWT_EXPIRY")
	config.Server.Port = os.Getenv("SERVER_PORT")

	return &config, nil
}

//func LoadConfig(configPath string) (*Config, error) {
//	viper.SetConfigFile(configPath) // Устанавливаем полный путь к файлу конфигурации
//	viper.SetConfigType("yaml")
//	viper.AutomaticEnv()
//
//	viper.BindEnv("mongo.uri", "AUTH_MONGO_URI")
//	viper.BindEnv("mongo.database", "AUTH_MONGO_DATABASE")
//	viper.BindEnv("mongo.timeout", "AUTH_MONGO_TIMEOUT")
//	viper.BindEnv("mongo.collection.collection_name", "AUTH_MONGO_COLLECTION_NAME")
//	viper.BindEnv("jwt.secret", "AUTH_JWT_SECRET")
//	viper.BindEnv("jwt.expiry", "AUTH_JWT_EXPIRY")
//	viper.BindEnv("server.port", "AUTH_SERVER_PORT")
//
//	if err := viper.ReadInConfig(); err != nil {
//		return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
//	}
//
//	var cfg Config
//
//	if err := viper.Unmarshal(&cfg); err != nil {
//		return nil, fmt.Errorf("ошибка распаковки конфигурации: %w", err)
//	}
//
//	fmt.Printf("Loaded Config: %+v\n", cfg)
//
//	if err := applyConfigOverrides(&cfg); err != nil {
//		return nil, err
//	}
//
//	return &cfg, nil
//}
//
//func applyConfigOverrides(cfg *Config) error {
//	_, err := time.ParseDuration(cfg.JWT.Expiry)
//	if err != nil {
//		return fmt.Errorf("не удалось распарсить срок действия токена: %w", err)
//	}
//
//	return nil
//}
