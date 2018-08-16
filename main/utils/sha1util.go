package utils

import (
	"crypto/sha1"
	"encoding/base32"
	"strconv"
	"time"
)

func Hash(msg string, len int) (string) {
	//SHA1
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(msg + strconv.FormatInt(time.Now().Unix(),10)) )
	return base32.HexEncoding.EncodeToString(Sha1Inst.Sum([]byte("")))[:len]
}
