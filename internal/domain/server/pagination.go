package server

type PaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// NewPaginationParams creates a new instance of PaginationParams with the given page and limit
// parameters.
//
// page - the page number starting from 0
// limit - the number of items per page
func NewPaginationParams(page int, limit int) *PaginationParams {
	return &PaginationParams{
		Page:  page,
		Limit: limit,
	}
}

func (p *PaginationParams) GetPage() *int64 {
	value := int64(p.Page)
	return &value
}

func (p *PaginationParams) GetLimit() *int64 {
	value := int64(p.Limit)
	return &value
}

func (p *PaginationParams) GetSkip() *int64 {
	value := int64(p.Page * p.Limit)
	return &value
}
