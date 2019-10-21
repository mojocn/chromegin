package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.New()

	//开发接口
	{
		//chromedp run screen shot
		r.GET("/python/ss", ChromedpShot)
	}

	log.Fatal(r.Run(":6666"))
}
