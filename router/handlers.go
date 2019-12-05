package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"korzhishe.ru/grabber/news"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}

func collectHandler(c *gin.Context) {
	cat := c.Param("catigory")

	news.Collect(cat)
	c.String(http.StatusOK, "Search is initialized")
}

func resultHandler(c *gin.Context) {
	cat := c.Param("catigory")
	topics := news.Result(cat)
	c.JSON(http.StatusOK, topics)
}
