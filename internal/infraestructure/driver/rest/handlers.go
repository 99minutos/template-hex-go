package rest

import (
	"example-service/internal/domain/core"
	"example-service/internal/domain/ports"
	driven_fiber "example-service/internal/infraestructure/driven/fiber"
	"github.com/gofiber/fiber/v2"
	"net"
)

type RestError struct {
	Cause string `json:"Cause"`
}
type RestHandler struct {
	acx       *core.AppContext
	Fiber     *driven_fiber.FiberServer
	exService ports.IExampleService
}

func NewRestHandler(acx *core.AppContext, exService ports.IExampleService) *RestHandler {
	Fiber := driven_fiber.NewFiberServer()
	return &RestHandler{
		acx:       acx,
		Fiber:     Fiber,
		exService: exService,
	}
}

func (r *RestHandler) InitializeRoutes(config *core.AppConfig) {
	v1 := r.Fiber.Server.Group("/api/v1")
	v1.Get("/order/:trackingId", r.GetExample)
	v1.Post("/order/create", r.CreateExample)
}

func (r *RestHandler) Start(is net.Listener) {
	err := r.Fiber.Server.Listener(is)

	if err != nil {
		r.acx.Fatalw("unable to listen", "error", err)
	}
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
