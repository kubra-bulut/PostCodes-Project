package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {

	url := "https://sandbox-api.iyzipay.com/payment/test"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
