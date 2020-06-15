package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email                string `gorm:"type:varchar(100);unique_index; not null"`
	Phone                string `gorm:"type:varchar(100);uniq; not null"`
	First_name           string `gorm:"type:varchar(100);"`
	Last_name            string `gorm:"type:varchar(100);"`
	Middle_name          string `gorm:"type:varchar(100);"`
	Age                  int    `gorm:"type:smallint;unique_index"`
	Birthdate            *time.Time
	Role                 int           `gorm:"type:smallint"`
	Directions           []*Direction  `gorm:"many2many:user_directions;"`
	Appointments         []Appointment `gorm:"foreignkey:Patient_id;"`
	AsDoctorAppointments []Appointment `gorm:"foreignkey:Doctor_id;"`
}
