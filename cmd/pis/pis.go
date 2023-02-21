package main

import (
	"github.com/serverkona1/prepaid-card-inquiry-service/api/routes"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func init() {
	debug.SetTraceback("all")
}

func main() {
	//routes.InitRoutes(gin.Default()).Run()
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "pong"})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	routers := routes.NewRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      routers,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.SetKeepAlivesEnabled(true)
	log.Println("pis_dev Server start : ", server.Addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Could not listen on : %v\n", err)
	}
}
