package main

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/gin-casbin-example/handler"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func setupRouter() (router *gin.Engine) {
	// load the casbin model and policy from files, database is also supported.
	e, err := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	if err != nil {
		return
	}

	// define your router, and use the Casbin authz middleware.
	// the access that is denied by authz will return HTTP 403 error.
	router = gin.New()
	router.Use(authz.NewAuthorizer(e))

	dataset1 := router.Group("/dataset1")
	dataset1.Any("/:resource", handler.Dataset1)
	dataset1.POST("/resource1", handler.CreateResource1)

	router.Any("/dataset2/resource1", handler.Dataset2Resource1)

	return
}

func main() {
	router := setupRouter()
	log.Fatal(router.Run(":8080"))
}
