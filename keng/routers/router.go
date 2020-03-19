package routers

import (
	v1 "open_iot/keng/routers/v1"

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

		management := GroupV1.Group("/management")
		{
			keng := management.Group("/keng")
			{
				keng.POST("/", v1.KengAdd)
				keng.POST("/:id", v1.KengUpdate)
				keng.GET("/", v1.KengGet)
				keng.GET("/list", v1.KengGetList)

			}
		}

		frontPage := GroupV1.Group("/front")
		{
			keng := frontPage.Group("/keng")
			{
				keng.GET("/", v1.KengGetFront)
				keng.GET("/history", v1.KengGetFrontHistory)
			}
		}
	}
}
