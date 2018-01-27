package action

import (
	"fmt"
	httpclient "github.com/Xu-Rui/meican-robot/util"
	"net/http"
)

func Login(username string, password string) (string, error) {

	request := httpclient.GetRequest(loginURL, "POST")

	//设置请求头
	request.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	request.Set("Origin", "https://meican.com")
	request.Set("Host", "meican.com")
	request.Set("Referer", "https://meican.com/login")
	request.Set("Cache-Control", "max-age=0")
	request.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Set("Accept-Encoding", "gzip, deflate, br")
	request.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6")

	machineId := http.Cookie{
		Name:"machineId",
		Value:"49e21695-5851-4599-a1e1-75b72d909dc3",
	}

	guestId:= http.Cookie{
		Name:"guestId",
		Value:"806721ae-962b-4cde-b759-392c8a618ff4",
	}

	request.AddCookie(&machineId)
	request.AddCookie(&guestId)

	paramMap := map[string]string{
		"username":    username,
		"password":    password,
		"loginType":   "username",
		"remember":    "true",
		"redirectUrl": "",
		"openId":      "",
	}

	resp, _, errs := request.Type("form").Send(paramMap).EndBytes()
	if errs != nil {
		fmt.Println(errs[0])
	}

	fmt.Println(resp.Header["Set-Cookie"])
	fmt.Println(resp.Request)
	fmt.Println(resp.StatusCode)

	return "", nil
}

func GetBuildList(sessionID string) error {
	request := httpclient.GetRequest(buildList, "GET")

	//设置cookies
	remember := http.Cookie{
		Name:  "remember",
		Value: sessionID,
	}
	request.AddCookie(&remember)

	_, body, errs := request.EndBytes()
	if errs != nil {
		return errs[0]
	}

	fmt.Println(string(body))
	return nil
}
