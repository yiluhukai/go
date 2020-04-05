package model

import "time"

type Record struct {
	Id         int       `json:"id" db:"id"`
	Add        int       `json:"add" db:"add"`
	Mark       int       `json:"mark" db:"mark"`
	OpenId     string    `json:"openId" db:"openId"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
	Note       string    `db:"note" json:"note"`
}
