package controller

import jsoniter "github.com/json-iterator/go"

type HashResp struct {
	Digest string
	Type   string
	Key    string
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary
