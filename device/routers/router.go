package routers

import (
	v1 "open_iot/device/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	GroupV1 := r.Group("/v1")
	{
		ManagementV1 := GroupV1.Group("/management")
		{
			user := ManagementV1.Group("/user")
			{
				user.POST("/", v1.UserAdd)
				user.GET("/", v1.UserGet)
			}

			device := ManagementV1.Group("/device")
			{
				device.POST("/", v1.DeviceAdd)
				device.POST("/:id", v1.DeviceUpdate)
				device.GET("/", v1.DeviceGet)
				device.GET("/list", v1.DeviceGetList)
			}

			location := ManagementV1.Group("/location")
			{
				location.POST("/", v1.LocationAdd)
				location.POST("/:id", v1.LocationUpdate)
				location.GET("/", v1.LocationGet)
				location.GET("/list", v1.LocationGetList)
			}
		}

		DeviceV1 := GroupV1.Group("/iot")
		{
			device := DeviceV1.Group("/device")
			{
				device.POST("/ip", v1.IPUpdate)
			}
		}

	}

}
