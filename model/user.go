package model

import (
	"errors"
)

type User struct {
	Email string `
	form:"email"
	json:"email" 
	example:"jjh123@naver.com"
	binding:"required"`
	Password string `
	form:"password"
	json:"password" 
	example:"123"
	binding:"required"`
	Name string `
	form:"name" 
	json:"name" 
	example:"jjh"
	binding:"required"`
	PhoneNumber string `
	form:"phone_number" 
	json:"phone_number" 
	example:"010-9966-5942"
	binding:"required"`
	Secretkey string `
	form:"secret_key" 
	json:"secret_key" 
	example:"2y0BcdVYH48Hxc8SEwfOucxAqMoL623K70j6OCWa"
	binding:"required"`
	Accesskey string `
	form:"access_key" 
	json:"access_key" 
	example:"cY158XlCRODQljHva8pMjORsoxrKRdfg4S9jT8qa"
	binding:"required"`
}

//  example
var (
	ErrNameInvalid = errors.New("name is empty")
)
