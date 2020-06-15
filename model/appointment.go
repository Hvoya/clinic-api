package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	date       *time.Time
	Patient_id int
	Doctor_id  int
}
