
package models

import (
	"time"
)



type Student struct{
	ID  string `gorm: "primaryKey" sql:"DEFAULT:id_hex"`
	Name string
	DateCreated time.Time `gorm:"autoCreateTime"`
	DateUpdated time.Time `gorm:"autoUpdateTime"`
	Subjects []*Subject `gorm:"many2many:student_subjects;"`
}


type Subject struct{
	ID string `gorm:"primaryKey" sql:"DEFAULT:id_hex"`
	SubName string
	SubCode int
	DateCreated time.Time `gorm:"autoCreateTime"`
	OfferedBy []*Student `gorm:"many2many:student_subjects;"`
}