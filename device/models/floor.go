package models

import (
	common "github.com/tsbxmw/gin_common"
)

type (
	FloorModel struct {
		common.BaseModel
		Name       string `json:"name"`
		BuildingId int    `json:"building_id"`
		LocationId int    `json:"location_id"`
		Remark     string `json:"remark"`
	}
)

func (FloorModel) TableName() string {
	return "floor"
}
