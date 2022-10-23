package entity

import (
	"time"
)

/*
https://gitea.com/xorm/reverse
一个灵活高效的数据库反转工具。
从数据库生成结构体工具
找发布版本，编译，运行
*/

type Doc struct {
	DocId       int    `xorm:"not null pk autoincr INTEGER" json:"docId,omitempty"`
	DocName     string `xorm:"TEXT" json:"docName,omitempty"`
	OwnerId     string `xorm:"TEXT" json:"ownerId,omitempty"`
	CreateDate  string `xorm:"DATE" json:"createDate"`
	DocType     string `xorm:"TEXT" json:"docType,omitempty"`
	IsRelease   string `xorm:"TEXT" json:"isRelease,omitempty"`
	IsOwnerEdit string `xorm:"TEXT" json:"isOwnerEdit,omitempty"`
	ProId       int    `xorm:"INTEGER" json:"proId,omitempty"`
	IsDiscard   string `xorm:"TEXT" json:"isDiscard,omitempty"`
}

type File struct {
	DocId         int    `xorm:"not null INTEGER" json:"docId,omitempty"`
	Version       int    `xorm:"not null integer" json:"version,omitempty"`
	VersionShow   string `xorm:"TEXT" json:"versionShow,omitempty"`
	UpdateDate    string `xorm:"DATE" json:"updateDate"`
	UpdateUserId  string `xorm:"text" json:"updateUserId,omitempty"`
	UpdateContent string `xorm:"text" json:"updateContent,omitempty"`
	FileName      string `xorm:"TEXT" json:"fileName,omitempty"`
}

type Project struct {
	ProId   int    `xorm:"not null pk autoincr INTEGER" json:"proId,omitempty"`
	ProName string `xorm:"TEXT" json:"proName,omitempty"`
	ProLogo string `xorm:"TEXT" json:"proLogo,omitempty"`
}

type User struct {
	PhoneNumber string    `xorm:"not null text" json:"phoneNumber,omitempty"`
	UserName    string    `xorm:"TEXT" json:"userName,omitempty"`
	Token       string    `xorm:"TEXT" json:"token,omitempty"`
	Expire      time.Time `xorm:"DATE" json:"expire"`
	Unique      string    `xorm:"TEXT" json:"unique,omitempty"`
	IsDisable   string    `xorm:"TEXT" json:"isDisable,omitempty"`
}
type Record struct {
	RecId      int    `xorm:"not null  pk autoincr INTEGER" json:"recId"`
	Token      string `xorm:"text" json:"token"`
	RouteName  string `xorm:"TEXT" json:"routeName"`
	CreateTime string `xorm:"date" json:"createTime"`
}
