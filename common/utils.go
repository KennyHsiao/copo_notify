package common

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//密码加密
func PasswordHash(plainpwd string) string {
	//谷歌的加密包
	hash, err := bcrypt.GenerateFromPassword([]byte(plainpwd), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePWD
}

//密码校验
func CheckPassword(plainpwd, cryptedpwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cryptedpwd), []byte(plainpwd)) //验证（对比）
	return err == nil
}

//時區轉換
func ParseTime(t string) string {
	timeString, _ := time.Parse(time.RFC3339, t)

	return timeString.String()
}

// 微服务密钥解析
func MicroServiceVerification(sing, key, publicKey string) (isOk bool, err error) {
	var singByte []byte
	if singByte, err = base64.StdEncoding.DecodeString(sing); err != nil {
		return
	}
	if singByte, err = DesCBCDecrypt(singByte, []byte(publicKey)); err != nil {
		return
	}

	decryptStr := string(singByte)

	if strings.Index(decryptStr, key) > -1 {
		trimStr := strings.Replace(decryptStr, key, "", 1)
		if timeX, err := time.ParseInLocation("200601021504", trimStr, time.Local); err == nil {
			if time.Now().Sub(timeX).Minutes() <= 10 {
				isOk = true
			}
		}
	}
	return
}

func DesCBCDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	//origData := make([]byte, len(crypted))
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)

	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
