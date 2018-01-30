package action

import (
	"github.com/Xu-Rui/meican-robot/model"
	"github.com/Xu-Rui/meican-robot/util"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strings"
)

//Login 模拟登陆美餐，将登陆获取到的 SessionID 返回
func Login(username string, password string) (string, error) {

	var remember string
	request := util.GetRequest(loginURL, "POST")

	request.RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
		return http.ErrUseLastResponse
	})

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
		return "", util.Errorf("请求失败，请检查网络")
	}

	if len(resp.Header["Set-Cookie"]) > 0 {
		for _, v := range resp.Header["Set-Cookie"] {
			if strings.HasPrefix(v, "remember") {
				remember = strings.Split(v[9:], ";")[0]
				return remember, nil
			}
		}
	}

	return "", util.Errorf("username or password error")
}

//GetBuildList 获取办公地点列表
func GetBuildList(sessionID string) (*model.BuildingList, error) {
	request := util.GetRequest(buildList, "GET")

	//设置cookies
	remember := http.Cookie{
		Name:  "remember",
		Value: sessionID,
	}
	request.AddCookie(&remember)

	_, body, errs := request.EndBytes()
	if errs != nil {
		return nil, errs[0]
	}

	var buildList model.BuildingList

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(body, &buildList)
	if err != nil {
		return nil, err
	}

	return &buildList, nil
}
