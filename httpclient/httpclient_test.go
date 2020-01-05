package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	fmt.Println("http client test")

	params := map[string]string {
		"aaa": "111",
		"bbb": "222",
	}
	url := "http://www.baidu.com"
	res, err := resty.New().R().SetQueryParams(params).Get(url)
	fmt.Println("res:", res)
	fmt.Println("err:", err)

	http.Client.Do(nil)
}
