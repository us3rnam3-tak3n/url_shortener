package Model

import "github.com/jinzhu/gorm"

type(
	Todo struct {
		gorm.Model
		Long	string `json:"long"`
		Short	string `json:"short"`
	}
)
