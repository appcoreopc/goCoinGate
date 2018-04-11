package goCoinGate

import (
	"github.com/parnurzeal/gorequest"
	)

type CoinGateClient struct {
	apiKey string
	secret string
}

func Connet(key string, secret string) (int, bool) {
	return  1, false
}

func CreateOrder() bool {
	return true;
}
