package timezone

import (
	"skeleton/pkg/config"
	"time"
)

func SetupTimezone(cfg *config.TimezoneProperties) {
	location, err := time.LoadLocation(cfg.Location)
	if err != nil {
		panic("Failed to load timezone location: " + err.Error())
	}

	time.Local = location
}
