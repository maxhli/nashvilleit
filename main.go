package main


import (
	"log"
	"net/http"
	"fmt"
)

func myhandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hellofdsfsdafd!")
}
func main() {
	http.Handle("/public/", http.StripPrefix("/public/",
		http.FileServer(http.Dir("public"))))

	http.HandleFunc("/hello", myhandler)

	// http.Handle("dfdf", fs)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}
