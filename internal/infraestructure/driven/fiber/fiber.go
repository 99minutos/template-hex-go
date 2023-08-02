package fiber

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	Server *fiber.App
}

func NewFiberServer() *FiberServer {
	server := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	server.Use(cors.New())
	return &FiberServer{
		Server: server,
	}
}

func (f *FiberServer) Start(port string) error {
	return f.Server.Listen(fmt.Sprintf(":%v", port))
}
