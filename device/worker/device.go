package worker

import (
	"open_iot/device/models"
	"time"

	"github.com/robfig/cron"
	common "github.com/tsbxmw/gin_common"
)

func CornWork() {
	common.LogrusLogger.Info("Corn work start")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", MessageSendCheckWork)
	c.Start()
	select {}
}

type (
	AllInfoLocations struct {
		LocationBuilding []LocationBuilding `json:"locations"`
	}
	LocationBuilding struct {
		LocationId    int             `json:"location_id"`
		LocationName  string          `json:"location_name"`
		BuildingFloor []BuildingFloor `json:"buildings`
	}

	BuildingFloor struct {
		BuildingId   int         `json:"building_id"`
		BuildingName string      `json:"building_name"`
		FloorRoom    []FloorRoom `json:"floors"`
	}
	FloorRoom struct {
		FloorId    int          `json:"floor_id"`
		FloorName  string       `json:"floor_name"`
		RoomDevice []RoomDevice `json:"rooms`
	}
	RoomDevice struct {
		RoomId     int          `json:"room_id"`
		RoomName   string       `json:"room_name"`
		DeviceGpio []DeviceGpio `json:"device"`
	}

	DeviceGpio struct {
		DeviceId   int    `json:"device_id"`
		DeviceName string `json:"device_name"`
		GpioInfo   []Gpio `json:"gpio_info"`
	}

	Gpio struct {
		GpioId     int     `json:"gpio_id"`
		GpioNumber int     `json:"gpio_number"`
		GpioStatus int     `json:"gpio_status"`
		GpioTime   float64 `json:"gpio_time"`
	}
)

func MessageSendCheckWork() {
	common.LogrusLogger.Info("定期更新所有数据到 redis，防止多次查询")
	allinfos := AllInfoLocations{
		LocationBuilding: make([]LocationBuilding, 0),
	}
	// 获取所有的地区信息
	locations := make([]models.LocationModel, 0)
	if err := common.DB.Table(models.LocationModel{}.TableName()).Find(&locations).Error; err != nil {
		common.LogrusLogger.Error(err)
	}

	for _, location := range locations {
		lb := LocationBuilding{
			BuildingFloor: make([]BuildingFloor, 0),
		}
		// 获取当地的建筑信息
		buildings := make([]models.BuildingModel, 0)
		if err := common.DB.Table(models.BuildingModel{}.TableName()).Where("location_id=?", location.ID).Find(&buildings).Error; err != nil {
			common.LogrusLogger.Error(err)
		}
		lb.LocationId = location.ID
		lb.LocationName = location.Name
		for _, building := range buildings {
			bf := BuildingFloor{
				FloorRoom: make([]FloorRoom, 0),
			}
			// 获取建筑的楼层信息
			floors := make([]models.FloorModel, 0)
			if err := common.DB.Table(models.FloorModel{}.TableName()).Where("building_id=?", building.ID).Find(&floors).Error; err != nil {
				common.LogrusLogger.Error(err)
			}
			bf.BuildingId = building.ID
			bf.BuildingName = building.Name
			for _, floor := range floors {
				fr := FloorRoom{
					RoomDevice: make([]RoomDevice, 0),
				}
				// 获取楼层所有的房间信息
				rooms := make([]models.RoomModel, 0)
				if err := common.DB.Table(models.RoomModel{}.TableName()).Where("floor_id=?", floor.ID).Find(&rooms).Error; err != nil {
					common.LogrusLogger.Error(err)
				}
				fr.FloorId = floor.ID
				fr.FloorName = floor.Name
				for _, room := range rooms {
					rd := RoomDevice{
						DeviceGpio: make([]DeviceGpio, 0),
					}
					// 获取所有房间的设备信息
					devices := make([]models.DeviceModel, 0)
					if err := common.DB.Table(models.DeviceModel{}.TableName()).Where("room_id=?", room.ID).Find(&devices).Error; err != nil {
						common.LogrusLogger.Error(err)
					}
					rd.RoomId = room.ID
					rd.RoomName = room.Name
					for _, device := range devices {
						dg := DeviceGpio{
							GpioInfo: make([]Gpio, 0),
						}
						// 获取设备的所有 gpio 信息
						gpios := make([]models.DeviceGpioModel, 0)
						if err := common.DB.Table(models.DeviceGpioModel{}.TableName()).Where("device_id=?", device.ID).Find(&gpios).Error; err != nil {
							common.LogrusLogger.Error(err)
						}

						dg.DeviceId = device.ID
						dg.DeviceName = device.Name

						for _, gpio := range gpios {
							gpio_record := models.DeviceGpioRecordModel{}
							if err := common.DB.Table(gpio_record.TableName()).Where("gpio_id=?", gpio.ID).Last(&gpio_record).Error; err != nil {
								common.LogrusLogger.Error(gpio.ID, err)
							}
							gi := Gpio{}
							gi.GpioId = gpio.ID
							gi.GpioNumber = gpio.GpioNumber
							gi.GpioStatus = gpio.GpioStatus
							now := time.Now()
							duration := now.Sub(gpio_record.CreationTime)
							gi.GpioTime = duration.Seconds()
							dg.GpioInfo = append(dg.GpioInfo, gi)
						}
						rd.DeviceGpio = append(rd.DeviceGpio, dg)
					}
					fr.RoomDevice = append(fr.RoomDevice, rd)
				}
				bf.FloorRoom = append(bf.FloorRoom, fr)
			}
			lb.BuildingFloor = append(lb.BuildingFloor, bf)
		}
		redisConn := common.RedisPool.Get()
		defer redisConn.Close()
		// 分别更新地区信息到 redis
		if code, err := common.RedisSetCommon(redisConn, lb.LocationName, lb); err != nil {
			common.LogrusLogger.Error(code, err)
		}
		allinfos.LocationBuilding = append(allinfos.LocationBuilding, lb)
	}
	redisConn := common.RedisPool.Get()
	defer redisConn.Close()
	// 更新信息到redis
	if code, err := common.RedisSetCommon(redisConn, "allinfos", allinfos); err != nil {
		common.LogrusLogger.Error(code, err)
	}
}
