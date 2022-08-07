package entity

type FileRecord struct {
	Id            int64
	FileName      string
	MajorVersion  int64
	MinorVersion  int64
	ProductName   string
	AbolishFlag   string `orm:"null;default(0)"`
	PublishFlag   string `orm:"null;default(0)"`
	Operationtype string
	CreateTime    string `orm:"null"`
	Creater       string `orm:"null"`
	ModifyTime    string `orm:"null"`
	Modifier      string `orm:"null"`
}
type FileInfo struct {
	Id           int64
	FileName     string
	MajorVersion int64
	MinorVersion int64
	ProductName  string
	AbolishFlag  string `orm:"default(0)"`
	PublishFlag  string `orm:"default(0)"`
	EditFlag     string `orm:"default(0)"`
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
type UserInfo struct {
	Id         int64
	Userid     string `orm:"unique"`
	Password   string
	Username   string
	UserRole   string
	CreateTime string `orm:"null"`
	Creater    string `orm:"null"`
	ModifyTime string `orm:"null"`
	Modifier   string `orm:"null"`
}
