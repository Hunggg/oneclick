package controller

import (
	cockroachdb "oneclick/services/query/categories/cockroachDB"

	"github.com/gin-gonic/gin"
)

func GetCategoryById(c *gin.Context, db cockroachdb.CockroachDB) {
	id := c.Param("id")
	db.GetCategoryById()
}