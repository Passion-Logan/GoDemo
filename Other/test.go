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
	formValues.Set("username", "13882901701")
	formValues.Set("password", "IfyZgMNvj2YmKXsTjME36uwzPuhdhA2BrSd0DIq+gHnEghBlNJrQW/CVcUySNP2/pM4nvl2ZwyriMBs+ZsaOknS0/HUB5cnYW5bJZIrPfoFRXUxgxF9amCxgd4oHkjlazuHL57X/remeF0nejFqIA5z/TNsTZwVUeEsj83nhr2hNEbj0vI5oUrxfkZWjhQ0Rs4m8scY06hRKqoPeaRCzpHeB2BEzJpoBDs20N4OorrDNOT3NSbqfcYON12xtIbPMVbjx2EEnEj49KY+q4EVL3MCCm0oa8RHpWgLKlYz8Kxdk19soe0uQOnxHNDqom9VfqdIBan16UIUUyHo9xLZRTA==")
	formValues.Set("grant_type", "password")
	formValues.Set("appTypeEnum", "WEB")
	formValues.Set("device", "Win32")
	formValues.Set("browser", "Chrome/98.0.4758.102")
	formValues.Set("address", "四川省成都市")
	formValues.Set("authCode", "null")
	formValues.Set("authKey", "null")
	formDataStr := formValues.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)

	req, err := http.NewRequest("POST", logInUri, formBytesReader)
	req.Header.Set("Authorization", "Basic Y2hhb3MtdWFhOmxyd3p5dlpGUEpvWEN5enc=")
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

	// Basic Y2hhb3MtdWFhOmxyd3p5dlpGUEpvWEN5enc=
	// 加密规则 用户密码->rsa加密->urlEncode
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmruM9GHsYNe1TBvDlBqjQVIzYnjeXdsC9xchkmrtmy3mSnk2SotSZJHGH6a59DE2pKb5kRhk3ozEry6JkJZRN4hDpsafbT8IER6L/gqhyad3ujEgSCW+nyddc4DLdySzoMPG9avuagK3yXbqFM+0RDBJRC43j02CZoIFhFsiHe3Isk2M2Ue08QmjF2A26R1B7DDpE/9lgfn4ZYuSRncWt62NZkwsvWcB8WUIaHvwxpJXYau+5zED2EqD3jVxoP3Q9Y8ar5tB6jrkqKCtvgCgc//g8jXGgfdWWLtsDEED66y1ltmdzLa5VM675HxPuFiH7v39t1pnfayS7sYNXQ65awIDAQAB

}
