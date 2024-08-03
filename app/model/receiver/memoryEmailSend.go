package receiver

import "eyesStars/app/common/validate"

/**
 * @author eyesYeager
 * @date 2024/7/25 22:46
 */

type MemoryEmailSend struct {
	IdList []uint32 `json:"idList" binding:"required"` // 星星id列表
}

func (memoryEmailSend MemoryEmailSend) GetMessages() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"idList.required": "星星id列表不能为空",
	}
}
