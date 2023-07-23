package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

type Courses struct {
	cou []Course
}

// fake DB
// var course []Course

// middleware, helper - files
func (c *Course) IsEmpty() bool {
	if c == nil {
		return true
	}
	return c.CourseId == "" && c.CourseName == ""
	// return false
}

func initializeDummyDataDB() *Courses {
	course := []Course{
		{CourseId: "12KA4", CourseName: "Python", CoursePrice: 200, Author: &Author{FullName: "Ramiz Raza", Website: "www.bhosdiwala.com"}},
		{CourseId: "42KA1", CourseName: "JAVA", CoursePrice: 600, Author: &Author{FullName: "Pritam pyare", Website: "www.chotta_sala.com"}},
		{CourseId: "92E18", CourseName: "GoLand", CoursePrice: 50, Author: &Author{FullName: "batla Chodi", Website: "www.bekchodi.com"}},
		{CourseId: "64K24", CourseName: "Jung(Rust)", CoursePrice: 2000, Author: &Author{FullName: "POK Pankaj", Website: "www.fat_ke_flower.com"}},
	}
	return &Courses{cou: course}
}

func main() {
	log.Println("Welcome to simple api server.....")
	http.HandleFunc("/", greeting)
	c := initializeDummyDataDB()
	http.HandleFunc("/courses", c.getAllCourse)
	http.HandleFunc("/course", c.getCourse)
	log.Println("Starting server at localhost:4000...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello lodu welcome to this world of API....</h1>"))
	// fmt.Fprintln(w, "<h1>Hello lodu welcome to this world....</h1>")
}

func (c *Courses) getAllCourse(w http.ResponseWriter, r *http.Request) {

	log.Println("getAllCourse api endpoint hit received")
	w.Header().Add("Content-type", "application/json")
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(c.cou)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(`{"Message":"WrongMethod Called"}`)
	}

}

func (c *Courses) getCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("GetCourse api end hit received")
	w.Header().Add("content-type", "application/json")
	val := r.URL.Query()
	fmt.Fprintf(w, "values received from url is : %v", val)
}
