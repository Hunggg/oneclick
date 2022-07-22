package model

type AuthLog struct {
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string ``
}