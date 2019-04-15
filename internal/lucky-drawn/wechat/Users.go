package wechat

type Users struct {
	Lang         string  `json:"lang"`
	GroupList    []Group `json:"group_list"`
	UserList     []User  `json:"user_list"`
	TotalUserNum string  `json:"total_user_num"`
}

type Group struct {
	Name       string `json:"name"`
	Cnt        string `json:"cnt"`
	CreateTime string `json:"create_time"`
	Id         string `json:"id"`
}

type User struct {
	Id         string   `json:"id"`
	NickName   string   `json:"nick_name"`
	RemarkName string   `json:"remark_name"`
	CreateTime string   `json:"create_time"`
	GroupId    []string `json:"group_id"`
}
