package dataanalysis

import (
	"context"

	"github.com/mztlive/go-pkgs/wechat/mini"
)

// RetainItem 留存项结构
type RetainItem struct {
	Key   int `json:"key"`   // 标识，0开始表示当天，1表示1甜后，以此类推
	Value int `json:"value"` // key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

type ResAnalysisRetain struct {
	mini.CommonError
	RefDate    string       `json:"ref_date"`     // 日期
	VisitUVNew []RetainItem `json:"visit_uv_new"` // 新增用户留存
	VisitUV    []RetainItem `json:"visit_uv"`     // 活跃用户留存
}

type Analysis interface {
	GetAnalysisDailyRetain(ctx context.Context, beginDate, endDate string) (*ResAnalysisRetain, error)
}
