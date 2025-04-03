package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func validateConfig() error {
	errorMessageTemplate := "no '%s' keyword found in config"

	if err := validateNestedYaml(errorMessageTemplate, "database", []string{"host", "user", "password", "dbname", "sslmode"}); err != nil {
		return err
	}

	if err := validateSingleKeyword(errorMessageTemplate, "domain_name"); err != nil {
		return err
	}

	if err := validateSingleKeyword(errorMessageTemplate, "jwks_endpoint"); err != nil {
		return err
	}

	if err := validateNestedYaml(errorMessageTemplate, "amqp", []string{"host", "user", "password"}); err != nil {
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
