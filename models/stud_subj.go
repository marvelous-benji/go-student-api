
package models

import (
	"github.com/google/uuid"
	"strings"
	"time"
)



func Id_hex() string{
	raw_uuid := uuid.NewString()
	uuid_hex := strings.ReplaceAll(raw_uuid, "-", "")
	return uuid_hex
}

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