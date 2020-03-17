package models

import (
	common "github.com/tsbxmw/gin_common"
)

type (
	DeviceModel struct {
		common.BaseModel
		Name       string `json:"name"`
		DeviceType int    `json:"device_type"`
		IpAddress  string `json:"ip_address"`
		MacAddress string `json:"mac_address"`
		LocationId int    `json:"location_id"`
		BuildingId int    `json:"building_id"`
		FloorId    int    `json:"floor_id"`
		RoomId     int    `json:"room_id"`
		Remark     string `json:"remark"`
	}

	DeviceTypeModel struct {
		common.BaseModel
		Name string `json:"name"`
	}

	DeviceGpioModel struct {
		common.BaseModel
		DeviceId   int `json:"device_id"`
		GpioNumber int `json:"gpio_number"`
		GpioStatus int `json:"gpio_status"`
	}

	DeviceGpioRecordModel struct {
		common.BaseModel
		DeviceId   int `json:"device_id"`
		GpioNumber int `json:"gpio_number"`
		GpioStatus int `json:"gpio_status"`
	}
)

func (DeviceModel) TableName() string {
	return "device"
}

func (DeviceTypeModel) TableName() string {
	return "device_type"
}

func (DeviceGpioModel) TableName() string {
	return "device_gpio"
}

func (DeviceGpioRecordModel) TableName() string {
	return "device_gpio_record"
}
