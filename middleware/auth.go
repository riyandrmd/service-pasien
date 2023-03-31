package middleware

import (
	"administrasi/apperror"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var AuthSecret = "rumah-sakit"
var RefreshAuthSecret = "rumah-sakit-refresh"
var From = "http://localhost:8082"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			apperror.Response(c, apperror.New(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)))
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, apperror.New(http.StatusForbidden, http.StatusText(http.StatusForbidden))
			}

			return []byte(AuthSecret), nil
		})
		if err != nil {
			apperror.Response(c, apperror.New(http.StatusForbidden, err.Error()))
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			apperror.Response(c, apperror.New(http.StatusForbidden, http.StatusText(http.StatusForbidden)))
			return
		}
		exp, _ := claims["exp"].(float64)
		if time.Now().After(time.Unix(int64(exp), 0)) {
			apperror.Response(c, apperror.New(http.StatusForbidden, http.StatusText(http.StatusForbidden)))
			return
		}
		uid, _ := claims["uid"].(float64)
		c.Set("id", int(uid))
		c.Next()
	}
}
