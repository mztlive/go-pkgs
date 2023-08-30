package structure

// 这个文件是用来定义一些常用的结构体的

// Collection 集合
type Collection[T any] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}
