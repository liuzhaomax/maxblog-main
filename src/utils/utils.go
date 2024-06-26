package utils

import (
	"github.com/liuzhaomax/maxblog-main/internal/core"
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
	if pageNoReq == core.EmptyString {
		pageNoReq = "1"
	}
	if pageSizeReq == core.EmptyString {
		pageSizeReq = "5"
	}
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
		pageSize = 5
	case pageSize >= 30:
		pageSize = 30
	}
	return pageNo, pageSize, nil
}
