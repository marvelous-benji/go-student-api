
package controllers

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"views/view"
	"os"
)


func SetRoutes() *mux.Router{
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("students", view.Get_students).Methods("GET")
	r.HandleFunc("student/{id}", view.Get_student).Methods("GET")
	r.HandleFunc("student", view.Add_student).Methods("POST")
	r.HandleFunc("student/{id}", view.Update_student).Methods("PUT")
	r.HandleFunc("student/{id}", view.Delete_student).Methods("DELETE")
	r.HandleFunc("subjects", view.Get_subjects).Methods("GET")
	r.HandleFunc("subject/{id}", view.Get_subject).Methods("GET")
	r.HandleFunc("subject", view.Add_subject).Methods("POST")
	r.HandleFunc("subject/{id}", view.Update_subject).Methods("POST")
	r.HandleFunc("subject/{id}", view.Delete_subject).Methods("POST")
	r.HandleFunc("student/subjects/{stud_id}", view.Get_student_subjects).Methods("GET")
	r.HandleFunc("student/subject/{stud_id}/{sub_id}", view.Remove_stud_sub).Methods("DELETE")

	r.PathPrefix("/api/v1/")
	logrouter := handlers.LoggingHandler(os.Stdout, r)
	return logrouter
}