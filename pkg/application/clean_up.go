package application

import "skeleton/pkg/config"

func (a *Application) CleanUp() {
	if err := a.Db.MasterClient.Close(); err != nil {
		a.Logger.Error("Failed to close database connection: " + err.Error())
	}

	if a.Env == config.EnvironmentProduction {
		if err := a.Logger.Sync(); err != nil {
			a.Logger.Warn(
				"Failed to flush logs during clean up: " + err.Error(),
			)
		}
	}
}
