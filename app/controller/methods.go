package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type HashMethod struct {
	Name         string
	Endpoint     string
	MinKeyLength int
	MaxKeyLength int
}

func ListMethods(c echo.Context) error {
	list := []HashMethod{
		{Name: "Blake2b-256", Endpoint: "/hash/BLAKE2B-256", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "Blake2b-384", Endpoint: "/hash/BLAKE2B-384", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "Blake2b-512", Endpoint: "/hash/BLAKE2B-512", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "Blake2s-128", Endpoint: "/hash/BLAKE2s-128", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "Blake2s-256", Endpoint: "/hash/BLAKE2s-256", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "HighwayHash-256", Endpoint: "/hash/HIGHWAY", MinKeyLength: 32, MaxKeyLength: 32},
		{Name: "HighwayHash-64", Endpoint: "/hash/HIGHWAY-64", MinKeyLength: 32, MaxKeyLength: 32},
		{Name: "HighwayHash-128", Endpoint: "/hash/HIGHWAY-128", MinKeyLength: 32, MaxKeyLength: 32},
		{Name: "MD4", Endpoint: "/hash/MD4", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "MD5", Endpoint: "/hash/MD5", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA1", Endpoint: "/hash/SHA1", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA256", Endpoint: "/hash/SHA256", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA384", Endpoint: "/hash/SHA384", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA512", Endpoint: "/hash/SHA512", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA512-256", Endpoint: "/hash/SHA512-256", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA3-256", Endpoint: "/hash/SHA3-256", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA3-384", Endpoint: "/hash/SHA3-384", MinKeyLength: 0, MaxKeyLength: 0},
		{Name: "SHA3-512", Endpoint: "/hash/SHA3-512", MinKeyLength: 0, MaxKeyLength: 0},
	}

	j, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
