package middleware

import (
	"fmt"
	"github.com/NuttayotSukkum/acn/acn-authentication/configs"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var hmacSampleSecret []byte

func ValidateTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := time.Now()
		cfg := configs.InitConfig()
		hmacSampleSecret = cfg.App.HmacSecret
		tokenHeader := c.Request().Header.Get("token")
		if tokenHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"httpStatus": strconv.Itoa(http.StatusUnauthorized),
				"time":       t.Format("2006-01-02 15:04:05"),
				"error":      "Token header is empty",
			})
		}
		tokenString := strings.TrimSpace(strings.Replace(tokenHeader, "Bearer", "", 1))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"httpStatus": strconv.Itoa(http.StatusForbidden),
				"time":       t.Format("2006-01-02 15:04:05"),
				"message":    "Invalid or expired token",
			})
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("userId", claims["userId"])
		return next(c)
	}
}
