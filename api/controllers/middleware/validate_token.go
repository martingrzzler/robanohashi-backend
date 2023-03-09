package middleware

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func ValidateFirebaseToken(auth *auth.Client, abortOnError bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		bearer := c.GetHeader("Authorization")

		if bearer == "" && abortOnError {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthenticated"})
			return
		} else if bearer == "" && !abortOnError {
			c.Next()
			return
		}

		idToken := strings.Split(bearer, " ")[1]

		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "unauthenticated"})
			return
		}

		uid := token.UID

		c.Set("uid", uid)
		c.Next()
	}
}
