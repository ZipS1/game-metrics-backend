package config

import (
	"crypto/ed25519"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type JwtTokenConfig struct {
	PublicKeyPemFilepath       string        `mapstructure:"public_key_filepath" env:"JWT_PUBLIC_KEY_PEM_FILEPATH"`
	PrivateKeyPemFilepath      string        `mapstructure:"private_key_filepath" env:"JWT_PRIVATE_KEY_PEM_FILEPATH"`
	JwtExpirationTime          time.Duration `mapstructure:"jwt_expiration_time"`
	RefreshTokenExpirationTime time.Duration `mapstructure:"refresh_token_expiration_time"`

	Ed25519PrivateKey ed25519.PrivateKey `mapstructure:"-"`
	Ed25519PublicKey  ed25519.PublicKey  `mapstructure:"-"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	SslMode  string `mapstructure:"sslmode"`
	Timezone string `mapstructure:"timezone"`
}

func (c *DatabaseConfig) GetConnectionString() string {
	templateString := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s"
	return fmt.Sprintf(templateString, c.Host, c.User, c.Password, c.DbName, c.Port, c.SslMode, c.Timezone)
}

type AMQPConfig struct {
	Host     string        `mapstructure:"host"`
	Port     int           `mapstructure:"port"`
	User     string        `mapstructure:"user"`
	Password string        `mapstructure:"password"`
	Timeout  time.Duration `mapstructure:"timeout"`
}

func (c *AMQPConfig) GetConnectionString() string {
	templateString := "amqp://%s:%s@%s:%d/"
	return fmt.Sprintf(templateString, c.User, c.Password, c.Host, c.Port)
}

type Config struct {
	DomainName        string         `mapstructure:"domain_name"`
	Port              int            `mapstructure:"port"`
	PublicUriPrefix   string         `mapstructure:"public_uri_prefix"`
	InternalUriPrefix string         `mapstructure:"internal_uri_prefix"`
	JwtToken          JwtTokenConfig `mapstructure:"jwt_token"`
	Database          DatabaseConfig `mapstructure:"database"`
	AMQP              AMQPConfig     `mapstructure:"amqp"`
}

func loadConfig(configPath string) (*Config, error) {
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(configPath), filepath.Ext(configPath)))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(configPath))
	configureEnvVarsAndDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := validateConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := loadJwtKeys(&cfg.JwtToken); err != nil {
		return nil, err
	}

	return &cfg, nil
}
