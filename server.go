package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/rep2recall/rep2recall-go/docs"
)

/*
@title Rep2recall API
@version 1.0
@description SRS quizzing API

@contact.name Pacharapol Withayasakpunt
@contact.url https://www.polv.cc
@contact.email polv@polv.cc

@license.name MIT
@license.url https://mit-license.org/
*/
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()

	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)
	if runtime.GOOS == "linux" {
		if data, err := ioutil.ReadFile("/proc/version"); err == nil && strings.Contains(string(data), "microsoft") {
			if p, err := strconv.Atoi(port); err == nil {
				swaggerURL = fmt.Sprintf("http://localhost:%d/swagger/doc.json", p+1)
			}
		}
	}

	url := ginSwagger.URL(swaggerURL)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Printf("Server running at http://localhost:%s", port)
	r.Run(":" + port)
}
