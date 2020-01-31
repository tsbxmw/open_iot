package v1

import (
	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/service"
)

func LocationAdd(c *gin.Context) {
	common.InitKey(c)
	req := service.LocationAddRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.LocationAdd(&req)

	c.JSON(200, &res)
}

func LocationUpdate(c *gin.Context) {
	common.InitKey(c)
	req := service.LocationUpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	locationId := c.Params.ByName("id")
	res := cps.LocationUpdate(locationId, &req)
	c.JSON(200, &res)
}

func LocationGet(c *gin.Context) {
	common.InitKey(c)
	req := service.LocationGetRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.LocationGet(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}

func LocationGetList(c *gin.Context) {
	common.InitKey(c)
	req := service.LocationGetListRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.LocationGetList(&req)
	c.JSON(common.HTTP_STATUS_OK, &res)
}
