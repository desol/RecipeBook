package settings

import (
	"encoding/json"
	"os"
	"time"
)

type unCleanSettings struct {
	StormDB           string `json:"StormDB"`
	Port              string `json:"Port"`
	ServerTimeout     int    `json:"ServerTimeout"`
	ServerIdleTimeout int    `json:"ServerIdleTimeout"`
}

type settings struct {
	StormDB           string
	Port              string
	ServerTimeout     time.Duration
	ServerIdleTimeout time.Duration
}

// Settings the application settings for the current instance.
var Settings settings

// Populate : Sets the values for the Settings object.
func Populate(debug bool) error {
	var settingsPath string
	var tempSettings unCleanSettings

	if debug {
		settingsPath = "./settings/settings.development.json"
	} else {
		settingsPath = "./settings/settings.production.json"
	}

	settingsFile, err := os.Open(settingsPath)
	if err != nil {
		return err
	}
	defer settingsFile.Close()

	decoder := json.NewDecoder(settingsFile)
	err = decoder.Decode(&tempSettings)
	if err != nil {
		return err
	}

	cleanSettings(tempSettings)

	return nil
}

func cleanSettings(temp unCleanSettings) {
	Settings.Port = temp.Port
	Settings.StormDB = temp.StormDB
	Settings.ServerTimeout = time.Duration(temp.ServerTimeout) * time.Minute
	Settings.ServerIdleTimeout = time.Duration(temp.ServerIdleTimeout) * time.Minute
}
