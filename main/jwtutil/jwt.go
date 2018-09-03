package jwtutil

import (
	"encoding/json"
	"encoding/base64"
	"crypto/sha256"
	"crypto/hmac"
	"encoding/hex"
	"strings"
	"github.com/sunlggggg/piconline/main/config/jwtconfig"
	"time"
)

// SALT 密钥

// Header 消息头部
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// PayLoad 负载
type PayLoad struct {
	ExpireTime int16  `json:"expire_time"`
	CreateTime int64  `json:"create_time"`
	Name       string `json:"name"`
	UserID     uint64 `json:"user_id"`
}

// JWT 完整的本体
type JWT struct {
	Header           `json:"header"`
	PayLoad          `json:"payload"`
	Signature string `json:"signature"`
}

// Encode 将 json 转成符合 JWT 标准的字符串
func (jwt *JWT) Encode() (string, error) {
	header, err := json.Marshal(jwt.Header)
	if err != nil {
		return "", err
	}
	headerString := base64.StdEncoding.EncodeToString(header)
	payload, err := json.Marshal(jwt.PayLoad)
	payloadString := base64.StdEncoding.EncodeToString(payload)
	if err != nil {
		return "", err
	}

	format := headerString + "." + payloadString
	signature := getHmacCode(format)

	return format + "." + signature, nil
}

func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte(jwtconfig.SALT))
	h.Write([]byte(s))
	key := h.Sum(nil)
	return hex.EncodeToString(key)
}

// Decode 验证 jwt 签名是否正确,并将json内容解析出来
func (jwt *JWT) Decode(code string) (bool) {

	arr := strings.Split(code, ".")
	if len(arr) != 3 {
		return false
	}

	// 验证签名是否正确
	format := arr[0] + "." + arr[1]
	signature := getHmacCode(format)
	if signature != arr[2] {
		return false
	}

	header, err := base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		return false
	}
	payload, err := base64.StdEncoding.DecodeString(arr[1])

	if err != nil {
		return false
	}

	json.Unmarshal(header, &jwt.Header)
	json.Unmarshal(payload, &jwt.PayLoad)
	timeInterval := time.Now().Unix() - jwt.PayLoad.CreateTime

	if timeInterval > int64(jwt.PayLoad.ExpireTime) {
		return false
	}
	return true
}