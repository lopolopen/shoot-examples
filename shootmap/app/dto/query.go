package dto

type Pagination struct {
	Page    int `json:page`
	PerPage int `json:perPage`
}
