package qpo

type Pagination struct {
	Page    int
	PerPage int
}

func (p Pagination) Offset() int {
	if p.Page <= 0 || p.PerPage <= 0 {
		return 0
	}
	return p.PerPage * (p.Page - 1)
}
