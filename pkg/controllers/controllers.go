package controllers

import (
	"encoding/json"
	"io"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muf002/go-basic/pkg/config"
	"github.com/muf002/go-basic/pkg/models"
)

func Basic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You've requested the book \n")
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var emp []models.Employee
	db := config.GetDb()
	db.Find(&emp)
	res, _ := json.Marshal(emp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	params := mux.Vars(r)
	id := params["employeeId"]
	db := config.GetDb()
	db.Find(&emp, id)
	res, _ := json.Marshal(emp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Create(w http.ResponseWriter, r *http.Request) {
	body := models.Employee{}
	json.NewDecoder(r.Body).Decode(&body)
	db := config.GetDb()
	db.Create(&body)
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(body)
	w.Write(res)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["employeeId"]
	db := config.GetDb()
	var employee models.Employee
	db.Unscoped().Delete(&employee, id)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	var empDetails models.Employee
	db := config.GetDb()
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &emp)
	db.Find(&empDetails, emp.ID)
	db.Model(&empDetails).Omit("id").Updates(emp)
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(empDetails)
	w.Write(res)
}
