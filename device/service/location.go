package service

import (
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/models"
	"time"
)

func (cps *ProjectService) LocationAdd(req *LocationAddRequest) *LocationAddResponse {
	res := LocationAddResponse{}
	location := models.LocationModel{}
	if err := common.DB.Table(location.TableName()).
		Where("name=? and location=?", req.Name, req.Location).First(&location).Error; err != nil {
		if err.Error() == "record not found" {
			location.Name = req.Name
			location.Location = req.Location
			location.Remark = req.Remark
			location.Remark = req.Remark
			location.BaseModel.ModifiedTime = time.Now()
			location.BaseModel.CreationTime = time.Now()

			if err = common.DB.Table(location.TableName()).Create(&location).Error; err != nil {
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
		res.Message = "Location already exists"
	}
	return &res
}

func (cps *ProjectService) LocationUpdate(id string, req *LocationUpdateRequest) *LocationUpdateResponse {
	res := LocationUpdateResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	location := models.LocationModel{}
	if err := common.DB.Table(location.TableName()).
		Where("name=? and location=?", req.Name, req.Location).First(&location).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
		panic(err)
	} else {
		if req.Name != "" {
			location.Name = req.Name
		}
		if req.Location != "" {
			location.Location = req.Location
		}
		if req.Remark != "" {
			location.Remark = req.Remark
		}
		location.BaseModel.ModifiedTime = time.Now()
		location.BaseModel.CreationTime = time.Now()

		if err = common.DB.Model(&location).Save(&location).Error; err != nil {
			common.LogrusLogger.Error(err)
			common.InitKey(cps.Ctx)
			cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
			panic(err)
		} else {
			res.Code = common.HTTP_RESPONSE_OK
			res.Message = common.HTTP_MESSAGE_OK
			res.Data = []string{}
		}
	}
	return &res
}

func (cps *ProjectService) LocationGet(req *LocationGetRequest) *LocationGetResponse {
	location := models.LocationModel{}
	common.LogrusLogger.Info(req.LocationId)
	if err := common.DB.Table(location.TableName()).
		First(&location, req.LocationId).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	res := LocationGetResponse{
		Response: common.Response{
			Message: common.HTTP_MESSAGE_OK,
			Code:    common.HTTP_RESPONSE_OK,
		},
		Data: location,
	}
	return &res
}

func (cps *ProjectService) LocationGetList(req *LocationGetListRequest) *locationGetListResponse {
	res := locationGetListResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
		Count:  0,
		Data:   make([]models.LocationModel, 0),
		Limmit: req.Limit,
		Offset: req.Offset,
	}
	db := common.DB.Table(models.LocationModel{}.TableName())

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
