package server

import (
	"fmt"

	"github.com/Hoi-Kun/letsgo-bongu.com/middleware"
	"github.com/Hoi-Kun/letsgo-bongu.com/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Port uint32

func (p Port) String() string { return fmt.Sprintf(":%d", p) }

type Server struct {
	port Port
	app  *fiber.App
}

func config() fiber.Config {
	appName := "Letsgo-Bongu"
	return fiber.Config{
		AppName:      appName,
		ServerHeader: appName,
		Views:        engine(),
	}
}

func New(portNumber uint32) *Server {
	s := &Server{}
	s.port = Port(portNumber)
	s.app = fiber.New(config())
	return s
}

func (s *Server) set() {
	s.app.Static("/static", "./static")
}

func (s *Server) middlewares() {
	s.app.Use(compress.New(), middleware.BindSiteConfig, middleware.BindMenu)
}

func (s *Server) routes() {
	route.HandleIndex(s.app.Group("/"))
}

func (s *Server) Serve() error {
	s.set()
	s.middlewares()
	s.routes()
	return s.app.Listen(s.port.String())
}
