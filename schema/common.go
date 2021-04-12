package schema

import (
	"strconv"
	"strings"
)

type PaginationReq struct {
	Page    int64 `form:"page" binding:""`
	PerPage int64 `form:"perPage" binding:""`
	Total   int64 `form:"total" binding:""`
}

type PaginationResp struct {
	Data  interface{} `form:"data"`
	Total int64       `form:"Total"`
}

type OrderBy string

func (a OrderBy) Values() (map[string]int, error) {
	values := make(map[string]int)
	s := string(a)
	elems := strings.Split(s, ",")
	for _, elem := range elems {
		kvs := strings.Split(elem, "=")
		if len(kvs) != 2 {
			return nil, ErrInvalidArguments
		}
		order, err := strconv.Atoi(kvs[1])
		if err != nil || (order != 1 && order != -1) {
			return nil, ErrInvalidArguments
		}
		values[kvs[0]] = order
	}
	return values, nil
}
