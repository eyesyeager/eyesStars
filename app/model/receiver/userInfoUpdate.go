package receiver

import (
	"eyesStars/app/common/validate"
)

/**
 * @author eyesYeager
 * @date 2023/2/1 8:58
 */

type UserInfoUpdate struct {
	Name string `json:"name" binding:"required,max=10"` // "昵称"
}

func (userInfoUpdate UserInfoUpdate) GetMessages() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"name.required": "昵称不能为空",
		"name.max":      "昵称长度不能大于10",
	}
}
