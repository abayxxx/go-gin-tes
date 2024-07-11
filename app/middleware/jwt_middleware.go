package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-gin/app/service"
	jwtServ "go-gin/app/service/jwt"
	"net/http"
	"strings"
)

// JwtMiddleware is a middleware to validate jwt token
func JwtMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("Authorization")
		if jwtToken == "" || jwtToken == "Bearer " {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		jwtToken = strings.TrimPrefix(jwtToken, "Bearer ")

		// Validate token with jwt service
		jwtService := jwtServ.NewJwtService()

		token, err := jwtService.ValidateToken(jwtToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		isExpired := jwtService.IsTokenExpired(jwtToken)
		if isExpired {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if claim["refresh_token"].(bool) {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
