package request

import (
	"administrasi/models"

	"github.com/gin-gonic/gin"
)

const (
	defaultLimit = 20
	maxLimit     = 1000
)

// Paginate validates pagination requests
func Paginate(c *gin.Context) (*models.Pagination, error) {
	p := new(models.Pagination)
	if err := c.ShouldBindQuery(p); err != nil {
		return nil, err
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
	if p.Limit > 1000 {
		p.Limit = maxLimit
	}
	p.Offset = (p.Page - 1) * p.Limit
	return p, nil
}
