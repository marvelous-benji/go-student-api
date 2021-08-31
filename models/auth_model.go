
package models


import (
	"time"
	"github.com/google/uuid"
	"strings"
)


func Id_hex() string{
	raw_uuid := uuid.NewString()
	uuid_hex := strings.ReplaceAll(raw_uuid,"-","")
	return uuid_hex
}


type User struct{
	ID string `gorm:"primaryKey" sql:"DEFAULT:Id_hex"`
	Email string `gorm:"index:unique"`
	Password string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}