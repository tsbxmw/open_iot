package service

import (
	"open_iot/keng/models"
	"time"

	common "github.com/tsbxmw/gin_common"
)

func (cps *ProjectService) KengAdd(req *KengAddRequest) *KengAddResponse {
	res := KengAddResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	count := 0
	common.LogrusLogger.Info(req.Index, req.RoomId)
	if err := common.DB.Table(models.KengModel{}.TableName()).
		Where("device_id=? and `index`=? and room_id=?", req.DeviceId, req.Index, req.RoomId).Count(&count).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	if count == 0 {
		keng := models.KengModel{}

		keng.Remark = req.Remark
		keng.Name = req.Name
		keng.RoomId = req.RoomId
		keng.DeviceId = req.DeviceId
		keng.Index = req.Index

		keng.BaseModel.ModifiedTime = time.Now()
		keng.BaseModel.CreationTime = time.Now()

		if err := common.DB.Table(keng.TableName()).Create(&keng).Error; err != nil {
			cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
			panic(err)
		} else {
			res.Code = common.HTTP_RESPONSE_OK
			res.Message = common.HTTP_MESSAGE_OK
			res.Data = []string{}
		}
	} else {
		res.Code = 0
		res.Message = "Keng already exists"
	}
	return &res
}

func (cps *ProjectService) KengUpdate(id string, req *KengUpdateRequest) *KengUpdateResponse {
	res := KengUpdateResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	keng := models.KengModel{}
	if err := common.DB.Table(keng.TableName()).First(&keng, id).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	keng.Name = req.Name
	keng.DeviceId = req.DeviceId
	keng.RoomId = req.RoomId
	keng.Index = req.Index
	keng.Remark = req.Remark
	if err := common.DB.Model(&keng).Save(&keng).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	return &res
}

func (cps *ProjectService) KengGet(req *KengGetRequest) *KengGetResponse {
	keng := models.KengModel{}

	if err := common.DB.Table(keng.TableName()).First(&keng, req.KengId).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	res := KengGetResponse{
		Response: common.Response{
			Message: common.HTTP_MESSAGE_OK,
			Code:    common.HTTP_RESPONSE_OK,
		},

		Data: keng,
	}
	return &res
}

func (cps *ProjectService) KengGetList(req *KengGetListRequest) *KengGetListResponse {
	res := KengGetListResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
		Data: make([]models.KengModel, 0),
	}
	db := common.DB.Table(models.KengModel{}.TableName())

	if req.RoomId != 0 {
		db = db.Where("room_id=?", req.RoomId)
	}
	if req.DeviceId != 0 {
		db = db.Where("device_id=?", req.DeviceId)
	}
	if err := db.Offset((req.Offset - 1) * req.Limit).Limit(req.Limit).Find(&res.Data).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}

	return &res
}
