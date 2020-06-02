package vision

import "errors"

type GroupMap struct {
	COMMUNITY string `json:"community"` // 智慧社区的小区id
	GROUP     int    `json:"group"`     // PMS的小区id
}

//如果不存在映射，则说明该社区没有开通对应服务
var OPEN_PLATFORM_GROUP_MAP = []GroupMap{
	{
		"oudianhuayuan",
		5,
	},
}

func GetGroupMap(groupID int) (error, GroupMap) {
	for _, gmap := range OPEN_PLATFORM_GROUP_MAP {
		if gmap.GROUP == groupID {
			return nil, gmap
		}
	}
	return errors.New("组团映射不存在"), GroupMap{}
}
