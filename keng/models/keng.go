package models

import (
	common "github.com/tsbxmw/gin_common"
)

type (
	KengModel struct {
		common.BaseModel
		Name         string `json:"name"`
		DeviceId     int    `json:"device_id"`
		RoomId       int    `json:"room_id"`
		Index        int    `json:"index"`
		Remark       string `json:"remark"`
		DeviceGpioId int    `json:"device_gpio_id"`
		Status       int    `json:"status"`
	}
)

func (KengModel) TableName() string {
	return "keng"
}
