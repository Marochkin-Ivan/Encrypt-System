package storage

import "github.com/gofiber/fiber/v2"

type Server struct {
	App      *fiber.App
	Cfg      Config
	keys     []string
	key4keys string
}

type Config struct {
	Address string
}

type EncryptRequest struct {
	Message string `json:"message"`
}

type DecryptRequest struct {
	CodedMessage string `json:"message"`
}
