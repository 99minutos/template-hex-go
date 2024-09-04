package fiber_server

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	Server *fiber.App
}

type RestError struct {
	Cause string `json:"cause"`
}

func NewFiberServer() *FiberServer {
	server := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return ctx.Status(code).JSON(RestError{Cause: err.Error()})
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
