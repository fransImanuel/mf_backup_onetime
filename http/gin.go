package http

import (
	"fmt"
	"mf_backup_onetime/dto"
	"mf_backup_onetime/http/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Gin       *gin.Engine
	AppConfig *dto.AppConfig
}

func InitServer() *Server {
	log.Info("InitServer() - starting...")
	log.Info("InitServer() - finished.")
	return &Server{}
}

func (s *Server) SetConfig(cfg *dto.AppConfig) {
	log.Info("Server - SetConfig() - starting...")
	log.Info("Server - SetConfig() - finished.")
	s.AppConfig = cfg
}

func (s *Server) InitGin() {
	log.Info("Server - InitGin() - starting...")
	appMiddleware := middleware.InitMiddleware()

	if !s.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	s.Gin = gin.Default()
	appCors := appMiddleware.CORS(s.AppConfig)
	s.Gin.Use(cors.New(appCors))
	s.Gin.Use(gin.Logger())
	s.Gin.Use(gin.Recovery())
	//s.Gin.Use(SetAuth(redis))
	log.Info("Server - InitGin() - finished.")
}

func (s *Server) Start() error {
	log.Info("Server - Start() - starting...")
	urlPath := fmt.Sprintf("%v:%v", s.AppConfig.Host, s.AppConfig.Port)
	var err error = nil
	if s.AppConfig.UseTLS && s.AppConfig.CertFilePath != "" && s.AppConfig.KeyFilePath != "" {
		err = s.Gin.RunTLS(urlPath, s.AppConfig.CertFilePath, s.AppConfig.KeyFilePath)
	} else {
		err = s.Gin.Run(urlPath)
	}

	log.Info("Server - Start() - finished.")
	return err
}
