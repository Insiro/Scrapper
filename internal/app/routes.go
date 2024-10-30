package app

import "github.com/gin-gonic/gin"

type IRouter interface {
	gin.IRouter
	gin.IRoutes
}

func ApiRoute(r IRouter) *gin.RouterGroup {
	g := r.Group("/api")
	g.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "halo"})
	})
	return g
}
