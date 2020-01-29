package v1

import (
	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/service"
)

func FloorAdd(c *gin.Context) {
	common.InitKey(c)
	req := service.FloorAddRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.FloorAdd(&req)

	c.JSON(200, &res)
}

func FloorUpdate(c *gin.Context) {
	common.InitKey(c)
	req := service.FloorUpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	floorId := c.Params.ByName("id")
	res := cps.FloorUpdate(floorId, &req)
	c.JSON(200, &res)
}

func FloorGet(c *gin.Context) {
	common.InitKey(c)
	req := service.FloorGetRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}

	cps := service.NewServiceMgr(c)
	res := cps.FloorGet(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}

func FloorGetList(c *gin.Context) {
	common.InitKey(c)
	req := service.FloorGetListRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.FloorGetList(&req)
	c.JSON(common.HTTP_STATUS_OK, &res)
}
