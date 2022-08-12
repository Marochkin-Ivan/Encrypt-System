package storage

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (s *Server) Encrypt(c *fiber.Ctx) error {
	var r EncryptRequest
	err := json.Unmarshal(c.Body(), &r)
	if err != nil || r.Message == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.Status(http.StatusOK).SendString(s.EncryptFunc(r.Message))
}

func (s *Server) Decrypt(c *fiber.Ctx) error {
	var r DecryptRequest
	err := json.Unmarshal(c.Body(), &r)
	if err != nil || r.CodedMessage == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	if decryptMessage := s.DecryptFunc(r.CodedMessage); decryptMessage == "" {
		return c.SendStatus(http.StatusConflict)
	} else {
		return c.Status(http.StatusOK).SendString(decryptMessage)
	}
}
