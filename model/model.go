package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	EmployeeId int    `json:"employee_id"`
	Name       string `json:"name"`
	Lastname   string `json:"lastname"`
	Birhtday   string `json:"birhtday"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
}
