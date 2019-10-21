package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		//logrus.WithError(err).Error("gin context http handler error")
		c.JSON(200, gin.H{"msg": err.Error()})
		return true
	}
	return false
}
func md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
func ChromedpShot(c *gin.Context) {
	var err error
	u := c.Query("u")
	//url decode 参数
	u, err = url.QueryUnescape(u)
	if handleError(c, err) {
		return
	}
	if !strings.HasPrefix(u, "http") {
		c.JSON(200, gin.H{"msg": u + " 地址无效"})
		return
	}

	timeString := c.Query("c")
	timeString, err = url.QueryUnescape(timeString)
	if handleError(c, err) {
		return
	}
	t, err := time.Parse("2006-01-02 15:04:05", timeString)
	if handleError(c, err) {
		return
	}
	fileName := fmt.Sprintf("%s_%s.png", t.Format("060102T150405"), md5Encode(u))
	//imagePath := path.Join(os.TempDir(), fileName)
	imagePath := path.Join("/data", fileName)
	if _, err := os.Stat(imagePath); err == nil {
		//如果图片存在就直接gin response 图片
		c.File(imagePath)
		return
	}

	if err := runChromedp(u, imagePath); handleError(c, err) {
		return
	}

	c.File(imagePath)

}

func runChromedp(targetUrl, imagePath string) error {
	// create context
	// timeout 90 秒
	timeContext, cancelFunc := context.WithTimeout(context.Background(), time.Second*90)
	defer cancelFunc()

	ctx, cancel := chromedp.NewContext(timeContext)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(targetUrl, 50, &buf)); err != nil {
		return err
	}
	log.Println(imagePath)
	return ioutil.WriteFile(imagePath, buf, 0644)
}

func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Emulate(device.IPad),
		chromedp.EmulateViewport(1024, 2048, chromedp.EmulateScale(2)),
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}
			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
