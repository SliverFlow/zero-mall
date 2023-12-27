package request

import "server/common"

type OrderListReq struct {
	common.PageInfo
	Role   int64
	UserID string
}

type PageListReq struct {
	BusinessID string
	Page       int64
	PageSize   int64
	StartTime  int64
	EndTime    int64
	Status     int64
	UserID     string
}

func (r *PageListReq) LimitAndOffset() (int, int) {
	var limit, offset int
	if r.Page == 0 {
		limit = 1
	}
	if r.PageSize == 0 {
		offset = 10
	}

	limit = int(r.PageSize)
	offset = int((r.Page - 1) * r.PageSize)
	return limit, offset
}
