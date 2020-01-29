package service

import (
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/models"
	"time"
)

func (cps *ProjectService) RoomAdd(req *RoomAddRequest) *RoomAddResponse {
	res := RoomAddResponse{
		common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	room := models.RoomModel{}
	if err := common.DB.Table(room.TableName()).
		Where("name=? and location_id=? and building_id=? and floor_id=?",
			req.Name, req.LocationId, req.BuildingId, req.FloorId).
		First(&room).Error; err != nil {
		if err.Error() == "record not found" {
			room.Name = req.Name
			room.LocationId = req.LocationId
			room.BuildingId = req.BuildingId
			room.FloorId = req.FloorId
			room.RoomType = req.RoomType
			room.Remark = req.Remark
			room.BaseModel.ModifiedTime = time.Now()
			room.BaseModel.CreationTime = time.Now()

			if err = common.DB.Table(room.TableName()).Create(&room).Error; err != nil {
				cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
				panic(err)
			} else {
				res.Code = common.HTTP_RESPONSE_OK
				res.Message = common.HTTP_MESSAGE_OK
				res.Data = []string{}
			}
		} else {
			cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
			panic(err)
		}
	} else {
		res.Code = 0
		res.Message = "Room already exists"
	}
	return &res
}

func (cps *ProjectService) RoomUpdate(id string, req *RoomUpdateRequest) *RoomUpdateResponse {
	res := RoomUpdateResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	room := models.RoomModel{}
	if err := common.DB.Table(room.TableName()).First(&room, id).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
		panic(err)
	}
	if req.Name != "" {
		room.Name = req.Name
	}
	if req.LocationId != 0 {
		room.LocationId = req.LocationId
	}
	if req.BuildingId != 0 {
		room.BuildingId = req.BuildingId
	}
	if req.Remark != "" {
		room.Remark = req.Remark
	}
	if err := common.DB.Model(&room).Save(&room).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
		panic(err)
	}
	return &res
}

func (cps *ProjectService) RoomGet(req *RoomGetRequest) *RoomGetResponse {
	room := models.RoomModel{}

	if err := common.DB.Table(room.TableName()).
		First(&room, req.RoomId).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	res := RoomGetResponse{
		Response: common.Response{
			Message: common.HTTP_MESSAGE_OK,
			Code:    common.HTTP_RESPONSE_OK,
		},
		Data: room,
	}
	return &res
}

func (cps *ProjectService) RoomGetList(req *RoomGetListRequest) *RoomGetListResponse {
	res := RoomGetListResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
		Count:  0,
		Data:   make([]models.RoomModel, 0),
		Limmit: req.Limit,
		Offset: req.Offset,
	}
	db := common.DB.Table(models.RoomModel{}.TableName())

	if err := db.Offset((req.Offset - 1) * req.Limit).Limit(req.Limit).Find(&res.Data).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(cps.Ctx)
			cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}
	if err := db.Count(&res.Count).Error; err != nil {
		if err.Error() != "record not found" {
			common.LogrusLogger.Error(err)
			common.InitKey(cps.Ctx)
			cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
			panic(err)
		}
	}
	return &res
}
