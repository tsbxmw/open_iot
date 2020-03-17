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

	LocationBuilding struct {
		LocationId   int             `json:"location_id"`
		LocationName string          `json:"location_name"`
		Buildings    []BuildingFloor `json:"buildings`
	}

	BuildingFloor struct {
		BuildingId   int         `json:"building_id"`
		BuildingName string      `json:"building_name"`
		Floors       []FloorRoom `json:"floors"`
	}
	FloorRoom struct {
		FloorId   int        `json:"floor_id"`
		FloorName string     `json:"floor_name"`
		RoomKeng  []RoomKeng `json:"rooms`
	}
	RoomKeng struct {
		RoomId   int        `json:"room_id"`
		RoomName string     `json:"room_name"`
		Keng     []KengInfo `json:"kengs"`
	}

	KengInfo struct {
		KengId     int    `json:"keng_id"`
		KengName   string `json:"keng_name"`
		KengIndex  int    `json:"keng_index"`
		KengStatus int    `json:"keng_status"`
	}

	KengGetFrontResponse struct {
		common.Response
		Data LocationBuilding `json:"data"`
	}
)
