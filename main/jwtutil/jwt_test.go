package jwtutil

import (
	"testing"
	"github.com/sunlggggg/piconline/main/config/jwtconfig"
	"time"
)

func Test_Encode(t *testing.T) {
	jwt := JWT{}
	jwt.Header = Header{"HS256", "JWT"}
	jwt.PayLoad = PayLoad{ExpireTime: jwtconfig.ValidTime, Name: "sunlg", CreateTime: time.Now().Unix()}
	result, _ := jwt.Encode()
	t.Log(result)
}

func Test_Decode(t *testing.T) {
	testStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MTAwLCJjcmVhdGVfdGltZSI6MTUzNDIyMjU2MywibmFtZSI6InN1bmxnIn0=.60d926139923e531b709870cb6d9028b67e8247fcb6de3e6ccb45496a47a338c"
	jwt := JWT{}
	status := jwt.Decode(testStr)
	if status {
		t.Log(jwt)
	} else {
		t.Error("error json content")
	}
}
