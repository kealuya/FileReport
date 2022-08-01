package controllers

type FileInfo struct {
	FlieName    string `json:"flieName"`
	MajorUpdate string `json:"majorUpdate"`
	Version     string `json:"version"`
	ProductName string `json:"productName"`
	Userid      string `json:"userid"`
}
type UserInfo struct {
	Userid   string `json:"userid"`
	Password string `json:"password"`
	Username string `json:"username"`
	UserRole string `json:"userRole"`
	Modifier string `json:"modifier"`
}
