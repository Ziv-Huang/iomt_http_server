package api

import (
	"fmt"
	"iomt_http_server/domain/configs"
	log "iomt_http_server/domain/logger"
	"iomt_http_server/domain/request"
	"iomt_http_server/usecase/statushandler"

	"github.com/gofiber/fiber/v2"
)

type API interface {
	Heartbeat(c *fiber.Ctx) error
	IoMT_task(c *fiber.Ctx) error
	IoMT_status(c *fiber.Ctx) error
}

type api struct {
	config        configs.Configuration
	statushandler statushandler.StatusHandler
}

func NewAPI(config configs.Configuration) API {
	return &api{config: config, statushandler: statushandler.NewStatusHandler(config.Http.Waittime)}
}

// @Summary		Heartbeat
// @Description	Heartbeat
// @Tags			Heartbeat
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/heartbeat [get]
func (a *api) Heartbeat(c *fiber.Ctx) error {
	log.Info(fmt.Sprintf("Heartbeat, IP: %s", c.IP()))
	c.Status(200)
	return c.SendString("OK")
}

// @Summary		IoMT_task
// @Description	IoMT_task
// @Tags			IoMT_task
// @Accept			json
// @Produce		json
// @Param payload body request.Notify true "payload"
// @Success		200
// @Router			/api/v2/troom/phone/IoMT_task [post]
func (a *api) IoMT_task(c *fiber.Ctx) error {
	log.Info(fmt.Sprintf("IoMT_task, IP: %s", c.IP()))
	a.statushandler.Countdown()
	c.Status(200)
	return c.SendString("OK")
}

// @Summary		IoMT_status
// @Description	IoMT_status
// @Tags			IoMT_status
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/api/v2/troom/bulletin/IoMT_status [get]
func (a *api) IoMT_status(c *fiber.Ctx) error {
	log.Info(fmt.Sprintf("IoMT_status, IP: %s", c.IP()))
	res := request.IoMTStatus{
		Suspend: a.statushandler.GetStatus(),
	}
	c.Status(200)
	return c.JSON(res)
}

// // @Summary		Notify
// // @Description	Notify
// // @Tags			Notify
// // @Accept			json
// // @Produce		json
// // @Param payload body request.Notify true "payload"
// // @Success		200
// // @Router			/Notify [post]
// func Notify(c *fiber.Ctx) error {
// 	var payload request.Notify
// 	err := c.BodyParser(&payload)
// 	if err != nil {
// 		c.Status(400)
// 		return c.SendString(err.Error())
// 	}
// 	c.Status(200)
// 	return c.JSON(payload)
// }

// // @Summary		Hello
// // @Description	Hello World
// // @Tags			Hello
// // @Accept			json
// // @Produce		json
// // @Success		200
// // @Router			/ [get]
// func Hello(c *fiber.Ctx) error {
// 	c.Status(200)
// 	return c.SendString("Hello, World 👋!")
// }

// // @Summary		Get Device Info
// // @Description	Get Device Info
// // @Tags			Get Device Info
// // @Accept			json
// // @Produce		json
// // @Success		200
// // @Router			/deviceinfo [get]
// func GetDeviceInfo(c *fiber.Ctx) error {
// 	// TODO
// 	var result request.DeviceInfo
// 	result.IsBusy = true
// 	result.IsAlive = true
// 	c.Status(200)
// 	return c.JSON(result)
// }

// // @Summary		Create A Job
// // @Description	Create A Job
// // @Tags			Create A Job
// // @Accept			json
// // @Produce		json
// // @Success		200
// // @Router			/job [post]
// func CreateAJob(c *fiber.Ctx) error {
// 	// TODO
// 	var result request.Task
// 	result.TaskID = "1"
// 	result.JobStatus = request.Running
// 	result.Msg = "OK"
// 	result.Code = 200
// 	c.Status(200)
// 	return c.JSON(result)
// }

// // @Summary		Delete A Job
// // @Description	Delete A Job
// // @Tags			Delete A Job
// // @Accept			json
// // @Produce		json
// // @Success		200
// // @Router			/job/:id [delete]
// func DeleteAJob(c *fiber.Ctx) error {
// 	// TODO
// 	var result request.Task
// 	result.TaskID = c.Params("id")
// 	result.JobStatus = request.Killed
// 	result.Msg = "OK"
// 	result.Code = 200
// 	c.Status(200)
// 	return c.JSON(result)
// }

// // @Summary		Get Result
// // @Description	Get Result
// // @Tags			Get Result
// // @Accept			json
// // @Produce		json
// // @Success		200
// // @Router			/result/:id [get]
// func GetResultByID(c *fiber.Ctx) error {
// 	// TODO
// 	var result request.Result
// 	result.TaskID = c.Params("id")
// 	result.JobStatus = request.Running
// 	result.Msg = "OK"
// 	result.Code = 200
// 	result.AnalysisResult = "OK"
// 	c.Status(200)
// 	return c.JSON(result)
// }
