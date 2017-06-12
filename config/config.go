package config

import (
	"encoding/json"
)

type DatabaseConfig struct {
	Type uint8 `json:"Type"`
	Host string `json:"Host"`
	Name string `json:"Database"`	
}

type RegistrationConfig struct {
	SelfRegistrationEnabled bool `json:"SelfRegistrationEnabled"`
	RequireConfirmationResponse bool `json:"RequireConfirmationResponse`
	SendConfirmationEmail bool `json:"SendConfirmationEmail"`
}

type AppConfig struct {
	ListenAddress string `json:"ListenAddress"`
	LogFile string `json:"LogFile"`

	Database DatabaseConfig `json:"Database"`
	Registration RegistrationConfig `json:"Registration"`
}

var SystemConfig AppConfig

func LoadConfig () (error) {
	unMarshalAppConfig := func (configData []byte) error {
		return json.Unmarshal(configData, &SystemConfig)
	}

	return Load(unMarshalAppConfig)
}
