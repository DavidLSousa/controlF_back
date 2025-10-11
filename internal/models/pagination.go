package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func NewPagination(c *gin.Context) *Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	return &Pagination{
		Limit: limit,
		Page:  page,
	}
}

func DefaultPagination() *Pagination {
	return &Pagination{
		Limit: 10,
		Page:  1,
	}
}

func (p *Pagination) GetScope(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	return db.Offset(offset).
		Limit(p.Limit)
}
