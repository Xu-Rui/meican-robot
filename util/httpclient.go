package util

import (
	"github.com/parnurzeal/gorequest"
)

// Get 封装GET方法
func Get(urlstr string, params map[string]string) ([]byte, error) {

	request := GetRequest(urlstr, "GET")

	for key, value := range params {
		request = request.Param(key, value)
	}

	_, body, errs := request.EndBytes()
	if errs != nil {
		return nil, errs[0]
	}

	return body, nil
}

// PostForm 封装POST方法
func PostForm(urlstr string, params map[string]string) ([]byte, error) {

	request := GetRequest(urlstr, "POST")

	_, body, errs := request.Type("form").Send(params).EndBytes()
	if errs != nil {
		return nil, errs[0]
	}

	return body, nil
}

//PostJSON 封装POST方法
func PostJSON(urlstr string, data map[string]interface{}) ([]byte, error) {

	request := GetRequest(urlstr, "POST")

	request.Set("Content-Type", "application/json")
	_, body, _ := request.Type("json").Send(data).EndBytes()

	return body, nil
}

func GetRequest(urlstr string, method string) (*gorequest.SuperAgent) {

	var request *gorequest.SuperAgent
	switch method {
	case "GET":
		request = gorequest.New().Get("http://" + urlstr)
	case "POST":
		request = gorequest.New().Post("http://" + urlstr)
	}

	return request
}
