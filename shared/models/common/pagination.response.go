package common

type Pagination struct {
	PageSize   int         `json:"pageSize"`
	PageNumber int         `json:"pageNumber"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
}

func (p *Pagination) GetOffset() int {
	return (p.PageNumber - 1) * p.PageSize
}

func (p *Pagination) GetLimit() int {
	if p.PageSize == 0 {
		return 10
	}
	return p.PageSize
}
