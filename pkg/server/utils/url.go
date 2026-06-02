package serverutils

import (
	"net/http"
	"skeleton/pkg/consts"
	"skeleton/pkg/utils"
)

func GetPaginationQueryParams(r *http.Request) (*int, *int) {
	query := r.URL.Query()
	pageQuery := query.Get(consts.PaginationPageQueryParamKey)
	sizeQuery := query.Get(consts.PaginationSizeQueryParamKey)

	var page, size *int
	if pageQuery != "" {
		pageVal := utils.ParseIntOrDefault(
			pageQuery,
			consts.DefaultPaginationPage,
		)
		page = &pageVal
	}
	if sizeQuery != "" {
		sizeVal := utils.ParseIntOrDefault(
			sizeQuery,
			consts.DefaultPaginationSize,
		)
		size = &sizeVal
	}
	return utils.SanitizePageSize(page, size)
}
