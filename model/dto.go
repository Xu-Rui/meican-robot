package model

//"groups":[]

// BuildingList 办公地点列表
type BuildingList struct {
	Groups []Building `json:"groups"`
	Corps  []Corp     `json:"corps"`
}

//"title":"xx大厦",
//"skip":false,
//"corps":[]

// Building 办公地点
type Building struct {
	Title string `json:"title"`
	Skip  bool   `json:"skip"`
	Corps []Corp `json:"corps"`
}

//"corpId":1234,
//"displayName":"",
//"name":"企业付xx楼",
//"namespace":"123456"

// Corp 企业付款 or 自己付款
type Corp struct {
	CorpID      int    `json:"corpId"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	NameSpace   string `json:"namespace"`
}
