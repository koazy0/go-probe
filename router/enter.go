package router

import "github.com/gin-gonic/gin"

func SetUp() {
	r := gin.Default()
	r.Run(":56324")
}
