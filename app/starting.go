package app

import (
	"github.com/gin-gonic/gin"
)



func StartApplication() {
	var router = gin.Default()

	router.Run(":3000")
}
