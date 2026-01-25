package application

import "skeleton/pkg/config"

func (a *Application) CleanUp() {
	if err := a.Database.MasterClient.Close(); err != nil {
		a.Logger.Error("Failed to close database connection: " + err.Error())
	}

	if a.Environment != nil &&
		*a.Environment == config.EnvironmentProduction {
		if err := a.Logger.Sync(); err != nil {
			a.Logger.Warn(
				"Failed to flush logs during clean up: " + err.Error(),
			)
		}
	}
}
