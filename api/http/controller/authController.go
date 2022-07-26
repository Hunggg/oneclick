package controller

import (
	"log"
	"net/http"
	"oneclick/config"
	"oneclick/entity"
	"oneclick/entity/model"
	cockroachdb "oneclick/services/query/account/cockroachDB"

	"github.com/gin-gonic/gin"
)

// Accout godoc
// @Summary      Account
// @Description  save Account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        Authentication	header	entity.Accounts	false	"Account"
// @Param        body	body	entity.Accounts	false	"Account"
// @Success      200  {array}   entity.Accounts
// @Failure      500  {object}  string
// @Router       /v1/account/create [post]
func Register(c *gin.Context) {
	var json entity.Accounts

	if err := c.ShouldBindJSON(&json); err == nil {
		db, erDB := config.NewGormDBConnection()
		if erDB != nil {
			log.Fatal(erDB)
		}
		cockroach, er := cockroachdb.NewCockroachDB(db)
		if er != nil {
			log.Fatal(er)
		}

		e := cockroach.SaveAccount(json)
		if e != nil {
			errorResponse := model.BuildErrorResponse("Not saved account", e.Error(), model.EmptyObj{})
			c.JSON(http.StatusRequestTimeout, errorResponse)
		}
		response := model.BuildResponse(true, "inserted", json)
		c.JSON(http.StatusOK, response)
	} else {
		errorResponse := model.BuildErrorResponse("Not saved account", err.Error(), model.EmptyObj{})
		c.JSON(http.StatusRequestTimeout, errorResponse)
	}
}

// func Login(c *gin.Context) {
// 	var json model.AuthLog

// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		db, erDB := config.NewGormDBConnection()
// 		if erDB != nil {
// 			log.Fatal(erDB)
// 		}
// 		cockroach, er := cockroachdb.NewCockroachDB(db)
// 		if er != nil {
// 			log.Fatal(er)
// 		}
// 		str, e := cockroach.Login(json)
// 		if e != nil {
// 			errorResponse := model.BuildErrorResponse("Have no accout", e.Error(), model.EmptyObj{})
// 			c.JSON(http.StatusRequestTimeout, errorResponse)
// 		}
// 		response := model.BuildResponse(true, "login success", str)
// 		c.JSON(http.StatusOK, response)
// 	} else {
// 		errorResponse := model.BuildErrorResponse("Not login", err.Error(), model.EmptyObj{})
// 		c.JSON(http.StatusRequestTimeout, errorResponse)
// 	}
// } 