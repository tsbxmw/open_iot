package service

type (
	UserAddRequest struct {
		UserNickname string `json:"user_nickname" binding:"required"`
		Age          int    `json:"age" binding:"required"`
		Remark       string `json:"remark"`
	}

	UserGetRequest struct {
		UserNickname string `form:"user_nickname" binding:"required"`
	}
)

type (
	KengAddRequest struct {
		Name         string `json:"name" binding:"required"`
		DeviceId     int    `json:"device_id" binding:"required"`
		DeviceGpioId int    `json:"device_gpio_id" binding:"required"`
		RoomId       int    `json:"room_id" binding:"required"`
		Index        int    `json:"index" binding:"required"`
		Remark       string `json:"remark"`
	}

	KengUpdateRequest struct {
		Name         string `json:"name" binding:"required"`
		DeviceId     int    `json:"device_id" binding:"required"`
		DeviceGpioId int    `json:"device_gpio_id" binding:"required"`
		RoomId       int    `json:"room_id" binding:"required"`
		Index        int    `json:"index" binding:"required"`
		Remark       string `json:"remark"`
	}

	KengGetRequest struct {
		KengId int `form:"keng_id" binding:"required"`
	}

	KengGetListRequest struct {
		Limit      int `form:"limit" binding:"required"`
		Offset     int `form:"offset" binding:"required"`
		BuildingId int `form:"building_id"`
		RoomId     int `form:"room_id"`
		DeviceId   int `form:"device_id"`
	}

	KengGetFrontRequest struct {
		BuildingId int `form:"building_id" binding:"required"`
	}
)
