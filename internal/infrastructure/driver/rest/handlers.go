package rest

import (
	"net"

	"service/internal/implementation/example"
	"service/internal/infrastructure/driven/dbg"
	driven_fiber "service/internal/infrastructure/driven/fiber_server"

	"github.com/gofiber/fiber/v2"
)

type RestError struct {
	Cause string `json:"Cause"`
}
type RestHandler struct {
	Fiber     *driven_fiber.FiberServer
	exService *example.ExampleService
}

func NewRestHandler(exService *example.ExampleService) *RestHandler {
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
	log := dbg.GetLogger()

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
