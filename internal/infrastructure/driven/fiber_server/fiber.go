package fiber_server

import (
	"errors"
	"fmt"
	"service/internal/domain/errcodes"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	Server *fiber.App
}

type RestError struct {
	Cause   string      `json:"cause"`
	Message interface{} `json:"message,omitempty"`
}

func NewFiberServer() *FiberServer {
	server := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var message interface{} = nil

			var eCus *errcodes.ErrorCode
			var fibErr *fiber.Error
			var errDis *ErrorDispatcher

			if errors.As(err, &errDis) {
				code = fiber.StatusUnprocessableEntity
				message = errDis.errors
			} else if errors.As(err, &eCus) {
				code = eCus.Code
				if eCus.Description != "" {
					desc := string(eCus.Description)
					message = &desc
				}
			} else if errors.As(err, &fibErr) {
				code = fibErr.Code
				message = fibErr.Message
			} else {
				code = fiber.StatusInternalServerError
				message = err.Error()
			}

			errCus := RestError{
				Cause:   err.Error(),
				Message: message,
			}
			return ctx.Status(code).JSON(errCus)
		},
	})
	server.Use(cors.New())
	return &FiberServer{
		Server: server,
	}
}

func (f *FiberServer) Start(port string) error {
	return f.Server.Listen(fmt.Sprintf(":%v", port))
}
