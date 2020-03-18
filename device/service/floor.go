package service

import (
	"open_iot/device/models"
	"time"

	common "github.com/tsbxmw/gin_common"
)

func (cps *ProjectService) FloorAdd(req *FloorAddRequest) *FloorAddResponse {
	res := FloorAddResponse{}
	floor := models.FloorModel{}
	if err := common.DB.Table(floor.TableName()).
		Where("name=? and location_id=? and building_id=?", req.Name, req.LocationId, req.BuildingId).
		First(&floor).Error; err != nil {
		if err.Error() == "record not found" {
			floor.Name = req.Name
			floor.LocationId = req.LocationId
			floor.BuildingId = req.BuildingId
			floor.Remark = req.Remark
			floor.BaseModel.ModifiedTime = time.Now()
			floor.BaseModel.CreationTime = time.Now()

			if err = common.DB.Table(floor.TableName()).Create(&floor).Error; err != nil {
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
		res.Message = "Floor already exists"
	}
	return &res
}

func (cps *ProjectService) FloorUpdate(id string, req *FloorUpdateRequest) *FloorUpdateResponse {
	res := FloorUpdateResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	floor := models.FloorModel{}
	if err := common.DB.Table(floor.TableName()).First(&floor, id).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
		panic(err)
	}
	if req.Name != "" {
		floor.Name = req.Name
	}
	if req.LocationId != 0 {
		floor.LocationId = req.LocationId
	}
	if req.BuildingId != 0 {
		floor.BuildingId = req.BuildingId
	}
	if req.Remark != "" {
		floor.Remark = req.Remark
	}
	if err := common.DB.Model(&floor).Save(&floor).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
		panic(err)
	}
	return &res
}

func (cps *ProjectService) FloorGet(req *FloorGetRequest) *FloorGetResponse {
	floor := models.FloorModel{}
	if err := common.DB.Table(floor.TableName()).
		First(&floor, req.FloorId).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	res := FloorGetResponse{
		Response: common.Response{
			Message: common.HTTP_MESSAGE_OK,
			Code:    common.HTTP_RESPONSE_OK,
		},
		Data: floor,
	}
	return &res
}

func (cps *ProjectService) FloorGetList(req *FloorGetListRequest) *FloorGetListResponse {
	res := FloorGetListResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
		Count:  0,
		Data:   make([]models.FloorModel, 0),
		Limmit: req.Limit,
		Offset: req.Offset,
	}
	db := common.DB.Table(models.FloorModel{}.TableName())

	if req.LocationId != 0 {
		db = db.Where("location_id=?", req.LocationId)
	}

	if req.BuildingId != 0 {
		db = db.Where("building_id=?", req.BuildingId)
	}

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
