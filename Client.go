package goCoinGate

import (
	"github.com/parnurzeal/gorequest"
	"fmt"
)

type CoinGateClient struct {
	apiKey string
	secret string
}

func Connect(key string, secret string) (int, bool) {
	return  1, false
}

func CreateOrder() bool {

	const createUrl = "https://api-sandbox.coingate.com/v1/orders"

	request := gorequest.New()
	resp, _, errs := request.Post(createUrl).
		Set("Notes","gorequst is coming!").
		Send(`{"name":"backy", "species":"dog"}`).
		End()

	if (errs != nil) {
		fmt.Print("ok to proceed")
	}

	fmt.Println(resp.Status)

	//resp.Body.Read()

	return true;
}
