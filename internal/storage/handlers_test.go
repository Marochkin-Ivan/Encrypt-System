package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Encrypt(t *testing.T) {
	tests := []struct {
		description  string
		expectedCode int
		request      string
		route        string
	}{
		{
			description:  "success (secretMessage)",
			expectedCode: http.StatusOK,
			request:      "{\"message\":\"secretMessage\"}",
			route:        "/encrypt",
		},
		{
			description:  "success(Hello!!!)",
			expectedCode: http.StatusOK,
			request:      "{\"message\":\"Hello!!!\"}",
			route:        "/encrypt",
		},
		{
			description:  "empty message",
			expectedCode: http.StatusBadRequest,
			request:      "{\"message\":\"\"}",
			route:        "/encrypt",
		},
		{
			description:  "bad json",
			expectedCode: http.StatusBadRequest,
			request:      "{\"message\":\"\"",
			route:        "/encrypt",
		},
	}
	var s Server
	s.Init(Config{}, fiber.New())
	s.App.Post("/encrypt", s.Encrypt)

	for _, test := range tests {
		body := bytes.NewBuffer([]byte(test.request))
		req := httptest.NewRequest(http.MethodPost, test.route, body)

		resp, err := s.App.Test(req, -1)
		if err != nil {
			log.Println(err)
			continue
		}
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		if test.expectedCode == http.StatusOK {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				continue
			}
			var res EncryptRequest
			_ = json.Unmarshal([]byte(test.request), &res)
			assert.Equalf(t, res.Message, s.DecryptFunc(string(respBody)), test.description)
		}
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestServer_Decrypt(t *testing.T) {
	var s Server
	s.Init(Config{}, fiber.New())
	s.App.Post("/decrypt", s.Decrypt)
	tests := []struct {
		description    string
		expectedCode   int
		request        string
		expectedResult string
		route          string
	}{
		{
			description:    "success1",
			expectedCode:   http.StatusOK,
			expectedResult: "StrangerThings",
			request:        fmt.Sprintf("{\"message\":\"%s\"}", s.EncryptFunc("StrangerThings")),
			route:          "/decrypt",
		},
		{
			description:    "success2",
			expectedCode:   http.StatusOK,
			expectedResult: "hello???",
			request:        fmt.Sprintf("{\"message\":\"%s\"}", s.EncryptFunc("hello???")),
			route:          "/decrypt",
		},
		{
			description:  "empty message",
			expectedCode: http.StatusBadRequest,
			request:      "{\"message\":\"\"}",
			route:        "/decrypt",
		},
		{
			description:  "bad json",
			expectedCode: http.StatusBadRequest,
			request:      "{\"message\":\"\"",
			route:        "/decrypt",
		},
	}

	for _, test := range tests {
		body := bytes.NewBuffer([]byte(test.request))
		req := httptest.NewRequest(http.MethodPost, test.route, body)

		resp, err := s.App.Test(req, -1)
		if err != nil {
			log.Println(err)
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		if test.expectedCode == http.StatusOK {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				continue
			}
			assert.Equalf(t, test.expectedResult, string(respBody), test.description)
		}
		resp.Body.Close()
	}
}
