package service

import (
	"encoding/json"
	"open_iot/device/models"
	"time"

	"github.com/garyburd/redigo/redis"
	common "github.com/tsbxmw/gin_common"
)

func (cps *ProjectService) DeviceAdd(req *DeviceAddRequest) *DeviceAddResponse {
	res := DeviceAddResponse{}
	device := models.DeviceModel{}
	if err := common.DB.Table(device.TableName()).
		Where("mac_address=?", req.MacAddress).First(&device).Error; err != nil {
		if err.Error() == "record not found" {
			device.Name = req.Name
			device.DeviceType = req.DeviceType
			device.MacAddress = req.MacAddress
			device.LocationId = req.LocationId
			device.BuildingId = req.BuildingId
			device.FloorId = req.FloorId
			device.RoomId = req.RoomId
			device.Remark = req.Remark
			device.BaseModel.ModifiedTime = time.Now()
			device.BaseModel.CreationTime = time.Now()

			if err = common.DB.Table(device.TableName()).Create(&device).Error; err != nil {
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
		res.Message = "Device already exists"
	}
	return &res
}

func (cps *ProjectService) DeviceUpdate(id string, req *DeviceUpdateRequest) *DeviceUpdateResponse {
	res := DeviceUpdateResponse{}
	device := models.DeviceModel{}
	if err := common.DB.Table(device.TableName()).First(&device, id).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
		panic(err)
	}
	if req.Name != "" {
		device.Name = req.Name
	}
	device.DeviceType = req.DeviceType
	if req.MacAddress != "" {
		device.MacAddress = req.MacAddress
	}
	if req.LocationId != 0 {
		device.LocationId = req.LocationId
	}
	if req.BuildingId != 0 {
		device.BuildingId = req.BuildingId
	}
	if req.FloorId != 0 {
		device.FloorId = req.FloorId
	}
	if req.RoomId != 0 {
		device.RoomId = req.RoomId
	}
	if req.Remark != "" {
		device.Remark = req.Remark
	}
	device.BaseModel.ModifiedTime = time.Now()
	device.BaseModel.CreationTime = time.Now()

	if err := common.DB.Model(&device).Save(&device).Error; err != nil {
		cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
		panic(err)
	} else {
		res.Code = common.HTTP_RESPONSE_OK
		res.Message = common.HTTP_MESSAGE_OK
		res.Data = []string{}
	}
	return &res
}

func (cps *ProjectService) IPUpdate(req *IPUpdateRequest) *IPUpdateResponse {
	res := IPUpdateResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	redisConn := common.RedisPool.Get()
	defer redisConn.Close()
	if value, err := common.RedisGetCommon(redisConn, req.MacAddress); err != nil {
		common.LogrusLogger.Error(err)
		if _, err := common.RedisSetCommon(redisConn, req.MacAddress, req.IpAddress); err != nil {
			common.LogrusLogger.Error(err)
		}
	} else {
		if value == req.IpAddress {
			return &res
		} else if _, err := common.RedisSetCommon(redisConn, req.MacAddress, req.IpAddress); err != nil {
			common.LogrusLogger.Error(err)
		}
	}
	device := models.DeviceModel{}

	if err := common.DB.Table(device.TableName()).
		Where("mac_address=?", req.MacAddress).First(&device).Error; err == nil {
		device.IpAddress = req.IpAddress
		common.LogrusLogger.Info(device.IpAddress)
		if err := common.DB.Model(&device).Save(&device).Error; err != nil {
			common.LogrusLogger.Error(err)
			common.InitKey(cps.Ctx)
			cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
			panic(err)
		}
	} else {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
		panic(err)
	}

	return &res
}

func (cps *ProjectService) SwitchUpdate(req *SwitchUpdateRequest) *SwitchUpdateResponse {
	res := SwitchUpdateResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}

	redisConn := common.RedisPool.Get()
	defer redisConn.Close()

	if value, err := redis.Bytes(redisConn.Do("Get", req.MacAddress)); err != nil {
		common.LogrusLogger.Info(value)
	} else {
		// 这里获取所有的存在 redis 里的信息，看是否要更新
		sur := SwitchUpdateRequest{}
		if err = json.Unmarshal(value, &sur); err != nil {
			common.LogrusLogger.Error(err)
		} else {
			all_same_flag := false
			for _, info := range sur.GpioInfos {
				for _, req_info := range req.GpioInfos {
					if info.GpioNumber == req_info.GpioNumber {
						if info.GpioStatus != req_info.GpioStatus {
							all_same_flag = true
						}
						break
					}
				}
			}
			if !all_same_flag && sur.IpAddress == req.IpAddress { // 如果是相同的，不必入库和更新 redis，不同则更新
				return &res
			}
		}
	}
	// 如果出现错误，则入 redis
	if _, err := common.RedisSetCommon(redisConn, req.MacAddress, &req); err != nil {
		common.LogrusLogger.Error(err)
	}

	// 更新 device 的 ip address
	device := models.DeviceModel{}
	if err := common.DB.Table(device.TableName()).
		Where("mac_address=?", req.MacAddress).First(&device).Error; err == nil {
		device.IpAddress = req.IpAddress
		if err := common.DB.Model(&device).Save(&device).Error; err != nil {
			common.LogrusLogger.Error(err)
			common.InitKey(cps.Ctx)
			cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
			panic(err)
		}
	} else {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
		panic(err)
	}

	// 更新 gpio 的状态信息
	for _, gpio_req := range req.GpioInfos {
		gpio := models.DeviceGpioModel{}
		if err := common.DB.Table(gpio.TableName()).
			Where("device_id=? and gpio_number=?", device.ID, gpio_req.GpioNumber).First(&gpio).Error; err != nil {
			common.LogrusLogger.Error(err)
			common.InitKey(cps.Ctx)
			cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
			panic(err)
		} else {
			if gpio_req.GpioStatus != gpio.GpioStatus {
				gpio.GpioStatus = gpio_req.GpioStatus
				if err := common.DB.Table(gpio.TableName()).Save(&gpio).Error; err != nil {
					common.LogrusLogger.Error(err)
					common.InitKey(cps.Ctx)
					cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
					panic(err)
				}
				// 更新记录
				gpio_record := models.DeviceGpioRecordModel{
					GpioId:     gpio.ID,
					DeviceId:   gpio.DeviceId,
					GpioNumber: gpio.GpioNumber,
					GpioStatus: gpio.GpioStatus,
					BaseModel: common.BaseModel{
						CreationTime: time.Now(),
						ModifiedTime: time.Now(),
					},
				}
				if err := common.DB.Table(gpio_record.TableName()).Create(&gpio_record).Error; err != nil {
					common.LogrusLogger.Error(err)
					common.InitKey(cps.Ctx)
					cps.Ctx.Keys["code"] = common.MYSQL_UPDATE_ERROR
					panic(err)
				}
			}
		}
	}

	return &res
}

func (cps *ProjectService) DeviceGet(req *DeviceGetRequest) *DeviceGetResponse {
	device := models.DeviceModel{}

	if err := common.DB.Find(&device, req.DeviceId).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	res := DeviceGetResponse{
		Response: common.Response{
			Message: common.HTTP_MESSAGE_OK,
			Code:    common.HTTP_RESPONSE_OK,
		},
		Data: device,
	}
	return &res
}

func (cps *ProjectService) DeviceGetList(req *DeviceGetListRequest) *DeviceGetListResponse {
	res := DeviceGetListResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
		Count:  0,
		Data:   make([]models.DeviceModel, 0),
		Limmit: req.Limit,
		Offset: req.Offset,
	}
	db := common.DB.Table(models.DeviceModel{}.TableName())
	if req.LocationId != 0 {
		db = db.Where("location_id=?", req.LocationId)
	}
	if req.BuildingId != 0 {
		db = db.Where("building_id=?", req.BuildingId)
	}
	if req.FloorId != 0 {
		db = db.Where("floor_id=?", req.FloorId)
	}
	if req.RoomId != 0 {
		db = db.Where("room_id=?", req.RoomId)
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
