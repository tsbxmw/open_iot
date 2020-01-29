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
	DeviceAddRequest struct {
		Name       string `json:"name" binding:"required"`
		DeviceType int    `json:"device_type" binding:"required"`
		MacAddress string `json:"mac_address" binding:"required"`
		LocationId int    `json:"location_id"`
		BuildingId int    `json:"building_id"`
		FloorId    int    `json:"floor_id"`
		RoomId     int    `json:"room_id"`
		Remark     string `json:"remark"`
	}

	DeviceUpdateRequest struct {
		Name       string `json:"name" binding:"required"`
		DeviceType int    `json:"device_type" binding:"required"`
		MacAddress string `json:"mac_address" binding:"required"`
		LocationId int    `json:"location_id"`
		BuildingId int    `json:"building_id"`
		FloorId    int    `json:"floor_id"`
		RoomId     int    `json:"room_id"`
		Remark     string `json:"remark"`
	}

	DeviceGetRequest struct {
		DeviceId int `form:"device_id" binding:"required"`
	}

	DeviceGetListRequest struct {
		Limit      int `form:"limit" binding:"required"`
		Offset     int `form:"offset" binding:"required"`
		LocationId int `form:"location_id"`
		BuildingId int `form:"building_id"`
		FloorId    int `form:"floor_id"`
		RoomId     int `form:"room_id"`
	}

	IPUpdateRequest struct {
		IpAddress  string `json:"ip_address" binding:"required"`
		MacAddress string `json:"mac_address" binding:"required"`
	}
)

type (
	LocationAddRequest struct {
		Name     string `json:"name" binding:"required"`
		Location string `json:"location" binding:"required"`
		Remark   string `json:"remark"`
	}

	LocationUpdateRequest struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Remark   string `json:"remark"`
	}

	LocationGetRequest struct {
		LocationId int `json:"location_id"`
	}

	LocationGetListRequest struct {
		Limit  int `form:"limit" binding:"required"`
		Offset int `form:"offset" binding:"required"`
	}
)

type (
	BuildingAddRequest struct {
		Name       string `json:"name" binding:"required"`
		LocationId int    `json:"location_id" binding:"required"`
		Remark     string `json:"remark"`
	}

	BuildingUpdateRequest struct {
		Name       string `json:"name" binding:"required"`
		LocationId int    `json:"location_id" binding:"required"`
		Remark     string `json:"remark"`
	}

	BuildingGetRequest struct {
		BuildingId int `form:"building_id" binding:"required"`
	}

	BuildingGetListRequest struct {
		Limit      int `form:"limit" binding:"required"`
		Offset     int `form:"offset" binding:"required"`
		LocationId int `form:"location_id" binding:"required"`
	}
)

type (
	FloorAddRequest struct {
		Name       string `json:"name" binding:"required"`
		LocationId int    `json:"location_id" binding:"required"`
		BuildingId int    `json:"building_id" bidding:"required"`
		Remark     string `json:"remark"`
	}

	FloorUpdateRequest struct {
		Name       string `json:"name" binding:"required"`
		LocationId int    `json:"location_id" binding:"required"`
		BuildingId int    `json:"building_id" bidding:"required"`
		Remark     string `json:"remark"`
	}
	FloorGetRequest struct {
		FloorId int `form:"floor_id"`
	}

	FloorGetListRequest struct {
		Limit      int `form:"limit" binding:"required"`
		Offset     int `form:"offset" binding:"required"`
		LocationId int `form:"location_id"`
		BuildingId int `form:"building_id"`
	}
)

type (
	RoomAddRequest struct {
		Name       string `json:"name" binding:"required"`
		LocationId int    `json:"location_id" binding:"required"`
		BuildingId int    `json:"building_id" bidding:"required"`
		FloorId    int    `json:"floor_id" binding:"required"`
		RoomType   int    `json:"room_type" binding:"required"`
		Remark     string `json:"remark"`
	}

	RoomUpdateRequest struct {
		Name       string `json:"name" binding:"required"`
		LocationId int    `json:"location_id" binding:"required"`
		BuildingId int    `json:"building_id" bidding:"required"`
		FloorId    int    `json:"floor_id" binding:"required"`
		RoomType   int    `json:"room_type" binding:"required"`
		Remark     string `json:"remark"`
	}

	RoomGetRequest struct {
		RoomId int `form:"room_id"`
	}

	RoomGetListRequest struct {
		Limit      int `form:"limit" binding:"required"`
		Offset     int `form:"offset" binding:"required"`
		LocationId int `form:"location_id"`
		BuildingId int `form:"building_id"`
		FloorId    int `form:"floor_id"`
		RoomType   int `form:"room_type"`
	}
)

type (
	RoomTypeAddRequest struct {
		Name   string `json:"name" binding:"required"`
		Remark string `json:"remark"`
	}

	RoomTypeUpdateRequest struct {
		Name   string `json:"name" binding:"required"`
		Remark string `json:"remark:`
	}

	RoomTypeGetRequest struct {
		RoomId int `form:"room_id" binding:"required"`
	}

	RoomTypeGetListRequest struct {
		Limit  int `form:"limit" binding:"required"`
		Offset int `form:"offset" binding:"required"`
	}
)
