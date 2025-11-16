package controllers

import (
	"net/http"
	"sc_police_api/utils"

	"github.com/gin-gonic/gin"
)

func CheckTokenb(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "缺少token"})
		return
	}
	username, err := utils.ParseJWT(token)
	if err != nil || username == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "验证成功", "username": username})
}
