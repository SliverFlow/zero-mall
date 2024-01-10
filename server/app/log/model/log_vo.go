package model

type LogPageReq struct {
	Page       int64
	PageSize   int64
	Username   string
	Method     string
	Path       string
	StatusCode int64
	StartTime  int64
	EndTime    int64
}
