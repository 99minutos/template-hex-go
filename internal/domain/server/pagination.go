package server

import "github.com/bytedance/sonic"

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

// GenericPaginationQuery - a generic query struct that can be used to query any type of data
type GenericPaginationQuery struct {
	Query      interface{}       `json:"query"`
	Pagination *PaginationParams `json:"pagination"`
}

func (a *GenericPaginationQuery) DecodeQuery(someType interface{}) error {
	jsonString, err := sonic.MarshalString(a.Query)
	if err != nil {
		return err
	}
	return sonic.UnmarshalString(jsonString, someType)
}

type GenericPaginationResponse struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
