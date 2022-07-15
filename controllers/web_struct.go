package controllers

type FileInfo struct {
	FlieName    string `json:"flieName"`
	MajorUpdate string `json:"majorUpdate"`
	ProductName string `json:"productName"`
	Creater     string `json:"creater"`
}
type UserInfo struct {
	Userid   string `json:"userid"`
	Password string `json:"password"`
	Username string `json:"username"`
	UserRole string `json:"userRole"`
}
