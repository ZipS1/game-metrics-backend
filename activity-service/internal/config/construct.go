package config

func ConstructConfig() (*Config, error) {
	args, err := parseArgs()
	if err != nil {
		return nil, err
	}

	cfg, err := loadConfig(args.ConfigPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
