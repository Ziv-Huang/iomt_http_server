package delivery

import (
	"iomt_http_server/domain/configs"
	"iomt_http_server/usecase/api"
	"sync"

	_ "iomt_http_server/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware

	"github.com/gofiber/fiber/v2"
)

type httpFiber struct {
	Wg     *sync.WaitGroup
	config configs.Configuration
	api    api.API
}

func NewHttpFiber(wg *sync.WaitGroup, config configs.Configuration) *httpFiber {
	return &httpFiber{Wg: wg, config: config, api: api.NewAPI(config)}
}

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8123
// @BasePath /
// @schemes http
func (s *httpFiber) Start() {

	go func() {
		defer s.Wg.Done()

		// Fiber instance
		app := fiber.New()
		troom := app.Group("/api/v2/troom")
		phone := troom.Group("/phone")
		bulletin := troom.Group("/bulletin")

		// Fiber operations
		app.Get("/heartbeat", s.api.Heartbeat)
		bulletin.Get("iomt_status", s.api.IoMT_status)
		phone.Post("iomt_task", s.api.IoMT_task)

		// app.Get("/", api.Hello)
		// app.Post("/notify", api.Notify)
		// app.Get("/deviceinfo", api.GetDeviceInfo)
		// app.Get("/result/:id", api.GetResultByID)
		// app.Post("/job", api.CreateAJob)
		// app.Delete("/job/:id", api.DeleteAJob)

		// Fiber swagger
		app.Get("/swagger/*", fiberSwagger.WrapHandler)

		// Start server
		app.Listen(":" + s.config.Http.Port)
	}()

}
