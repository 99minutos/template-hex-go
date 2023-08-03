package rest

import (
	"example-service/internal/application/ports"
	"example-service/internal/config"
	driven_fiber "example-service/internal/infraestructure/adapters/driven/fiber"
	"example-service/internal/infraestructure/adapters/driven/logger"
	"github.com/gofiber/fiber/v2"
	"net"
)

type RestError struct {
	Cause string `json:"Cause"`
}
type RestHandler struct {
	Fiber     *driven_fiber.FiberServer
	exService ports.IExampleService
}

func NewRestHandler(exService ports.IExampleService) *RestHandler {
	Fiber := driven_fiber.NewFiberServer()
	return &RestHandler{
		Fiber:     Fiber,
		exService: exService,
	}
}

func (r *RestHandler) InitializeRoutes(config config.AppConfig) {
	v1 := r.Fiber.Server.Group("/api/v1")
	v1.Get("/order/:trackingId", r.GetExample)
	v1.Post("/order/create", r.CreateExample)
}

func (r *RestHandler) Start(is net.Listener) {
	err := r.Fiber.Server.Listener(is)

	if err != nil {
		logger.Logger.Fatal(err)
	}
}

func (r *RestHandler) GetExample(fctx *fiber.Ctx) error {
	order, err := r.exService.GetExample(fctx.Context(), fctx.Params("exampleId"))
	if err != nil {
		fctx.Status(fiber.StatusNotFound)
		return fctx.JSON(RestError{Cause: err.Error()})
	}

	return fctx.JSON(order)
}

func (r *RestHandler) CreateExample(fctx *fiber.Ctx) error {
	err := r.exService.CreateExample(fctx.Context())
	if err != nil {

		return nil
	}

	fctx.SendStatus(200)
	fctx.SendString("OK")
	return nil
}
