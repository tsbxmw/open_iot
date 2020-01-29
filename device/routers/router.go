package routers

import (
	v1 "open_iot/device/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	GroupV1 := r.Group("/v1")
	{
		user := GroupV1.Group("/user")
		{
			user.POST("/", v1.UserAdd)
			user.GET("/", v1.UserGet)
		}

		device := GroupV1.Group("/device")
		{
			device.POST("/", v1.DeviceAdd)
			device.POST("/:id", v1.DeviceUpdate)
			device.GET("/", v1.DeviceGet)
			device.GET("/list", v1.DeviceGetList)
		}

		location := GroupV1.Group("/location")
		{
			location.POST("/", v1.LocationAdd)
			location.POST("/:id", v1.LocationUpdate)
			location.GET("/", v1.LocationGet)
			location.GET("/list", v1.LocationGetList)
		}

	}
}
