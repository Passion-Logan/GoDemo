package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// 登录接口
	logInUri := "https://uaa-zxj-prod.cdcdn.cn:22668/backend/auth/oauth/token"
	
	formValues := url.Values{}
	formValues.Set("test", "a")
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)

	req, err := http.NewRequest("POST", logInUri, formBytesReader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Print(err)
		}
	}(req.Body)

}
