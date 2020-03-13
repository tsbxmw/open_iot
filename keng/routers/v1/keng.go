package v1

import (
	"open_iot/keng/service"

	"github.com/gin-gonic/gin"
	common "github.com/tsbxmw/gin_common"
)

func KengAdd(c *gin.Context) {
	common.InitKey(c)
	req := service.KengAddRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	cps := service.NewServiceMgr(c)
	res := cps.KengAdd(&req)

	c.JSON(200, &res)
}

func KengUpdate(c *gin.Context) {
	common.InitKey(c)
	req := service.KengUpdateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}
	kengId := c.Params.ByName("id")
	cps := service.NewServiceMgr(c)
	res := cps.KengUpdate(kengId, &req)

	c.JSON(200, &res)
}

func KengGet(c *gin.Context) {
	common.InitKey(c)
	req := service.KengGetRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}

	cps := service.NewServiceMgr(c)
	res := cps.KengGet(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}

func KengGetList(c *gin.Context) {
	common.InitKey(c)
	req := service.KengGetListRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.Keys["code"] = common.HTTP_MISS_PARAMS
		panic(err)
	}

	cps := service.NewServiceMgr(c)
	res := cps.KengGetList(&req)

	c.JSON(common.HTTP_STATUS_OK, &res)
}

func KengGetFront(c *gin.Context) {
	common.InitKey(c)
	req := service.KengGetFrontRequest{}

}
