package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type HashMethod struct {
	Name     string
	Endpoint string
	Keyed    bool
}

func ListMethods(c echo.Context) error {
	list := []HashMethod{
		{Name: "Blake2b-256", Endpoint: "/hash/BLAKE2B-256", Keyed: false},
		{Name: "Blake2b-384", Endpoint: "/hash/BLAKE2B-384", Keyed: false},
		{Name: "Blake2b-512", Endpoint: "/hash/BLAKE2B-512", Keyed: false},
		{Name: "Blake2s-128", Endpoint: "/hash/BLAKE2s-128", Keyed: false},
		{Name: "Blake2s-256", Endpoint: "/hash/BLAKE2s-256", Keyed: false},
		{Name: "HighwayHash-256", Endpoint: "/hash/HIGHWAY", Keyed: true},
		{Name: "HighwayHash-64", Endpoint: "/hash/HIGHWAY-64", Keyed: true},
		{Name: "HighwayHash-128", Endpoint: "/hash/HIGHWAY-128", Keyed: true},
		{Name: "MD4", Endpoint: "/hash/MD4", Keyed: false},
		{Name: "MD5", Endpoint: "/hash/MD5", Keyed: false},
		{Name: "SHA1", Endpoint: "/hash/SHA1", Keyed: false},
		{Name: "SHA256", Endpoint: "/hash/SHA256", Keyed: false},
		{Name: "SHA384", Endpoint: "/hash/SHA384", Keyed: false},
		{Name: "SHA512", Endpoint: "/hash/SHA512", Keyed: false},
		{Name: "SHA512-256", Endpoint: "/hash/SHA512-256", Keyed: false},
		{Name: "SHA3-256", Endpoint: "/hash/SHA3-256", Keyed: false},
		{Name: "SHA3-384", Endpoint: "/hash/SHA3-384", Keyed: false},
		{Name: "SHA3-512", Endpoint: "/hash/SHA3-512", Keyed: false},
	}

	j, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, j)
}
