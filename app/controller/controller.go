package controller

import jsoniter "github.com/json-iterator/go"

type BasicError struct {
	Error string
}

type HashResp struct {
	Digest    string
	DigestEnc string
	Type      string
	Key       string
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary
