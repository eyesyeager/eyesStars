package receiver

import (
	"eyesStars/app/common/validate"
)

/**
 * @author eyesYeager
 * @date 2023/2/1 21:25
 */

type ContextUpdate struct {
	Name    string `json:"name" binding:"required"`    // 配置名
	Content string `json:"content" binding:"required"` // 配置内容
}

func (contextUpdate ContextUpdate) GetMessages() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"name.required":    "配置名不能为空",
		"content.required": "配置内容不能为空",
	}
}
