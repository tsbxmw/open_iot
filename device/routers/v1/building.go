package v1

import (
	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/service"
)

func BuildingAdd(c *gin.Context) {
	common.InitKey(c)
	req := service.BuildingAddRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.BuildingAdd(&req)

	c.JSON(200, &res)
}

func BuildingUpdate(c *gin.Context) {
	common.InitKey(c)
	req := service.BuildingUpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	buildingId := c.Params.ByName("id")
	res := cps.BuildingUpdate(buildingId, &req)
	c.JSON(200, &res)
}

func BuildingGet(c *gin.Context) {
	common.InitKey(c)
	req := service.BuildingGetRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}

	cps := service.NewServiceMgr(c)
	res := cps.BuildingGet(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}

func BuildingGetList(c *gin.Context) {
	common.InitKey(c)
	req := service.BuildingGetListRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.BuildingGetList(&req)
	c.JSON(common.HTTP_STATUS_OK, &res)
}
