package model

// Pagination 分页请求参数
type Pagination struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Status   string `json:"status,omitempty"`
}

// PaginatedResult 分页响应结构 (用于嵌入其他数据)
type PaginatedResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
