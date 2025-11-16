package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ScUserPolice struct {
	gorm.Model
	Code     string         `gorm:"column:code;type:varchar(12);not null;" json:"code"`
	PoliceId string         `gorm:"column:police_id;type:varchar(20);not null;unique;" json:"police_id"`
	Password string         `gorm:"column:password;type:varchar(255);not null;" json:"password"`
	Name     string         `gorm:"column:name;type:varchar(50)" json:"name"`
	Works    datatypes.JSON `gorm:"column:works;type:jsonb" json:"works"`
}
