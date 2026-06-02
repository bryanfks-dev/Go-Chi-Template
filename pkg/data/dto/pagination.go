package basedto

import (
	"math"
	"reflect"
	"skeleton/pkg/utils"
)

type PaginationDTO struct {
	Page int
	Size int
}

func NewPaginationDTOFromQueryParamDTO(qp any) *PaginationDTO {
	typ := reflect.TypeOf(qp)
	isPointerOfStruct := typ.Kind() == reflect.Pointer &&
		typ.Elem().Kind() == reflect.Struct
	isStruct := typ.Kind() == reflect.Struct
	if !isPointerOfStruct && !isStruct {
		return nil
	}

	var val reflect.Value
	if isPointerOfStruct {
		val = reflect.ValueOf(qp).Elem()
	} else {
		val = reflect.ValueOf(qp)
	}
	if !val.IsValid() {
		return nil
	}

	pageField := val.FieldByName("Page")
	sizeField := val.FieldByName("Size")

	var page, size int
	if pageField.IsValid() && !pageField.IsNil() {
		page = int(pageField.Elem().Int())
	}
	if sizeField.IsValid() && !sizeField.IsNil() {
		size = int(sizeField.Elem().Int())
	}

	sanPage, sanSize := utils.SanitizePageSize(&page, &size)
	if sanPage == nil || sanSize == nil {
		return nil
	}
	return &PaginationDTO{
		Page: *sanPage,
		Size: *sanSize,
	}
}

type PaginationResultDTO struct {
	TotalItems  int
	SizePerPage int
}

func (d *PaginationResultDTO) GetTotalPages() int {
	return int(math.Ceil(float64(d.TotalItems) / float64(d.SizePerPage)))
}
