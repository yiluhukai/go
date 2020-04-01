package model

type Opinion struct {
	Id      int    `json:"id" db:"id"`
	Opinion string `json:"opinion" db:"opinion"`
	Src     string `json:"src" db:"src"`
	WeChat  string `json:"wechat" db:"wechat"`
	OpenId  string `json:"openId" db:"openId"`
}
