package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/grade-tracker-backend/pkg/jwttoken"
	"github.com/linqcod/grade-tracker-backend/pkg/response"
	"github.com/linqcod/grade-tracker-backend/pkg/security/authorization"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwttoken.CheckTokenIsValid(c.Request)
		if err != nil {
			response.ResponseError(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}

func RoleCheckMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if role != authorization.ExtractRoleFromRequest(c.Request) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8084"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
