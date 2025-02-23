package config

type ServiceConfig struct {
	PathPrefix string
	Name       string
	URL        string
}

func GetServices() []ServiceConfig {
	return []ServiceConfig{
		{PathPrefix: "/auth", Name: "auth", URL: "http://auth-service:8080"},
		{PathPrefix: "/profiles", Name: "profiles", URL: "http://profiles-service:8080"},
		{PathPrefix: "/players", Name: "players", URL: "http://players-service:8080"},
		{PathPrefix: "/games", Name: "games", URL: "http://games-service:8080"},
	}
}
