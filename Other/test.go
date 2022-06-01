package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Print(getToken())
}

const (
 Partake = "https://uaa-zxj-prod.cdcdn.cn:22668/backend/pm/mgr/project/partake"
 TaskList = "https://uaa-zxj-prod.cdcdn.cn:22668/backend/pm/mgr/dailyReport/queryTaskList"
)

// 622959ede4b08db2022070ca
func getToken() map[string]interface{} {
	
	req, err := http.NewRequest("GET", Partake, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer 622959ede4b08db2022070ca")
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Print(err)
		}
	}(req.Body)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	return Transformation(resp)
}


// Transformation 转换接口返回结果 -> map
func Transformation(response *http.Response) map[string]interface{} {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil
		}
	}
	return result
}

