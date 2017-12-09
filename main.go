package main


import (
	"log"
	"net/http"
	"fmt"
	"os"
	"html/template"
	"io/ioutil"
)

func myhandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hehhehehehheheheheheh!")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := "templates/index.tmpl.html"
	t, _ := template.ParseFiles(
		title, "templates/header.tmpl.html", "templates/nav.tmpl.html")
	fmt.Println("t is: ", t)
	t.ExecuteTemplate(w, title, nil)
	}

func RootHandler(res http.ResponseWriter, req *http.Request) {
	file, _ := ioutil.ReadFile("templates/index.tmpl.html")
	res.Write(file)
}

func main() {

	//http.HandleFunc("/", viewHandler)
	http.HandleFunc("/", RootHandler)

	http.Handle("/public/", http.StripPrefix("/public/",
		http.FileServer(http.Dir("public"))))

	http.HandleFunc("/hello", myhandler)


	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Listening...")
	http.ListenAndServe(":" + port,nil)

	}