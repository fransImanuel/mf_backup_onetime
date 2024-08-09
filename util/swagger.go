package util

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GenerateSwagger(g *gin.Engine) string {
	log.Info("Swagger - GenerateSwagger() - starting...")

	log.Info("Swagger - GenerateSwagger() - finish...")
	return "Success"
}
