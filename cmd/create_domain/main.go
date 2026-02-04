// Command to create a new domain structure in the project
// Usage: go run ./cmd/create_domain <domain_name>

package main

import (
	"fmt"
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
	"toUpper":      strings.ToUpper,
	"toLowerKebab": utils.ToLowerKebab,
	"toTitle":      utils.ToTitle,
	"toCamel":      utils.ToCamel,
}

func main() {
	if len(os.Args) < 2 {
		panic(
			"Domain name argument is required, Usage: go run ./cmd/create_domain <domain_name>",
		)
	}

	domainName := os.Args[1]
	validateCreateDomainInput(
		CreateDomainInput{
			DomainName: domainName,
		},
	)

	createAPIDir(domainName)
	createRouteFile(domainName)
	createDeliveryLayerFiles(domainName)
	createUsecaseLayerFiles(domainName)
	createRepositoryLayerFiles(domainName)
	createDataLayerFiles(domainName)
	createDomainLayerDir(domainName)
}

func validateCreateDomainInput(input CreateDomainInput) {
	if strings.Trim(input.DomainName, " ") == "" {
		panic("Domain name cannot be empty")
	}

	if utils.HasSpecialChars(input.DomainName) {
		panic("Domain name can only contain alphanumeric characters")
	}
}

func createRouteFile(domainName string) {
	routeSchema := schema.RouteSchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/route.go.tmpl",
		"internal/api/"+domainName+"/route.go",
		routeSchema,
	)
}

func createDeliveryLayerFiles(domainName string) {
	createLayerDir(domainName, "delivery")

	factorySchema := schema.DeliveryFactorySchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/delivery/factory.go.tmpl",
		"internal/api/"+domainName+"/delivery/factory.go",
		factorySchema,
	)

	handlerSchema := schema.DeliveryHandlerSchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/delivery/handler.go.tmpl",
		"internal/api/"+domainName+"/delivery/handler.go",
		handlerSchema,
	)
}

func createUsecaseLayerFiles(domainName string) {
	createLayerDir(domainName, "usecase")

	factorySchema := schema.UsecaseFactorySchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/usecase/factory.go.tmpl",
		"internal/api/"+domainName+"/usecase/factory.go",
		factorySchema,
	)

	functionSchema := schema.UsecaseFunctionSchema{
		DomainName: domainName,
	}
	usecaseFilePath := fmt.Sprintf(
		"internal/api/%s/usecase/get_%s.go",
		domainName,
		domainName,
	)
	createFileFromTemplate(
		"cmd/create_domain/template/usecase/get_domain.go.tmpl",
		usecaseFilePath,
		functionSchema,
	)
}

func createRepositoryLayerFiles(domainName string) {
	createLayerDir(domainName, "repository")

	factorySchema := schema.RepositoryFactorySchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/repository/factory.go.tmpl",
		"internal/api/"+domainName+"/repository/factory.go",
		factorySchema,
	)

	functionSchema := schema.RepositoryFunctionSchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/repository/postgres.go.tmpl",
		"internal/api/"+domainName+"/repository/postgres.go",
		functionSchema,
	)
}

func createDataLayerFiles(domainName string) {
	createLayerDir(domainName, "data")

	dtoSchema := schema.DataDTOSchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/data/dto/example.go.tmpl",
		"internal/api/"+domainName+"/data/dto/"+utils.ToCamel(domainName)+".go",
		dtoSchema,
	)
	createFileFromTemplate(
		"cmd/create_domain/template/data/dto/request.go.tmpl",
		"internal/api/"+domainName+"/data/dto/request.go",
		dtoSchema,
	)
	createFileFromTemplate(
		"cmd/create_domain/template/data/dto/response.go.tmpl",
		"internal/api/"+domainName+"/data/dto/response.go",
		dtoSchema,
	)
}

func createDomainLayerDir(domainName string) {
	createLayerDir(domainName, "domain")

	errsSchema := schema.DomainErrorSchema{
		DomainName: domainName,
	}
	createFileFromTemplate(
		"cmd/create_domain/template/domain/error.go.tmpl",
		"internal/api/"+domainName+"/domain/error.go",
		errsSchema,
	)
}

func createAPIDir(domainName string) {
	baseDirPath := "internal/api/" + domainName
	if err := os.MkdirAll(baseDirPath, 0755); err != nil {
		panic("Could not create API directory: " + err.Error())
	}
}

func createLayerDir(domainName string, layerName string) {
	baseDirPath := "internal/api/" + domainName + "/" + layerName
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
