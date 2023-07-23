package main

import (
	"github.com/gin-gonic/gin"
)

func healthzServer(c *gin.Context) {
	c.String(200, "%s", "ok")
}
