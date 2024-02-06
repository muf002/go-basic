package routes

import (
	"github.com/gorilla/mux"
	"github.com/muf002/go-basic/src/pkg/controllers"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/employee", controllers.GetEmployees).Methods("GET")
	r.HandleFunc("/{employeeId}", controllers.GetEmployeeById).Methods("GET")
	r.HandleFunc("/employee", controllers.Create).Methods("POST")
	r.HandleFunc("/employee", controllers.Update).Methods("PUT")
	r.HandleFunc("/employee/{employeeId}", controllers.Delete).Methods("DELETE")
	return r
}
