package routes

import (
	"github.com/gorilla/mux"
	"github.com/muf002/go-basic/src/pkg/controllers"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetEmployees).Methods("GET")
	r.HandleFunc("/{employeeId}", controllers.GetEmployeeById).Methods("GET")
	r.HandleFunc("/create", controllers.Create).Methods("POST")
	r.HandleFunc("/update", controllers.Update).Methods("PUT")
	r.HandleFunc("/delete/{employeeId}", controllers.Delete).Methods("DELETE")
	return r
}
