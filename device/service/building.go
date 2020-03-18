package service

import (
	"open_iot/device/models"
	"time"

	common "github.com/tsbxmw/gin_common"
)

// 增加建筑
func (cps *ProjectService) BuildingAdd(req *BuildingAddRequest) *BuildingAddResponse {
	res := BuildingAddResponse{}
	building := models.BuildingModel{}
	// 检查检查建筑是否存在
	if err := common.DB.Table(building.TableName()).
		Where("name=? and location_id=?", req.Name, req.LocationId).
		First(&building).Error; err != nil {
		if err.Error() == "record not found" {
			// 不存在则增加
			building.Name = req.Name
			building.LocationId = req.LocationId
			building.Remark = req.Remark
			building.BaseModel.ModifiedTime = time.Now()
			building.BaseModel.CreationTime = time.Now()

			if err = common.DB.Table(building.TableName()).
				Create(&building).Error; err != nil {
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
		res.Message = "Building already exists"
	}
	return &res
}

// 更新建筑信息
func (cps *ProjectService) BuildingUpdate(id string, req *BuildingUpdateRequest) *BuildingUpdateResponse {
	res := BuildingUpdateResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
	}
	building := models.BuildingModel{}
	// 查找当前建筑是否存在
	if err := common.DB.Table(building.TableName()).First(&building, id).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	// 更新名称
	if req.Name != "" {
		building.Name = req.Name
	}
	// 更新地区 id
	if req.LocationId != 0 {
		building.LocationId = req.LocationId
	}
	// 更新 备注
	if req.Remark != "" {
		building.Remark = req.Remark
	}
	building.BaseModel.ModifiedTime = time.Now()
	building.BaseModel.CreationTime = time.Now()

	// 创建建筑信息
	if err := common.DB.Table(building.TableName()).
		Create(&building).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_CREATE_ERROR
		panic(err)
	}
	return &res
}

// 获取单个建筑信息
func (cps *ProjectService) BuildingGet(req *BuildingGetRequest) *BuildingGetResponse {
	building := models.BuildingModel{}

	if err := common.DB.Table(building.TableName()).
		Find(&building, req.BuildingId).Error; err != nil {
		common.LogrusLogger.Error(err)
		common.InitKey(cps.Ctx)
		cps.Ctx.Keys["code"] = common.MYSQL_QUERY_ERROR
		panic(err)
	}
	res := BuildingGetResponse{
		Response: common.Response{
			Message: common.HTTP_MESSAGE_OK,
			Code:    common.HTTP_RESPONSE_OK,
		},
		Data: building,
	}
	return &res
}

// 获取建筑列表
func (cps *ProjectService) BuildingGetList(req *BuildingGetListRequest) *BuildingGetListResponse {
	res := BuildingGetListResponse{
		Response: common.Response{
			Code:    common.HTTP_RESPONSE_OK,
			Message: common.HTTP_MESSAGE_OK,
		},
		Count:  0,
		Data:   make([]models.BuildingModel, 0),
		Limmit: req.Limit,
		Offset: req.Offset,
	}
	db := common.DB.Table(models.BuildingModel{}.TableName())

	if req.LocationId != 0 {
		db = db.Where("location_id=?", req.LocationId)
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
