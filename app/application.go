package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplicatioin() {
	mapurls()
	router.Run(":9092")
}
