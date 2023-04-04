package paginator

import (
	"administrasi/models"
	"math"
)

func Paging(p *models.Pagination) *models.Pagination {
	p.TotalPage = int(math.Ceil(float64(p.Count) / float64(p.Limit)))

	if p.Page > 1 {
		p.PreviousPage = p.Page - 1
	} else {
		p.PreviousPage = p.Page
	}

	if p.Page == p.TotalPage {
		p.NextPage = p.Page
	} else {
		p.NextPage = p.Page + 1
	}
	return p
}
