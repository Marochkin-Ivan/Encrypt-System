package storage

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (s *Server) Init(c Config, a *fiber.App) {
	s.Cfg = c
	s.App = a
	s.keys = []string{
		"kjdnfgiuwerntdfkvndf",
		"dlkfnvknsdkfnvkjdsfkvnkjdnfveruh",
		"ppweplrk",
		"qqnrkjhvlnrjtjohwnjtrng",
		"kjneiuguergionaefnvkadmfvnadngengniuadnfnsdklfnglksdnfhsdfhj",
		"owojtnnjcviahbbrbngnoibsnfgh",
		"cbbtyjkhjagjbknzxcnm",
		"jkfliguqerhtibqerkjnvanvjkniufdnglwnertkjgnjkdnbfgsdhfj",
		"kjehrhhhfhfhfhhjnekjrbhahjyeqkgerjodbmnjkbhgvrbfkovdfjnefkemwnjerbslhjnavdsjktrbnerfewrevb",
		"kek",
	}
	s.key4keys = "savagerandom"
}

func (c *Config) Init() {
	c.Address = "localhost:8080"
}

func (s *Server) EncryptFunc(message string) string {
	keyIdx := GetKeyIdx(len(s.keys))
	return XorStrings(strconv.Itoa(keyIdx), s.key4keys) +
		XorStrings(message, s.keys[keyIdx]) +
		"!" +
		RandomStringGen()
}

func (s *Server) DecryptFunc(encryptMessage string) string {
	keyIdx, err := CheckEncryptMessage(encryptMessage, s.key4keys)
	if err != nil {
		log.Println(err)
		return ""
	}

	lastIdx, err := CheckLastIdx(encryptMessage)
	if err != nil {
		log.Println(err)
		return ""
	}

	return ReverseXorStrings(encryptMessage[1:lastIdx], s.keys[keyIdx])
}
