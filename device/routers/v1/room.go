package v1

import (
	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
	"open_iot/device/service"
)

func RoomAdd(c *gin.Context) {
	common.InitKey(c)
	req := service.RoomAddRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.RoomAdd(&req)

	c.JSON(200, &res)
}

func RoomUpdate(c *gin.Context) {
	common.InitKey(c)
	req := service.RoomUpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	roomId := c.Params.ByName("id")
	res := cps.RoomUpdate(roomId, &req)
	c.JSON(200, &res)
}

func RoomGet(c *gin.Context) {
	common.InitKey(c)
	req := service.RoomGetRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}

	cps := service.NewServiceMgr(c)
	res := cps.RoomGet(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}

func RoomGetList(c *gin.Context) {
	common.InitKey(c)
	req := service.RoomGetListRequest{}

	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.RoomGetList(&req)
	c.JSON(common.HTTP_STATUS_OK, &res)
}
