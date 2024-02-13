package rest

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
)

const (
	Apikey string = "3caf85347f7e49e481d110120241401"
	City   string = "Batumi"
)

func NewServer() {
	g := gin.Default()

	g.GET("/weather", func(ctx *gin.Context) {
		url := fmt.Sprintf("%weather.json?key=%s&q=%s", City, Apikey)
		ctx.String(http.StatusOK, "city")
		resp, err := http.Get(url)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"data": url})
			return
		}
		defer resp.Body.Close()

		ctx.JSON(resp.StatusCode, resp.Body)
	})

	g.Run(":8080")

}
