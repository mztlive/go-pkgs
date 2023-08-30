// 这里是一些常用的结构体定义

package structure

// Paginator 分页参数
type Paginator struct {
	//页码
	Page int64

	//每页数量
	PageSize int64
}

// Offset 返回数据未知偏移
func (p Paginator) Offset() int64 {
	if p.Page <= 0 {
		p.Page = 1
	}

	return (p.Page - 1) * p.PageSize
}

// Limit 返回要取的数据量
func (p Paginator) Limit() int64 {
	if p.PageSize == 0 {
		p.PageSize = 20
	}

	return p.PageSize
}

// NewPaginator 创建一个分页参数
func NewPaginator(page, pageSize int64) Paginator {
	return Paginator{
		Page:     page,
		PageSize: pageSize,
	}
}
