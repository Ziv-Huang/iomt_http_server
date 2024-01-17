package api

import (
	"fmt"
	"iomt_http_server/domain/request"

	"github.com/gofiber/fiber/v2"
)

// @Summary		Heartbeat
// @Description	Heartbeat
// @Tags			Heartbeat
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/heartbeat [get]
func Heartbeat(c *fiber.Ctx) error {
	c.Status(200)
	return c.SendString("OK")
}

// @Summary		Notify
// @Description	Notify
// @Tags			Notify
// @Accept			json
// @Produce		json
// @Param payload body request.Notify true "payload"
// @Success		200
// @Router			/Notify [post]
func Notify(c *fiber.Ctx) error {
	var payload request.Notify
	err := c.BodyParser(&payload)
	if err != nil {
		c.Status(400)
		return c.SendString(err.Error())
	}
	c.Status(200)
	return c.SendString(fmt.Sprintf("Phone: %s, Organization: %s", payload.Phone, payload.Organization))
}

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
