package models

import (
	common "github.com/tsbxmw/gin_common"
)

type (
	LocationModel struct {
		common.BaseModel
		Name     string `json:"name"`
		Location string `json:"location"`
		Remark   string `json:"remark"`
	}
)

func (LocationModel) TableName() string {
	return "location"
}
