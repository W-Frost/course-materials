package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
}
//added the response for class
type ClassesResponse struct{
	Classes []Class `json:"classes"`
}

type Assignment struct {
	Id string `json:"id"`
	Title string `json:"title`
	Description string `json:"desc"`
	Points int `json:"points"`
}
//added the struct for class
type Class struct {
	Id string `json:"id"`
	Name string `json:"name`
	CourseDescription string `json:"course desc"`
	Credits int `json:"credits"`
}

var Classes []Class
var Assignments []Assignment
const Valkey string = "FooKey"

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	
	var response Response
	response.Assignments = Assignments
	params := mux.Vars(r)
	var assignmnet Assignment	
	r.ParseForm()
	for index, assignment := range Assignments {
		if assignment.Id == params["id"]{
			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			assignmnet.Id =  r.FormValue("id")
		    assignmnet.Title =  r.FormValue("title")
			assignmnet.Description =  r.FormValue("desc")
			assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
			Assignments = append(Assignments, assignmnet)
			break
		}
	}


}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if(r.FormValue("id") != ""){
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)
}

//Creates the classes
func CreateClass(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var class Class
	r.ParseForm()
	if(r.FormValue("id") !=""){
		class.Id =r.FormValue("id")
		class.Name = r.FormValue("name")
		class.CourseDescription = r.FormValue("course desc")
		class.Credits, _ = strconv.Atoi(r.FormValue("credits"))
		Classes = append(Classes, class)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)
}