package main


import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func myhandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hellofdsfsdafd!")
}
func main() {
	//http.Handle("/public/", http.StripPrefix("/public/",
	//	http.FileServer(http.Dir("public"))))
	//
	//http.HandleFunc("/hello", myhandler)
	//
	//http.HandleFunc("/", myhandler)
	//
	//// http.Handle("dfdf", fs)
	//
	//log.Println("Listening...")
	//http.ListenAndServe("",nil)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.Static("/public", "public")



	router.GET("/onlinetraces", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	router.GET("/traces", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	router.GET("/fileupload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fileupload.tmpl.html", nil)
	})

	router.GET("/online", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	//router.GET("/repeat", repeatHandler)
	// the above one causes problem. Max Li


	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl.html", nil)

	})

	router.Run(":" + port)
	}