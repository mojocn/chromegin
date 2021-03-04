package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

const staticDir = "/pic"

var (
	BuildAt, GitHash string
)

func main() {

	r := gin.New()
	r.Use(static.Serve("/", static.LocalFile(staticDir, true)))

	//开发接口
	r.POST("/api", func(c *gin.Context) {
		c.Header("built_at", BuildAt)
		c.Header("git_hash", GitHash)
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

	log.Fatal(r.Run(":6666"))
}
