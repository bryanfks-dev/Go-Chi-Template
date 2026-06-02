package utils

import "skeleton/pkg/consts"

func GetOffsetLimitFromPageSize(page int, size int) (*int, *int) {
	if page <= 0 || size <= 0 {
		return nil, nil
	}
	offset := (page - 1) * size
	limit := size
	return &offset, &limit
}

func SanitizePageSize(page *int, size *int) (*int, *int) {
	sanPage := consts.DefaultPaginationPage
	sanSize := consts.DefaultPaginationSize
	if page != nil && *page > 0 {
		sanPage = *page
	}
	if size != nil && *size > 0 {
		sanSize = *size
	}
	return &sanPage, &sanSize
}
