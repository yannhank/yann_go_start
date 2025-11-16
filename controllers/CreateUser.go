package controllers

import (
	"net/http"
	"sc_police_api/global"
	"sc_police_api/models"
	"sc_police_api/utils"

	"github.com/gin-gonic/gin"
)

// 应用启动时调用一次，创建表
func MigrateTables() error {
	return global.Db.AutoMigrate(&models.ScUserPolice{})
}

func CreateUser(ctx *gin.Context) {
	var user models.ScUserPolice
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 10002, "msg": "参数错误"})
		return
	}

	user.Code = utils.GetRandomString(12)
	user.Password, _ = utils.HashPassword(user.Password)
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 10003, "msg": "创建用户失败，请检查后重试"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 10001, "msg": "创建用户成功"})
}
