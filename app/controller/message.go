package controller

import (
	"eyesStars/app/common"
	"eyesStars/app/common/result"
	"eyesStars/app/common/validate"
	"eyesStars/app/constant"
	"eyesStars/app/model/receiver"
	"eyesStars/app/service"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2024/7/24 22:29
 */

// SendMemoryEmail
// @Summary			发送回忆邮件
// @Description		权限：Admin
// @Tags			message
// @Accept			json
// @Produce			json
// @Security		sToken,lToken
// @Param			sToken		header		string		false			"短token"
// @Param			lToken		header		string 		false			"长token"
// @Param			receiver	body    	receiver.MemoryEmailSend	true		"请求参数"
// @Success			200			{object}	result.Response
// @Failure			400			{object}	result.Response
// @Router			/message/sendMemoryEmail [post]
func SendMemoryEmail(c *gin.Context) {
	// 权限校验
	if !common.CheckAuth(c, constant.Roles.Admin) {
		result.FailByCustom(c, result.Results.AuthInsufficient)
		return
	}

	// 参数校验
	var data receiver.MemoryEmailSend
	if err := c.ShouldBindJSON(&data); err != nil {
		result.FailAttachedMsg(c, validate.GetErrorMsg(data, err))
		return
	}

	// 执行业务
	if err := service.MessageService.SendMemoryEmail(c, data); err != nil {
		result.FailAttachedMsg(c, err.Error())
	} else {
		result.SuccessDefault(c, nil)
	}
}
