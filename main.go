
package main


import (
	"go-student-api/controllers"
	"net/http"
	"time"
	"log"
)


func main(){
	r := router.SetRoute()

	srv := &http.Server{
		Handler: r,
		Addr:  "127.0.0.1:9000",
		WriteTimeout:  15 * time.Second,
		ReadTimeout:   15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}