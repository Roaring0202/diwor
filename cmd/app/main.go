package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/Valghall/diwor/pkg/experimentator"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("web/template/*")
	r.StaticFile("/favicon.ico", "./web/static/favicon.ico")
	r.Static("/css", "web/static/css")
	r.Static("/image", "./web/static/image")

	r.GET("/", func(c *gin.Context) {
		var warning string

		if _, ok := c.GetQuery("reason"); ok {
			warning = "Алгоритмы не должны совпадать!"
		}

		c.HTML(http.StatusOK, "index.gohtml", warning)
	})

	r.GET("/experiment", func(c *gin.Context) {
		sampleA, sampleB, mode := c.Query("sample-a"), c.Query("sample-b"), c.Query("mode")

		if sampleA == sampleB {
			c.Redirect(http.StatusTemporaryRedirect, "/?reason=equal")
		}
		_, encryptionRes := experimentator.HoldExperiment(sampleA, sampleB, mode)
		c.HTML(http.StatusOK, "experiment.gohtml", string(encryptionRes))
	})

	log.Fatalln(r.Run())
}
