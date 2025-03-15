package config

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

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

type JwtTokenConfig struct {
	PublicKeyPemFilepath       string        `mapstructure:"public_key_filepath" env:"JWT_PUBLIC_KEY_PEM_FILEPATH"`
	PrivateKeyPemFilepath      string        `mapstructure:"private_key_filepath" env:"JWT_PRIVATE_KEY_PEM_FILEPATH"`
	JwtExpirationTime          time.Duration `mapstructure:"jwt_expiration_time"`
	RefreshTokenExpirationTime time.Duration `mapstructure:"refresh_token_expiration_time"`

	Ed25519PrivateKey ed25519.PrivateKey `mapstructure:"-"`
	Ed25519PublicKey  ed25519.PublicKey  `mapstructure:"-"`
}

type Config struct {
	DomainName    string         `mapstructure:"domain_name"`
	Port          int            `mapstructure:"port"`
	BaseUriPrefix string         `mapstructure:"base_uri_prefix"`
	JwtToken      JwtTokenConfig `mapstructure:"jwt_token"`
	Database      DatabaseConfig `mapstructure:"database"`
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

func validateConfig() error {
	errorMessageTemplate := "no '%s' keyword found in config"

	if err := validateNestedYaml(errorMessageTemplate, "database", []string{"host", "user", "password", "dbname", "sslmode"}); err != nil {
		return err
	}

	if err := validateNestedYaml(errorMessageTemplate, "jwt_token", []string{"public_key_filepath", "private_key_filepath"}); err != nil {
		return err
	}

	if err := validateSingleKeyword(errorMessageTemplate, "domain_name"); err != nil {
		return err
	}

	return nil
}

func validateSingleKeyword(errorMessageTemplate string, keyword string) error {
	if !viper.IsSet(keyword) {
		return fmt.Errorf(errorMessageTemplate, keyword)
	}
	return nil
}

func validateNestedYaml(errorMessageTemplate string, topLevelKeyword string, requiredChildren []string) error {
	if !viper.IsSet(topLevelKeyword) {
		return fmt.Errorf(errorMessageTemplate, topLevelKeyword)
	}
	for _, child := range requiredChildren {
		keywordFullName := fmt.Sprintf("%s.%s", topLevelKeyword, child)
		if !viper.IsSet(keywordFullName) {
			return fmt.Errorf(errorMessageTemplate, keywordFullName)
		}
	}
	return nil
}

func loadJwtKeys(jwtConfig *JwtTokenConfig) error {
	privData, err := os.ReadFile(jwtConfig.PrivateKeyPemFilepath)
	if err != nil {
		return fmt.Errorf("error reading private key file: %w", err)
	}
	privBlock, _ := pem.Decode(privData)
	if privBlock == nil {
		return errors.New("failed to decode PEM block from private key file")
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse ED25519 private key: %w", err)
	}
	ed25519Priv, ok := parsedKey.(ed25519.PrivateKey)
	if !ok {
		return errors.New("provided private key is not an ED25519 key")
	}
	jwtConfig.Ed25519PrivateKey = ed25519Priv

	pubData, err := os.ReadFile(jwtConfig.PublicKeyPemFilepath)
	if err != nil {
		return fmt.Errorf("error reading public key file: %w", err)
	}
	pubBlock, _ := pem.Decode(pubData)
	if pubBlock == nil {
		return errors.New("failed to decode PEM block from public key file")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse ED25519 public key: %w", err)
	}
	ed25519Pub, ok := pubInterface.(ed25519.PublicKey)
	if !ok {
		return errors.New("provided public key is not an ED25519 key")
	}
	jwtConfig.Ed25519PublicKey = ed25519Pub

	return nil
}
