package util

import (
	"bufio"
	"github.com/Xu-Rui/meican-robot/model"
	validator "gopkg.in/go-playground/validator.v9"
	"io"
	"os"
	"strings"
)

// Validate 用来对 request struct 进行 validate
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

//初始化配置文件
func InitConfig(filePath string) (*model.UserConfig, error) {

	var user model.UserConfig
	scanState := 0

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		lineBuffer, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lineStr := string(lineBuffer)

		//匹配分割线
		if strings.Contains(lineStr, "---") {
			scanState++
			continue
		}

		//匹配到账号
		if strings.HasPrefix(lineStr, "账号") {
			tempArray := strings.Split(lineStr, "[")
			tempArray = strings.Split(tempArray[1], "]")
			user.Password = strings.Trim(tempArray[0], " ")
			user.Username = tempArray[0]
			continue
		}

		//匹配到密码
		if strings.HasPrefix(lineStr, "密码") {
			tempArray := strings.Split(lineStr, "[")
			tempArray = strings.Split(tempArray[1], "]")
			user.Password = strings.Trim(tempArray[0], " ")
			continue
		}

		//匹配办公地点
		if strings.Contains(lineStr, "@") && scanState == 2 {
			tempStr := strings.TrimRight(lineStr, "@")
			user.Workspace = strings.Trim(tempStr, " ")
			continue
		}

		//匹配选项
		if strings.Contains(lineStr, "@") && scanState == 3 {
			//TODO 一期只开发随机订餐
			user.IsRandom = true
			continue
		}
	}

	//校验用户参数
	if err := Validate.Struct(user); err != nil {
		return nil, err
	}

	return &user, nil
}
