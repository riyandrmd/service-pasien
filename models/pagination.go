package models

type Pagination struct {
	Limit        int `json:"limit,omitempty" form:"limit"`
	Offset       int `json:"-"`
	Page         int `json:"page,omitempty" form:"page" binding:"min=0"`
	NextPage     int `json:"next_page,omitempty"`
	PreviousPage int `json:"previous_page,omitempty"`
	Count        int `json:"count,omitempty"`
	TotalPage    int `json:"total_page,omitempty"`
}