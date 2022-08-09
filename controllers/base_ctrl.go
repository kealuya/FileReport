package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JsonStruct struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
type TokenStruct struct {
	UserInfo  interface{} `json:"user_info"`
	Token     string      `json:"token"`
	ExpiresAt string      `json:"expires_at"`
}

func NewJsonStruct(data interface{}) *JsonStruct {
	return &JsonStruct{
		Success: true,
		Msg:     "0000",
		Data:    data,
	}
}

type Claims struct {
	UserId             string
	jwt.StandardClaims // jwt中标准格式,主要是设置token的过期时间
}

const (
	sign = "fileReport"
)

// GenerateToken() 生成token
// 调用库的NewWithClaims(加密方式,载荷).SignedString(签名) 生成token
func GetnerateToken(user string) (string, string, error) {
	// 当前时间
	nowTime := time.Now()
	// 过期时间
	expireTime := nowTime.Add(30 * time.Minute)
	//   签发人
	issuer := user
	//	 赋值给结构体
	claims := Claims{
		UserId: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 转成纳秒
			Issuer:    issuer,
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==》 头部，载荷，签证
	toke, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(sign))
	exdate := expireTime.Format("2006-01-02 15:04:05")
	return toke, exdate, err
}

//  ParseToken 解析token
// 调用ParseWithClaims进行解析，使用签名解密出数据
func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 使用签名解析用户传入的token,获取载荷部分数据
		return []byte(sign), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if c, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return c, nil
		}
	}
	return nil, err
}
