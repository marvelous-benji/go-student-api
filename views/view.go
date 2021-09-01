
package views

import (
	"github.com/gorilla/mux"
	"go-student-api/models"
	"gorm.io/gorm"
	//"go-student-api/controllers"
	"net/http"
	"encoding/json"
	"log"
)



var db *gorm.DB

db = db_conn.InitDb()



type student struct{
	ID  string `json:"id"`
	Name string `json:"name"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time  `json:"date_updated"`
	Subjects []*Subject    `json:"subjects"`
}

type subject struct{
	ID string  `json:"id"`
	SubName string  `json:"sub_name"`
	SubCode int  `json:"sub_code"`
	DateCreated time.Time `json:"date_created"`
	OfferedBy []*Student   `json:"offered_by"`
}


db.Automigrate(&student{}, &subject{})

func Get_students(w *http.ResponseWriter, r *http.Request){
	students := []student{}
	result := db.Find(&students)
	w.Header.Set("Content-Type", "application/json")
	res, err := json.Marshal(result)
	if err != nil{
		log.Fatalf("Error: %s occured", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}