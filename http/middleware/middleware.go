package middleware

import (
	"mf_backup_onetime/dto"

	"github.com/gin-contrib/cors"
)

type Middleware struct {
	MiddlewareConfig *cors.Config
}

func InitMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CORS(appCfg *dto.AppConfig) cors.Config {
	var httpCors cors.Config
	if appCfg.Debug {
		httpCors = cors.DefaultConfig()
		httpCors.AllowAllOrigins = true
		httpCors.AllowCredentials = true
		httpCors.AddAllowHeaders("authorization")
	} else {
		httpCors = cors.Config{
			AllowAllOrigins:        appCfg.Cors.AllowAllOrigins,
			AllowOrigins:           appCfg.Cors.AllowOrigins,
			AllowMethods:           appCfg.Cors.AllowMethods,
			AllowHeaders:           appCfg.Cors.AllowHeaders,
			ExposeHeaders:          appCfg.Cors.ExposeHeaders,
			AllowCredentials:       appCfg.Cors.AllowCredentials,
			AllowWildcard:          appCfg.Cors.AllowWildcard,
			AllowBrowserExtensions: appCfg.Cors.AllowBrowserExtensions,
			AllowWebSockets:        appCfg.Cors.AllowWebSockets,
			AllowFiles:             appCfg.Cors.AllowFiles,
			MaxAge:                 appCfg.Cors.MaxAge,
		}
	}

	return httpCors
}
