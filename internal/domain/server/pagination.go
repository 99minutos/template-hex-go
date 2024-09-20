package server

type PaginationParams struct {
	page  int `json:"page"`
	limit int `json:"limit"`
}

// NewPaginationParams creates a new instance of PaginationParams with the given page and limit
// parameters.
//
// page - the page number starting from 0
// limit - the number of items per page
func NewPaginationParams(page int, limit int) *PaginationParams {
	return &PaginationParams{
		page:  page,
		limit: limit,
	}
}

func (p *PaginationParams) GetPage() *int64 {
	value := int64(p.page)
	return &value
}

func (p *PaginationParams) GetLimit() *int64 {
	value := int64(p.limit)
	return &value
}

func (p *PaginationParams) GetSkip() *int64 {
	value := int64(p.page * p.limit)
	return &value
}
