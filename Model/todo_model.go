package Model

import "github.com/jinzhu/gorm"

//Structure to create table on Database.
type(
	Todo struct {
		gorm.Model
		Long	string `json:"long"`
		Short	string `json:"short"`
	}
)
