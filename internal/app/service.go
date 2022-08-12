package app

import (
	"encryptionsystem/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func StartServer() {
	var (
		cfg storage.Config
		s   storage.Server
	)

	// _______start INIT_______
	cfg.Init()
	s.Init(cfg, fiber.New())
	s.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	s.App.Use(recover.New(recover.Config{EnableStackTrace: true}))
	s.App.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))
	// _______end INIT_______

	// ____start HANDLERS____
	s.App.Post("/encrypt", s.Encrypt)
	s.App.Post("/decrypt", s.Decrypt)
	// ____end HANDLERS____

	log.Fatal(s.App.Listen(s.Cfg.Address))
}
