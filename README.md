# GOLANG NET/HTTP SERVICE TEMPLATE

## APPLICATION ENVIRONMENT

To uses this application in production environment, set the `APP_ENVIRONMENT` environment variable to `production`. This will enable production-specific configurations and optimizations.

## CONFIG FILES

This directory contains configuration files for the Go project. These files are used to set up and manage various aspects of the service, including environment variables, logging settings, and other configuration parameters.

### Considerations

- Keep configuration files organized and well-documented to facilitate easy maintenance and updates.
- Syncronize configuration files with the deployment environment to avoid discrepancies between development, staging, and production settings.
- After modifying configuration files, ensure to modify the `internal/config/factory.go` file to reflect the changes in the application's configuration loading mechanism.
- Reload or restart the service after making changes to configuration files to apply the new settings.
