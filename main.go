package main


import (
	"log"
	"net/http"
	"fmt"
	"os"
)

func myhandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hellofdsfsdafd!")
}
func main() {
	http.Handle("/public/", http.StripPrefix("/public/",
		http.FileServer(http.Dir("public"))))

	http.HandleFunc("/hello", myhandler)

	http.HandleFunc("/", myhandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Listening...")
	http.ListenAndServe(":" + port,nil)

	}