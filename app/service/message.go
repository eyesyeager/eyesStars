package service

import (
	"eyesStars/app/common"
	"eyesStars/app/constant/template"
	"eyesStars/app/model/entity"
	"eyesStars/app/model/receiver"
	"eyesStars/app/rpc/generate/userThrift"
	"eyesStars/app/rpc/rpc"
	"eyesStars/app/utils"
	"eyesStars/global"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/gin-gonic/gin"
)

/**
 * @author eyesYeager
 * @date 2024/7/25 22:42
 */

type messageService struct {
}

var MessageService = messageService{}

// SendMemoryEmail 发送回忆邮件
func (messageService messageService) SendMemoryEmail(c *gin.Context, data receiver.MemoryEmailSend) error {
	// 获取星星信息
	var stars []entity.Star
	if err := global.DB.Find(&stars, data.IdList).Error; err != nil {
		return err
	}

	// 组装uid
	var uidList []uint64
	for _, v := range stars {
		if !utils.Contains(uidList, uint64(v.Uid)) {
			uidList = append(uidList, uint64(v.Uid))
		}
	}

	// 获取用户信息
	var uid2InfoMap = map[uint64]*userThrift.UserInfoReturnee{}
	err, client, transport := rpc.User()
	if err != nil {
		return common.CustomError{}.SetErrorMsg("用户中心服务似乎宕机了")
	}
	defer func(transport thrift.TTransport) {
		closeErr := transport.Close()
		if closeErr != nil {
			global.Log.Error("UserService:thrift连接关闭异常！" + closeErr.Error())
		}
	}(transport)
	_, tUidList := utils.Uint64ToInt64Slice(uidList)
	userInfoList, err := client.GetBatchUserInfo(c, tUidList)
	for _, v := range userInfoList {
		uid2InfoMap[uint64(v.ID)] = v
	}

	// 开始发邮件
	for _, v := range stars {
		userInfo := uid2InfoMap[uint64(v.Uid)]
		emailInfo := template.GetMemoryEmailTemplate(v.Content, v.CreateTime, v.Name)
		if err = utils.EmailUtils.SendEmailWithHTML(emailInfo.Subject, []string{userInfo.Email}, emailInfo.Content); err != nil {
			return err
		}
	}

	return nil
}
