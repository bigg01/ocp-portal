package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"github.com/golang/glog"

	"github.com/bigg01/ocp-portal/server/routes"
)

func main() {

	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	// Ping test
	server.GET("/ping", ping)

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})



	server.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})


	server.POST("/post", func(c *gin.Context) {

			id := c.Query("id")
			page := c.DefaultQuery("page", "0")
			name := c.PostForm("name")
			message := c.PostForm("message")
		    glog.Info("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		//c.String(200, "post")
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
			//fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		})

	// Listen and Server in 0.0.0.0:8080
	//server.Run(":8080")

	s := &http.Server{
		Addr:           ":8088",
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	fmt.Println("end server")
	glog.Flush()

}
