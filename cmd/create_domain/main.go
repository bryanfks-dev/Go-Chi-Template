// Command to create a new domain structure in the project
// Usage: go run ./cmd/create_domain <domain_name>

package main

import (
	"os"
	"path/filepath"
	"skeleton/cmd/create_domain/data/schema"
	"skeleton/cmd/create_domain/utils"
	"strings"
	"text/template"
)

type CreateDomainInput struct {
	DomainName string
}

var funcMap = template.FuncMap{
	"normalize":    utils.Normalize,
	"toLower":      strings.ToLower,
	"toLowerKebab": utils.ToLowerKebab,
	"toTitle":      utils.ToTitle,
	"toCamel":      utils.ToCamel,
}

func main() {
	DomainName := os.Args[1]
	validateCreateDomainInput(
		CreateDomainInput{
			DomainName: DomainName,
		},
	)

	createAPIDir(DomainName)
	createRouteFile(DomainName)
	createDeliveryLayerFiles(DomainName)
	createUsecaseLayerFiles(DomainName)
	createRepositoryLayerFiles(DomainName)
	createDataLayerFiles(DomainName)
	createDomainLayerDir(DomainName)
}

func validateCreateDomainInput(input CreateDomainInput) {
	if strings.Trim(input.DomainName, " ") == "" {
		panic("Domain name cannot be empty")
	}

	if utils.HasSpecialChars(input.DomainName) {
		panic("Domain name can only contain alphanumeric characters")
	}
}

func createRouteFile(DomainName string) {
	routeSchema := schema.RouteSchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/route.go.tmpl",
		"internal/api/"+DomainName+"/route.go",
		routeSchema,
	)
}

func createDeliveryLayerFiles(DomainName string) {
	createLayerDir(DomainName, "delivery")

	factorySchema := schema.DeliveryFactorySchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/delivery/factory.go.tmpl",
		"internal/api/"+DomainName+"/delivery/factory.go",
		factorySchema,
	)

	handlerSchema := schema.DeliveryHandlerSchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/delivery/handler.go.tmpl",
		"internal/api/"+DomainName+"/delivery/handler.go",
		handlerSchema,
	)
}

func createUsecaseLayerFiles(DomainName string) {
	createLayerDir(DomainName, "usecase")

	factorySchema := schema.UsecaseFactorySchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/usecase/factory.go.tmpl",
		"internal/api/"+DomainName+"/usecase/factory.go",
		factorySchema,
	)

	functionSchema := schema.UsecaseFunctionSchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/usecase/read.go.tmpl",
		"internal/api/"+DomainName+"/usecase/read.go",
		functionSchema,
	)
}

func createRepositoryLayerFiles(DomainName string) {
	createLayerDir(DomainName, "repository")

	factorySchema := schema.RepositoryFactorySchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/repository/factory.go.tmpl",
		"internal/api/"+DomainName+"/repository/factory.go",
		factorySchema,
	)

	functionSchema := schema.RepositoryFunctionSchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/repository/postgres_example.go.tmpl",
		"internal/api/"+DomainName+"/repository/postgres_example.go",
		functionSchema,
	)
}

func createDataLayerFiles(DomainName string) {
	createLayerDir(DomainName, "data")

	dtoSchema := schema.DataDTOSchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/data/dto/example.go.tmpl",
		"internal/api/"+DomainName+"/data/dto/"+utils.ToCamel(DomainName)+".go",
		dtoSchema,
	)
	createFileFromTemplate(
		"cmd/create_domain/template/data/dto/request.go.tmpl",
		"internal/api/"+DomainName+"/data/dto/request.go",
		dtoSchema,
	)
	createFileFromTemplate(
		"cmd/create_domain/template/data/dto/response.go.tmpl",
		"internal/api/"+DomainName+"/data/dto/response.go",
		dtoSchema,
	)
}

func createDomainLayerDir(DomainName string) {
	createLayerDir(DomainName, "domain")

	errsSchema := schema.DomainErrorsSchema{
		DomainName: DomainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/domain/errors.go.tmpl",
		"internal/api/"+DomainName+"/domain/errors.go",
		errsSchema,
	)
}

func createAPIDir(DomainName string) {
	baseDirPath := "internal/api/" + DomainName
	if err := os.MkdirAll(baseDirPath, 0755); err != nil {
		panic("Could not create API directory: " + err.Error())
	}
}

func createLayerDir(DomainName string, layerName string) {
	baseDirPath := "internal/api/" + DomainName + "/" + layerName
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
