package middleware

import (
	"net/http"
	"strings"

	"github.com/Caknoooo/Golang-BLOG/services"
	"github.com/Caknoooo/Golang-BLOG/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Ditemukan", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !token.Valid {
			response := utils.BuildResponseFailed("Gagal Memproses Request", "Akses Ditolak", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed("Gagal Memproses Request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Set("token", authHeader) // -> di set dlu tokennya, bisa manggil mu
		ctx.Set("userID", userID)
		ctx.Next()
	}
}