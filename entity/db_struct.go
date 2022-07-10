package entity

type FileInfo struct {
	Id           int64 `orm:"null"`
	FileName     string
	MajorVersion int8   `orm:"null"`
	MinorVersion int8   `orm:"null"`
	ProductName  string `orm:"null"`
	CreateTime   string `orm:"null"`
	Creater      string `orm:"null"`
	ModifyTime   string `orm:"null"`
	Modifier     string `orm:"null"`
}
type ProductStatus struct {
	ProductName  string //项目名称
	FileNumber   int32  //文件数量
	PersonNumber int32  //维护人数
	VersionCount int32  //发布次数
}

type ProductInfo struct {
	Id             int64
	ProductName    string
	LastUpdateTime string
	LastCreater    string
}
