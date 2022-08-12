package storage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_EncryptDecryptFunc(t *testing.T) {
	tests := []struct {
		description string
		message     string
	}{
		{
			description: "encrypt works correct",
			message:     "secondSecretMessage",
		},
		{
			description: "encrypt difficult word works correct",
			message:     "second Secret SUPER (%^ Message.....",
		},
	}
	var s Server
	s.Init(Config{}, fiber.New())

	for _, test := range tests {
		encryptMessage := s.EncryptFunc(test.message)
		decryptMessage := s.DecryptFunc(encryptMessage)

		assert.Equalf(t, test.message, decryptMessage, test.description)
	}
}
