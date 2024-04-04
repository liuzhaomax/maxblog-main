package utils

import (
	"strconv"
)

func Str2Uint32(str string) (uint32, error) {
	if str == "" {
		str = "0"
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}

func Paginate(pageNoReq string, pageSizeReq string) (int, int, error) {
	pageNo, err := strconv.Atoi(pageNoReq)
	if err != nil {
		return 0, 0, err
	}
	pageSize, err := strconv.Atoi(pageSizeReq)
	if err != nil {
		return 0, 0, err
	}
	if pageNo <= 0 {
		pageNo = 1
	}
	switch {
	case pageSize <= 0:
		pageSize = 10
	case pageSize >= 100:
		pageSize = 100
	}
	return pageNo, pageSize, nil
}
