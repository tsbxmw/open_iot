package v1

import (
	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/service"
)

func DeviceAdd(c *gin.Context) {
	common.InitKey(c)
	req := service.DeviceAddRequest{}
	if err := c.ShouldBind(&req); err != nil {
		common.LogrusLogger.Error(err)
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.DeviceAdd(&req)

	c.JSON(200, &res)
}

func DeviceUpdate(c *gin.Context) {
	common.InitKey(c)
	req := service.DeviceUpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		common.LogrusLogger.Error(err)
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	deviceId := c.Params.ByName("id")
	res := cps.DeviceUpdate(deviceId, &req)
	c.JSON(common.HTTP_STATUS_OK, &res)
}

func IPUpdate(c *gin.Context) {
	common.InitKey(c)
	req := service.IPUpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		common.LogrusLogger.Error(err)
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.IPUpdate(&req)
	c.JSON(common.HTTP_STATUS_OK, &res)
}

func DeviceGet(c *gin.Context) {
	common.InitKey(c)
	req := service.DeviceGetRequest{}
	if err := c.ShouldBind(&req); err != nil {
		common.LogrusLogger.Error(err)
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}

	cps := service.NewServiceMgr(c)
	res := cps.DeviceGet(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}

func DeviceGetList(c *gin.Context) {
	common.InitKey(c)
	req := service.DeviceGetListRequest{}

	if err := c.ShouldBind(&req); err != nil {
		common.LogrusLogger.Error(err)
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.DeviceGetList(&req)
	c.JSON(common.HTTP_STATUS_OK, &res)
}
