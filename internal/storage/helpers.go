package storage

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const firstReadableSymbol byte = 35
const maxStringLen int = 25

func XorStrings(str1, str2 string) string {
	buf := XorBytes([]byte(str1), []byte(str2), len(str2))
	return string(buf)
}

func XorBytes(str1, str2 []byte, keyLen int) []byte {
	n := len(str1)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = (str1[i] ^ str2[i%keyLen]) + firstReadableSymbol
	}
	return buf
}

func ReverseXorStrings(str1, str2 string) string {
	buf := ReverseXorBytes([]byte(str1), []byte(str2), len(str2))
	return string(buf)
}

func ReverseXorBytes(str1, str2 []byte, keyLen int) []byte {
	n := len(str1)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = (str1[i] - firstReadableSymbol) ^ str2[i%keyLen]
	}
	return buf
}

func GetKeyIdx(arraySize int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(arraySize)
}

func CheckEncryptMessage(message string, key4keys string) (int, error) {
	keyIdx, err := strconv.Atoi(ReverseXorStrings(message[0:1], key4keys))
	if err != nil || keyIdx > 9 {
		return -1, errors.New("invalid data")
	}
	return keyIdx, nil
}

func RandomStringGen() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#$%^&*()_+{}[]|/`")

	rand.Seed(time.Now().Unix())
	l := rand.Intn(maxStringLen)

	randStr := make([]rune, l)
	for i := range randStr {
		randStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(randStr)
}

func CheckLastIdx(message string) (int, error) {
	if lastIdx := strings.Index(message, "!"); lastIdx == -1 {
		return lastIdx, errors.New("invalid data")
	} else {
		return lastIdx, nil
	}
}
