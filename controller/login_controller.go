package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/common"
	"user-service/config"
	"user-service/model"
	"user-service/repo"
	"user-service/service"
)

func Login(appCtx config.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetDatabaseConnection()
		var userLogin model.UserLogin
		if err := ctx.ShouldBind(&userLogin); err != nil {
			panic(err)
		}
		userRepo := repo.NewDbObj(db)
		loginService := service.NewLoginService(userRepo)

		token, err := loginService.Login(ctx.Request.Context(), &userLogin)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.NewAppResponseSimple(token))

	}
}
