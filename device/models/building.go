package models

import (
	common "github.com/tsbxmw/gin_common"
)

type (
	BuildingModel struct {
		common.BaseModel
		Name       string `json:"name"`
		LocationId int    `json:"location_id"`
		Remark     string `json:"remark"`
	}
)

func (BuildingModel) TableName() string {
	return "building"
}
