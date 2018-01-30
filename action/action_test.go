package action

import (
	"fmt"
	"github.com/Xu-Rui/meican-robot/util"
	"testing"
)

func Test_Login(t *testing.T) {
	Login("testdata", "testdata")

}

func Test_GetBuildList(t *testing.T) {
	GetBuildList("testdata")
}

func Test_Whole_Process(t *testing.T) {

	user, err := util.InitConfig("../config.txt")

	sessionID, err := Login(user.Username, user.Password)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	buildList, err := GetBuildList(sessionID)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fmt.Println(buildList)
}
