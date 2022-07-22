package controller

import (
	"log"
	"oneclick/config"
	"oneclick/entity"
	"oneclick/entity/model"
	cockroachdb "oneclick/services/query/categories/cockroachDB"

	"github.com/gin-gonic/gin"
)

// Category godoc
// @Summary      Category
// @Description  get category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        id		path	int		false	"category_id"
// @Success      200  {array}   entity.Categories
// @Failure      500  {object}  string
// @Router       /v1/categories/{id} [get]
func GetCategoryById(c *gin.Context) {
	id, err := config.ConvertStringToInt(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "Id not found",
		})
	}

	db, erDB := config.NewGormDBConnection()
	if erDB != nil {
		log.Fatal(erDB)
	}
	cockroach, er := cockroachdb.NewCockroachDB(db)
	if er != nil {
		log.Fatal(er)
	}
	category, er := cockroach.GetCategoryById(uint64(id))
	if er != nil {
		log.Fatal(er)
	}
	c.JSON(200, category)
}
// Category godoc
// @Summary      ListCategory
// @Description  get list category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        offset		header	int		false	"offset"
// @Param        limit		header	int		false	"limit"
// @Success      200  {array}   entity.Categories
// @Failure      500  {object}  string
// @Router       /v1/categories/list [get]
func GetListCategory(c *gin.Context) {
	offset, err := config.ConvertStringToInt(c.GetHeader("offset"))
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "offset not found",
		})
	}
	limit, err := config.ConvertStringToInt(c.GetHeader("limit"))
	if err != nil {
		c.JSON(500, gin.H{
			"Message": "limit not found",
		})
	}

	db, erDB := config.NewGormDBConnection()
	if erDB != nil {
		log.Fatal(erDB)
	}
	cockroach, er := cockroachdb.NewCockroachDB(db)
	if er != nil {
		log.Fatal(er)
	}
	category, count, er := cockroach.GetListCategory(offset, limit)
	if er != nil {
		log.Fatal(er)
	}
	c.JSON(200, gin.H{
		"total":   count,
		"content": category,
	})
}
// Category godoc
// @Summary      Category
// @Description  save category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        body	body	entity.Categories	false	"category"
// @Success      200  {array}   entity.Categories
// @Failure      500  {object}  string
// @Router       /v1/categories/create [post]
func SaveCategory(c *gin.Context) {
	var json entity.Categories

	if err := c.ShouldBindJSON(&json); err == nil {
		db, erDB := config.NewGormDBConnection()
		if erDB != nil {
			log.Fatal(erDB)
		}
		cockroach, er := cockroachdb.NewCockroachDB(db)
		if er != nil {
			log.Fatal(er)
		}

		e := cockroach.SaveCategory(json)
		if e != nil {
			c.JSON(500, gin.H{
				"message": e,
			})
		}
		c.JSON(200, gin.H{
			"message": "inserted",
		})
	} else {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}

// Category godoc
// @Summary      Category
// @Description  update category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        body		body	entity.Categories		false	"category"
// @Success      200  {array}   entity.Categories
// @Failure      500  {object}  string
// @Router       /v1/categories/update [patch]
func UpdateCategory(c *gin.Context) {
	var json entity.Categories

	if err := c.ShouldBindJSON(&json); err == nil {
		db, erDB := config.NewGormDBConnection()
		if erDB != nil {
			log.Fatal(erDB)
		}
		cockroach, er := cockroachdb.NewCockroachDB(db)
		if er != nil {
			log.Fatal(er)
		}

		e := cockroach.UpdateCategory(json)
		if e != nil {
			c.JSON(500, gin.H{
				"message": e,
			})
		}
		c.JSON(200, gin.H{
			"message": "updated",
		})
	} else {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}

// Category godoc
// @Summary      Category
// @Description  delete category
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        id		path	int		false	"id"
// @Success      200  {array}   entity.Categories
// @Failure      500  {object}  string
// @Router       /v1/categories/delete/{id} [delete]
func DeleteCategoryById(c *gin.Context) {
	id, err := config.ConvertStringToInt(c.Param("id"))
	if err != nil {
		responseErr := model.BuildErrorResponse("Id not found", err.Error(), model.EmptyObj{})
		c.JSON(500, responseErr)
	}

	db, erDB := config.NewGormDBConnection()
	if erDB != nil {
		log.Fatal(erDB)
	}
	cockroach, er := cockroachdb.NewCockroachDB(db)
	if er != nil {
		log.Fatal(er)
	}

	if err := cockroach.DeleteCategoryById(uint64(id)); err != nil {
		c.JSON(500, gin.H{
			"Error query DB": err,
		})
	}
	c.JSON(200, gin.H{
		"Message": "Deleted",
	})
}
