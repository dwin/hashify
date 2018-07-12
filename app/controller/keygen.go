package controller

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type KeyResp struct {
	KeyHex string
	Length int
}

func KeyGen(c echo.Context) error {
	lenStr := c.Param("length")
	// Parse Int from URI Param
	len, err := strconv.Atoi(lenStr)
	if err != nil {
		log.Println("KeyGen error parsing int from parameter, error: ", err)
		return err
	}
	if len > 256 {
		log.Println("Key length request over limit of 256, requested: ", len)
		return fmt.Errorf("Key length request over limit of 256, requested: %v", len)
	}
	// Generate Random Key
	k, err := randKey(len)
	if err != nil {
		log.Println("KeyGen error unable to generate randKey, error: ", err)
		return err
	}
	j, err := json.Marshal(&KeyResp{
		KeyHex: hex.EncodeToString(k),
		Length: len,
	})
	return c.JSONBlob(http.StatusOK, j)
}

func randKey(len int) (hexVal []byte, err error) {
	b := make([]byte, len)
	_, err = rand.Read(b)
	if err != nil {
		log.Printf("randKey - rand.Read() error: %s\n", err)
		return
	}
	return b, err
}
