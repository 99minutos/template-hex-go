package fiber_server

import (
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
	Cause   interface{} `json:"cause"`
	Message interface{} `json:"message,omitempty"`
}

func NewFiberServer() *FiberServer {
	server := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var cause interface{} = nil
			var message interface{} = nil

			switch e := err.(type) {
			case *ErrorDispatcher:
				code = fiber.StatusUnprocessableEntity
				cause = e.Error()
				message = e.errors
			case *errcodes.ErrorCode:
				code = e.Code
				cause = e.Message
				if e.Description != "" {
					desc := string(e.Description)
					message = &desc
				}
			case *fiber.Error:
				code = e.Code
				cause = e.Error()
			default:
				code = fiber.StatusInternalServerError
				cause = err.Error()
			}

			errCus := RestError{
				Cause:   cause,
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
