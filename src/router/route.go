package router

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func initStatic(r *gin.Engine) {
	// r.StaticFS("", gin.Dir("./static/", true))
	r.Static("/web", "./static")

}
func Init(r *gin.Engine) {
	initStatic(r)
	r.GET("/ping", func(ctx *gin.Context) {
		log.Println(ctx.Request.URL.Path)
		ctx.JSON(http.StatusOK, "ok")
	})
	r.POST("/service/upload", func(ctx *gin.Context) {
		var upload struct {
			Data string
		}
		if ctx.Bind(&upload) == nil {
			re, err := regexp.Compile(`data:image/jpeg;base64,(.*)`)
			if err != nil {
				ctx.Error(err)
				return
			}
			s := re.FindStringSubmatch(upload.Data)
			b, _ := base64.RawStdEncoding.DecodeString(s[1])
			i := time.Now().UnixMicro()
			var savePath string
			savePath = "./image/image-" + strconv.Itoa(int(i)) + ".jpeg"
			fmt.Printf("savePath: %v\n", savePath)
			ioutil.WriteFile(savePath, b, 0644)
		}
	})
}
