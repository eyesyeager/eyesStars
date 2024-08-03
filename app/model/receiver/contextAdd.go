package receiver

import (
	"eyesStars/app/common/validate"
)

/**
 * @author eyesYeager
 * @date 2023/2/1 20:37
 */

type ContextAdd struct {
	Name    string `json:"name" binding:"required"`    // 配置名
	Content string `json:"content" binding:"required"` // 配置内容
	Remarks string `json:"remarks"`                    // 配置说明
}

func (contextAdd ContextAdd) GetMessages() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"name.required":    "配置名不能为空",
		"content.required": "配置内容不能为空",
	}
}
