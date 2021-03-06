package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//给配置chatId发送tg信息
func sendTgMsg2ConfigChatId(ctx *gin.Context) {
	var request model.SendTgMsg2ConfigChatIdRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("给配置chatId发送tg信息，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateResponseByErr(err))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("给配置chatId发送tg信息")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.SendTgMsg2ConfigChatId(ctx, request)))
}
