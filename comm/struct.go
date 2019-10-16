package comm

import "time"

//User 对应user表
type User struct {
	ID              int    `orm:"column(id);pk;auto"`
	UserID          uint64 `orm:"column(user_id)"`
	Account         string
	Password        string
	Platform        uint32
	Imei            string
	DeviceID        string `orm:"column(device_id)"`
	Mac             string
	Deleted         uint32
	UserType        uint32
	AutoAddFriend   uint32
	AppSource       uint32    `orm:"column(app_source)"`        // app来源字段, 0:feige  1:新app
	AppSrouceRemark string    `orm:"column(app_source_remark)"` // app来源 说明0:飞鸽  1:闲博
	EncryptKey      string    `orm:"column(encrypt_key)"`       // 客户端加密的key
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
}


type StResp struct {
	Ret uint32 `json:"ret"`
	ErrMsg string `json:"err_msg"`
	UserList []uint64 `json:"user_list"`
}