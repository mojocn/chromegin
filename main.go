package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	r := gin.New()

	//开发接口
	{
		//chromedp run screen shot
		r.GET("/python/ss", ChromedpShot)
		r.GET("/open/chromedp/screen/shot", ChromedpShot)
	}

	log.Fatal(r.Run(":6666"))
}

func init() {

	// open a file
	f, err := os.OpenFile("/data/golang_chrome_dp.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	defer f.Close()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}
