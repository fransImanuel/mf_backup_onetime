package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	//"mf_backup_onetime/module/user"
	"mf_backup_onetime/util"
)

//type Auths struct {
//	UserService user.Service
//}

//func Auth(c *gin.Context) {
//	resp := dto.ResponseBaseDto{
//		Code:    http.StatusUnauthorized,
//		Message: "error",
//	}
//
//	token := c.Request.Header.Get("Authorization")
//
//
//	redis := c.MustGet("redisClient").(*util.Redis)
//
//	if redis == nil {
//		resp.Message = "error on get redis client"
//		c.JSON(http.StatusUnauthorized, resp)
//		c.Abort()
//		return
//	}
//
//	userJSON, err := redis.Client.Get("user:" + token).Result()
//	if err != nil || userJSON == "" {
//		resp.Message = "token invalid"
//		c.JSON(http.StatusUnauthorized, resp)
//		c.Abort()
//		return
//	}
//}

func Auth(c *gin.Context) {

	JwtSecret := os.Getenv("JWT_SECRET")
	token := c.Request.Header.Get("Authorization")
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JwtSecret), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": http.StatusUnauthorized,
			"message":     err.Error(),
			"data":        nil,
		})
		c.Abort()
	}
}

func SetAuth(redis *util.Redis) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redisClient", redis)
		c.Next()
	}
}
