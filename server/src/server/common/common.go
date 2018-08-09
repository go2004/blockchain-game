package common

import (
	"time"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

//获取当前时间戳
func GetTimestamp() int64 {
	t := time.Now()
	return t.Unix()

}


// 生成32位MD5
func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(strlen int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < strlen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}