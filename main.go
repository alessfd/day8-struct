package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project/{id}", projectDetail).Methods("GET")
	route.HandleFunc("/addproject", addProject).Methods("GET")
	route.HandleFunc("/submitproject", submitProject).Methods("POST")
	route.HandleFunc("/editproject/{id}", editProject).Methods("GET")
	route.HandleFunc("/submitedit", submitEdit).Methods("POST")
	route.HandleFunc("/deleteproject/{id}", deleteProject).Methods("GET")

	port := "5000"

	fmt.Print("Server sedang berjalan di port " + port + "\n")
	http.ListenAndServe("localhost:"+port, route)
}

// Home
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// fmt.Println(projects)
	dataProject := map[string]interface{}{
		"Projects": projects,
	}

	tmpt.Execute(w, dataProject)
}

// Contact
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

// Add Project
func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/add-project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

// Edit Project
func editProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/edit-project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

// Project struct
type Project struct {
	Title    string
	Content  string
	Duration string
}

// var projects = []
var projects = []Project{
	{
		Title:    "Judul",
		Content:  "Halo Dumbways",
		Duration: "1 bulan",
	},
}

// Project Form Submit
func submitProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	var title = r.PostForm.Get("title")
	var content = r.PostForm.Get("content")

	var newProject = Project{
		Title:   title,
		Content: content,
	}

	// fmt.Println(
	// 	"Title: "+r.PostForm.Get("title"),
	// 	"\nContent: "+r.PostForm.Get("content"),
	// 	"\nDate Start: "+r.PostForm.Get("datestart"),
	// 	"\nDate Start: "+r.PostForm.Get("datestart"),
	// 	"\nTechnologies: ",
	// 	"\n Node Js: "+r.PostForm.Get("nodejs"),
	// 	"\n React Js: "+r.PostForm.Get("reactjs"),
	// 	"\n Next Js: "+r.PostForm.Get("nextjs"),
	// 	"\n TypeScript: "+r.PostForm.Get("typescript"),
	// )

	// projects.push(newProject)
	projects = append(projects, newProject)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Submit Edit
func submitEdit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var title = r.PostForm.Get("title")
	var content = r.PostForm.Get("content")

	projects[id].Title = title
	projects[id].Content = content

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Project Detail
func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// Id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var projectInfo = Project{}

	for index, data := range projects {
		if index == id {
			projectInfo = Project{
				Title:    data.Title,
				Content:  data.Content,
				Duration: data.Duration,
			}
		}
	}

	dataDetail := map[string]interface{}{
		"Project": projectInfo,
	}

	tmpt.Execute(w, dataDetail)
}

// Delete Project
func deleteProject(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	fmt.Println(id)

	projects = append(projects[:id], projects[id+1:]...)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
