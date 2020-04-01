package model

type UserInfo struct {
	OpenId    string `json:"openId"`
	NickName  string `json:"nickName" db:"nickName"`
	Gender    int    `json:"gender" db:"gender"`
	City      string `json:"city" db:"city"`
	Province  string `json:"province" db:"province"`
	Country   string `json:"country" db:"country"`
	AvatarUrl string `json:"avatarUrl" db:"avatarUrl"`
	UnionId   string `json:"unionId"`
}

type User struct {
	Id          int      `db:"id" json:"id"`
	UserId      string   `db:"user_id" json:"user_id"`
	Sky         string   `db:"sky" json:"sky"`
	SessionKey  string   `db:"session_key" json:"session_key"`
	CreateTime  string   `db:"create_time" json:"create_time"`
	OpenId      string   `json:"openId" db:"open_id"`
	UserInfoObj UserInfo `json:"userInfo"`
	UserInfo    string   `db:"userInfo"`
}
