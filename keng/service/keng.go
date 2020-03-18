package service

import (
	"encoding/json"
	"open_iot/keng/models"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
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
	if req.BuildingId != 0 {
		db = db.Where("building_id=?", req.BuildingId)
	}
	if err := db.Offset((req.Offset - 1) * req.Limit).Limit(req.Limit).Find(&res.Data).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}

	return &res
}

func (cps *ProjectService) KengGetFront(req *KengGetFrontRequest) *KengGetFrontResponse {
	res := KengGetFrontResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
		Data: make([]LocationBuilding, 0),
	}
	redisConn := common.RedisPool.Get()
	defer redisConn.Close()
	if value, err := redis.Bytes(redisConn.Do("Get", "allinfos")); err != nil {
		common.LogrusLogger.Error(err)
		common.LogrusLogger.Info(value)
	} else {
		// 这里获取所有的存在 redis 里的信息，看是否要更新
		allinfos := AllInfoLocations{}
		if err = json.Unmarshal(value, &allinfos); err != nil {
			common.LogrusLogger.Error(err)
		}
		for _, lb := range allinfos.LocationBuilding {
			common.LogrusLogger.Info(lb.LocationId)
			kengs := make([]models.KengModel, 0)
			if err := common.DB.Table(models.KengModel{}.TableName()).Find(&kengs).Error; err != nil {
				common.LogrusLogger.Error(err)
				common.InitKey(cps.Ctx)
				cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
				panic(err)
			}
			for _, bf := range lb.BuildingFloor {
				for _, fr := range bf.FloorRoom {
					for rdIndex, rd := range fr.RoomDevice {
						if len(rd.DeviceGpio) == 0 {
							continue
						}
						kengInfo := make([]KengInfo, 0)
						dg := rd.DeviceGpio[0]
						for _, gi := range dg.GpioInfo {
							for _, keng := range kengs {
								if keng.DeviceGpioId == gi.GpioId {
									var kengTime string
									if gi.GpioTime > 3600 {
										kengTime = "超过1小时"
									} else if gi.GpioTime > 60 {
										mins := int(gi.GpioTime / 60)
										seconds := int(gi.GpioTime) % 60
										kengTime = strconv.Itoa(mins) + "分钟" + strconv.Itoa(seconds) + "秒"
									} else {
										seconds := int(gi.GpioTime) % 60
										kengTime = strconv.Itoa(seconds) + "秒"
									}
									ki := KengInfo{
										KengId:     keng.ID,
										KengIndex:  keng.Index,
										KengName:   keng.Name,
										KengStatus: gi.GpioStatus,
										KengTime:   kengTime,
									}
									kengInfo = append(kengInfo, ki)
									break
								}
							}
						}
						fr.RoomDevice[rdIndex].KengInfo = kengInfo
					}
				}
			}
			res.Data = append(res.Data, lb)
		}

	}
	return &res
}
