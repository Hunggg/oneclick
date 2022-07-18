package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"oneclick/config"
	"oneclick/entity"
	cockroachdb "oneclick/services/query/categories/cockroachDB"

	"github.com/gin-gonic/gin"
)

func GetCategoryById(c *gin.Context) {
	id, err := config.ConvertStringToInt(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H {
			"Message" : "Id not found",
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

func GetListCategory(c *gin.Context) {
	offset, err := config.ConvertStringToInt(c.GetHeader("offset"))
	if err != nil {
		c.JSON(500, gin.H {
			"Message" : "offset not found",
		})
	}	
	limit, err := config.ConvertStringToInt(c.GetHeader("limit"))
	if err != nil {
		c.JSON(500, gin.H {
			"Message" : "limit not found",
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
		"total": count,
		"content": category,
	})
}

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
			c.JSON(500, gin.H {
				"message" : e,
			})
		}
		c.JSON(200, gin.H {
			"message": "inserted",
		})
	} else {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}

func SaveBatchCategory(c *gin.Context) {
	var categoryList []entity.Categories
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Body not found",
		})
	}

	err := json.Unmarshal(jsonData, &categoryList)
	db, erDB := config.NewGormDBConnection()
		if erDB != nil {
			log.Fatal(erDB)
		}
	cockroach, er := cockroachdb.NewCockroachDB(db)
		if er != nil {
			log.Fatal(er)
		}

		e := cockroach.SaveBatchCategory()
		if e != nil {
			c.JSON(500, gin.H {
				"message" : e,
			})
		}
	}