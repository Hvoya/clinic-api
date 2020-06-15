package model

import "github.com/jinzhu/gorm"

type Direction struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100);unique_index"`
	Users []*User `gorm:"many2many:user_directions;"`
}
