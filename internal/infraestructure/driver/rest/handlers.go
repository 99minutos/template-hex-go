package rest

import (
	"example-service/internal/domain/ports"
	"example-service/internal/infraestructure/driven/core"
	driven_fiber "example-service/internal/infraestructure/driven/fiber_server"
	"net"

	"github.com/gofiber/fiber/v2"
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

func (r *RestHandler) InitializeRoutes() {
	v1 := r.Fiber.Server.Group("/api/v1")
	v1.Get("/health", r.HealthCheck)
	v1.Get("/order/:trackingId", r.GetExample)
	v1.Post("/order/create", r.CreateExample)
}

func (r *RestHandler) Start(is net.Listener) {
	err := r.Fiber.Server.Listener(is)
	log := core.GetDefaultLogger()

	if err != nil {
		log.Fatalw("unable to listen", "error", err)
	}
}

func (r *RestHandler) HealthCheck(fctx *fiber.Ctx) error {
	fctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	return nil
}

func (r *RestHandler) GetExample(fctx *fiber.Ctx) error {
	order, err := r.exService.GetExample(fctx.Context(), fctx.Params("trackingId"))
	if err != nil {
		fctx.Status(fiber.StatusNotFound)
		return fctx.JSON(RestError{Cause: err.Error()})
	}

	return fctx.JSON(order)
}

func (r *RestHandler) CreateExample(fctx *fiber.Ctx) error {
	order, err := r.exService.CreateExample(fctx.Context())
	if err != nil {
		return nil
	}
	return fctx.Status(fiber.StatusCreated).JSON(order)
}
