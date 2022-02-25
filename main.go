package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

const staticDir = "/pic"

var (
	GitHash string
	BuildAt string
)

func main() {

	r := gin.New()

	//开发接口
	r.POST("/api", func(c *gin.Context) {
		arg := new(ReqJob)
		err := c.ShouldBind(arg)
		if handleError(c, err) {
			return
		}
		res, err := takeShot(arg)
		if handleError(c, err) {
			return
		}

		res.Url = fmt.Sprintf("http://%s/%s", c.Request.Host, res.Uri)

		c.JSON(200, res)
	})

	r.GET("/files/:file", func(context *gin.Context) {
		fileName := filepath.Join(staticDir, filepath.Base(context.Param("file")))

		if _, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
			context.JSON(404, ResJob{
				Code: 404,
				Msg:  "not found",
			})
			return
		}

		defer os.Remove(fileName)
		context.Header("Content-Type", "image/jpeg")
		context.File(fileName)
	})

	r.GET("version", func(context *gin.Context) {
		context.JSON(200, gin.H{"BuildAt": BuildAt, "GitHash": GitHash})
	})
	log.Fatal(r.Run(":8888"))
}
