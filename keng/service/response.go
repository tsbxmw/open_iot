package service

import (
	"open_iot/keng/models"

	common "github.com/tsbxmw/gin_common"
)

type (
	UserAddResponse struct {
		common.Response
	}

	UserGetResponse struct {
		common.Response
		Data models.UserModel `json:"data"`
	}
)

type (
	KengAddResponse struct {
		common.Response
	}

	KengUpdateResponse struct {
		common.Response
	}

	KengGetResponse struct {
		common.Response
		Data models.KengModel `json:"data"`
	}

	KengGetListResponse struct {
		common.Response
		Data   []models.KengModel `json:"data"`
		Limit  int                `json:"limit"`
		Offset int                `json:"offset"`
		Count  int                `json:"count"`
	}

	AllInfoLocations struct {
		LocationBuilding []LocationBuilding `json:"locations"`
	}
	LocationBuilding struct {
		LocationId    int             `json:"location_id"`
		LocationName  string          `json:"location_name"`
		BuildingFloor []BuildingFloor `json:"buildings`
	}

	BuildingFloor struct {
		BuildingId   int         `json:"building_id"`
		BuildingName string      `json:"building_name"`
		FloorRoom    []FloorRoom `json:"floors"`
	}
	FloorRoom struct {
		FloorId    int          `json:"floor_id"`
		FloorName  string       `json:"floor_name"`
		RoomDevice []RoomDevice `json:"rooms`
	}
	RoomDevice struct {
		RoomId     int          `json:"room_id"`
		RoomName   string       `json:"room_name"`
		DeviceGpio []DeviceGpio `json:"device"`
		KengInfo   []KengInfo   `json:"keng_info"`
	}

	DeviceGpio struct {
		DeviceId   int    `json:"device_id"`
		DeviceName string `json:"device_name"`
		GpioInfo   []Gpio `json:"gpio_info"`
	}

	Gpio struct {
		GpioId     int     `json:"gpio_id"`
		GpioNumber int     `json:"gpio_number"`
		GpioStatus int     `json:"gpio_status"`
		GpioTime   float32 `json:"gpio_time"`
	}

	KengInfo struct {
		KengId     int     `json:"keng_id"`
		KengName   string  `json:"keng_name"`
		KengIndex  int     `json:"keng_index"`
		KengStatus int     `json:"keng_status"`
		KengTime   float32 `json:"keng_time"`
	}

	KengGetFrontResponse struct {
		common.Response
		Data []LocationBuilding `json:"data"`
	}
)
