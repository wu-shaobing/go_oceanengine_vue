package utils

// Pagination 分页参数
type Pagination struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

// GetPage 获取页码
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

// GetPageSize 获取每页数量
func (p *Pagination) GetPageSize() int {
	if p.PageSize <= 0 {
		return 10
	}
	if p.PageSize > 100 {
		return 100
	}
	return p.PageSize
}

// GetOffset 获取偏移量
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

// GetLimit 获取限制数量
func (p *Pagination) GetLimit() int {
	return p.GetPageSize()
}

// DefaultPagination 默认分页
func DefaultPagination() *Pagination {
	return &Pagination{
		Page:     1,
		PageSize: 10,
	}
}

// NewPagination 创建分页
func NewPagination(page, pageSize int) *Pagination {
	p := &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
	p.Page = p.GetPage()
	p.PageSize = p.GetPageSize()
	return p
}

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[i%len(charset)]
	}
	return string(b)
}
