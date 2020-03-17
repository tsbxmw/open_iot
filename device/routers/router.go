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

			location := ManagementV1.Group("/location")
			{
				location.POST("/", v1.LocationAdd)
				location.POST("/:id", v1.LocationUpdate)
				location.GET("/", v1.LocationGet)
				location.GET("/list", v1.LocationGetList)
			}

			building := ManagementV1.Group("/building")
			{
				building.POST("/", v1.BuildingAdd)
				building.POST("/:id", v1.BuildingUpdate)
				building.GET("/", v1.BuildingGet)
				building.GET("/list", v1.BuildingGetList)
			}

			floor := ManagementV1.Group("/floor")
			{
				floor.POST("/", v1.FloorAdd)
				floor.POST("/:id", v1.FloorUpdate)
				floor.GET("/", v1.FloorGet)
				floor.GET("/list", v1.FloorGetList)
			}

			room := ManagementV1.Group("/room")
			{
				room.POST("/", v1.RoomAdd)
				room.POST("/:id", v1.RoomUpdate)
				room.GET("/", v1.RoomGet)
				room.GET("/list", v1.RoomGetList)
			}
		}

		DeviceV1 := GroupV1.Group("/iot")
		{
			device := DeviceV1.Group("/device")
			{
				device.POST("/ip", v1.IPUpdate)
				device.POST("/switch", v1.SwitchUpdate)
			}
		}

	}
}
