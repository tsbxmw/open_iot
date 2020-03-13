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

	KengRoom struct {
	}

	KengGetFrontResponse struct {
		common.Response
	}
)
