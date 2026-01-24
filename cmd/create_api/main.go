// Command to create a new API service using net/http in Go
// Usage: go run ./cmd/create_api <service_name>

package main

import (
	"os"
	"path/filepath"
	"skeleton/cmd/create_api/data/schema"
	"skeleton/cmd/create_api/utils"
	"strings"
	"text/template"
)

type CreateAPIInput struct {
	APIName string
}

var funcMap = template.FuncMap{
	"toLower": strings.ToLower,
	"toTitle": utils.ToTitle,
}

func main() {
	apiName := os.Args[1]
	validateCreateAPIInput(
		CreateAPIInput{
			APIName: apiName,
		},
	)

	createAPIDir(apiName)
	createRouteFile(apiName)
	createDeliveryLayerFiles(apiName)
	createUsecaseLayerFiles(apiName)
	createRepositoryLayerFiles(apiName)
	createDataLayerFiles(apiName)
	createDomainLayerDir(apiName)
}

func validateCreateAPIInput(input CreateAPIInput) {
	if strings.Trim(input.APIName, " ") == "" {
		panic("API name cannot be empty")
	}

	if utils.HasSpecialChars(input.APIName) {
		panic("API name can only contain alphanumeric characters")
	}
}

func createRouteFile(apiName string) {
	routeSchema := schema.RouteSchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/route.go.tmpl",
		"internal/api/"+apiName+"/route.go",
		routeSchema,
	)
}

func createDeliveryLayerFiles(apiName string) {
	createLayerDir(apiName, "delivery")

	factorySchema := schema.DeliveryFactorySchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/delivery/factory.go.tmpl",
		"internal/api/"+apiName+"/delivery/factory.go",
		factorySchema,
	)

	handlerSchema := schema.DeliveryHandlerSchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/delivery/handler.go.tmpl",
		"internal/api/"+apiName+"/delivery/handler.go",
		handlerSchema,
	)
}

func createUsecaseLayerFiles(apiName string) {
	createLayerDir(apiName, "usecase")

	factorySchema := schema.UsecaseFactorySchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/usecase/factory.go.tmpl",
		"internal/api/"+apiName+"/usecase/factory.go",
		factorySchema,
	)

	functionSchema := schema.UsecaseFunctionSchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/usecase/read.go.tmpl",
		"internal/api/"+apiName+"/usecase/read.go",
		functionSchema,
	)
}

func createRepositoryLayerFiles(apiName string) {
	createLayerDir(apiName, "repository")

	factorySchema := schema.RepositoryFactorySchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/repository/factory.go.tmpl",
		"internal/api/"+apiName+"/repository/factory.go",
		factorySchema,
	)

	functionSchema := schema.RepositoryFunctionSchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/repository/postgres_example.go.tmpl",
		"internal/api/"+apiName+"/repository/postgres_example.go",
		functionSchema,
	)
}

func createDataLayerFiles(apiName string) {
	createLayerDir(apiName, "data")

	dtoSchema := schema.DataDTOSchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/data/dto/example.go.tmpl",
		"internal/api/"+apiName+"/data/dto/example.go",
		dtoSchema,
	)
	createFileFromTemplate(
		"cmd/create_api/template/data/dto/request.go.tmpl",
		"internal/api/"+apiName+"/data/dto/request.go",
		dtoSchema,
	)
	createFileFromTemplate(
		"cmd/create_api/template/data/dto/response.go.tmpl",
		"internal/api/"+apiName+"/data/dto/response.go",
		dtoSchema,
	)
}

func createDomainLayerDir(apiName string) {
	createLayerDir(apiName, "domain")

	entitySchema := schema.DomainEntitySchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/domain/entity.go.tmpl",
		"internal/api/"+apiName+"/domain/entity.go",
		entitySchema,
	)

	errsSchema := schema.DomainErrorsSchema{
		APIName: apiName,
	}
	createFileFromTemplate(
		"cmd/create_api/template/domain/errors.go.tmpl",
		"internal/api/"+apiName+"/domain/errors.go",
		errsSchema,
	)
}

func createAPIDir(apiName string) {
	baseDirPath := "internal/api/" + apiName
	if err := os.MkdirAll(baseDirPath, 0755); err != nil {
		panic("Could not create API directory: " + err.Error())
	}
}

func createLayerDir(apiName string, layerName string) {
	baseDirPath := "internal/api/" + apiName + "/" + layerName
	if err := os.MkdirAll(baseDirPath, 0755); err != nil {
		panic("Could not create layer directory: " + err.Error())
	}
}

func createFileFromTemplate(
	templFilePath string,
	destFilePath string,
	schema any,
) {
	if _, err := os.Stat(templFilePath); err != nil {
		if !os.IsNotExist(err) {
			panic("Could not access template file: " + err.Error())
		}
	}

	baseTemplFileName := templFilePath[strings.LastIndex(
		templFilePath,
		"/",
	)+1:]
	templ, err := template.New(baseTemplFileName).Funcs(funcMap).ParseFiles(
		templFilePath,
	)
	if err != nil {
		panic("Could not parse template file: " + err.Error())
	}

	if _, err := os.Stat(destFilePath); err == nil {
		panic(
			"Destination file already exists: " +
				destFilePath + " Please remove the file first.",
		)
	}

	destDir := filepath.Dir(destFilePath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		panic("Could not create destination file directory: " + err.Error())
	}

	destFile, err := os.Create(destFilePath)
	if err != nil {
		panic("Could not create destination file: " + err.Error())
	}
	defer destFile.Close()

	if err := templ.Execute(destFile, schema); err != nil {
		panic("Could not execute template: " + err.Error())
	}
}
