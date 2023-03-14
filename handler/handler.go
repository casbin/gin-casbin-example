package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dataset1(c *gin.Context) {
	sub, _, _ := c.Request.BasicAuth()
	c.String(http.StatusOK, "%s %s %s", sub, c.Request.Method, "dataset1/"+c.Param("resource"))
}

func CreateResource1(c *gin.Context) {
	sub, _, _ := c.Request.BasicAuth()
	c.String(http.StatusOK, "%s POST resource1", sub)
}

func Dataset2Resource1(c *gin.Context) {
	sub, _, _ := c.Request.BasicAuth()
	c.String(http.StatusOK, "%s %s dataset2/resource1", sub, c.Request.Method)
}
