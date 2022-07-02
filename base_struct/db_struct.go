package base_struct

type FileInfo struct {
	Id           int64
	FlieName     string
	MajorVersion int8
	MinorVersion int8
	ProductName  string
	CreateTime   string
	Creater      string
	ModifyTime   string
	Modifier     string
}
type ProductInfo struct {
	ProductName  string
	FileNumber   int32
	PersonNumber int32
	VersionCount int32
}
