package apihandlers

import (
	"github.com/kataras/iris"
	"github.com/wonderivan/logger"
	"lucky-draw/internal/lucky-drawn/reward"
	"lucky-draw/internal/lucky-drawn/wechat"
	"net/url"
)

func StartLuckyDraw(ctx iris.Context) {
	logger.Info("*************************************** Start Lucky Draw ***************************************")
	users := reward.GetAllUsers()
	user := reward.GetAwardUser(users)
	logger.Info("Congratulations! The award user is %+v", user)
	ctx.ResponseWriter().WriteString(`{"Id":"` + user.Id + `","NickName":"` + user.NickName + `"}`)
	return
}

func GetHeadImg(ctx iris.Context) {
	var fakeid = ""
	params, _ := url.ParseQuery(ctx.Request().URL.RawQuery)
	if values, ok := params["fakeid"]; ok {
		fakeid = values[0]
	}

	ctx.ResponseWriter().Header().Set("Content-Type", "image/jpg")
	ctx.ResponseWriter().Write(wechat.GetHeadImg(fakeid))
	return
}
