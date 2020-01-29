package models

import (
	common "github.com/tsbxmw/gin_common"
)

type (
	RoomModel struct {
		common.BaseModel
		Name       string `json:"name"`
		LocationId int    `json:"location_id"`
		BuildingId int    `json:"building_id"`
		FloorId    int    `json:"floor_id"`
		RoomType   int    `json:"room_type"`
		Remark     string `json:"remark"`
	}

	RoomTypeModel struct {
		common.BaseModel
		Name   string `json:"name"`
		Remark string `json:"remark"`
	}
)

func (RoomModel) TableName() string {
	return "room"
}

func (RoomTypeModel) TableName() string {
	return "room_type"
}
