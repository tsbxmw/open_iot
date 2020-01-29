package service

import (
	"open_iot/device/models"

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
	DeviceAddResponse struct {
		common.Response
	}

	DeviceUpdateResponse struct {
		common.Response
	}

	DeviceGetResponse struct {
		common.Response
		Data models.DeviceModel `json:"data"`
	}

	DeviceGetListResponse struct {
		common.Response
		Limmit int                  `json:"limit"`
		Offset int                  `json:"offset"`
		Count  int                  `json:"count"`
		Data   []models.DeviceModel `json:"data"`
	}

	IPUpdateResponse struct {
		common.Response
	}
)

type (
	LocationAddResponse struct {
		common.Response
	}

	LocationUpdateResponse struct {
		common.Response
	}

	LocationGetResponse struct {
		common.Response
		Data models.LocationModel `json:"data"`
	}

	locationGetListResponse struct {
		common.Response
		Limmit int                    `json:"limit"`
		Offset int                    `json:"offset"`
		Count  int                    `json:"count"`
		Data   []models.LocationModel `json:"data"`
	}
)

type (
	BuildingAddResponse struct {
		common.Response
	}

	BuildingUpdateResponse struct {
		common.Response
	}

	BuildingGetResponse struct {
		common.Response
		Data models.BuildingModel `json:"data"`
	}

	BuildingGetListResponse struct {
		common.Response
		Limmit int                    `json:"limit"`
		Offset int                    `json:"offset"`
		Count  int                    `json:"count"`
		Data   []models.BuildingModel `json:"data"`
	}
)

type (
	FloorAddResponse struct {
		common.Response
	}

	FloorUpdateResponse struct {
		common.Response
	}

	FloorGetResponse struct {
		common.Response
		Data models.FloorModel `json:"data"`
	}

	FloorGetListResponse struct {
		common.Response
		Limmit int                 `json:"limit"`
		Offset int                 `json:"offset"`
		Count  int                 `json:"count"`
		Data   []models.FloorModel `json:"data"`
	}
)

type (
	RoomAddResponse struct {
		common.Response
	}

	RoomUpdateResponse struct {
		common.Response
	}

	RoomGetResponse struct {
		common.Response
		Data models.RoomModel `json:"data"`
	}

	RoomGetListResponse struct {
		common.Response
		Limmit int                `json:"limit"`
		Offset int                `json:"offset"`
		Count  int                `json:"count"`
		Data   []models.RoomModel `json:"data"`
	}
)
