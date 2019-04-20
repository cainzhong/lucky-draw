package apihandlers

import (
	"github.com/kataras/iris"
	"github.com/wonderivan/logger"
	"lucky-draw/internal/lucky-drawn/reward"
)

func StartLuckyDraw(ctx iris.Context){
	logger.Info("*************************************** Start Lucky Draw ***************************************")
	users := reward.GetAllUsers()
	user := reward.GetAwardUser(users)
	logger.Info("Congratulations! The award user is %+v",user)
	ctx.View("congratulation.html")
	return
}