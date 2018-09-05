package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"happylemon/conf"
	"net/url"
)

//解密函数
func Decryper(body []byte) (bool, []byte) {

	okbody := string(body)
	decodeData, err1 := UrlDecode(okbody)
	if !err1 {
		return false, []byte{}
	}
	base64Data, err2 := Base64Decode(decodeData)
	if !err2 {
		return false, []byte{}
	}

	key := conf.Config.Server.AeskeyBytes
	re, err3 := CbcPkcs5Decryper(base64Data, key, make([]byte, 16))
	if err3 != nil {
		return false, []byte{}
	}
	return true, re
}

//加密函数
func Encryper(body []byte) (bool, string) {
	key := conf.Config.Server.AeskeyBytes
	ww, err1 := CbcPkcs5Encryper(body, key, make([]byte, 16))
	if err1 != nil {
		return false, ""
	}
	re := UrlEncode(Base64Encode(ww))
	return true, re

}

// CbcPkcs5Encryper 加密 CBC/PKCS5
func CbcPkcs5Encryper(source []byte, key []byte, iv []byte) ([]byte, error) {
	sourceBlock, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	source = PKCS5Padding(source, sourceBlock.BlockSize()) //补全位数，长度必须是 16 的倍数
	sourceCrypted := make([]byte, len(source))
	sourceAes := cipher.NewCBCEncrypter(sourceBlock, iv)
	sourceAes.CryptBlocks(sourceCrypted, source)
	return sourceCrypted, err
}

// CbcPkcs5Decryper 解密 CBC/PKCS5
func CbcPkcs5Decryper(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	var err error
	emptyBytes := []byte{}
	sourceBlock, err := aes.NewCipher(key)
	if err != nil {
		return emptyBytes, err
	}
	if len(crypted)%sourceBlock.BlockSize() != 0 {
		err = errors.New("crypto/cipher: input not full blocks")
		return emptyBytes, err
	}

	source := make([]byte, len(crypted))
	sourceAes := cipher.NewCBCDecrypter(sourceBlock, iv)
	sourceAes.CryptBlocks(source, crypted)
	source = PKCS5UnPadding(source)
	return source, err
}

// CbcEncryper CBC
func CbcEncryper(source []byte, key []byte, iv []byte) ([]byte, error) {
	sourceBlock, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	sourceCrypted := make([]byte, len(source))
	sourceAes := cipher.NewCBCEncrypter(sourceBlock, iv)
	sourceAes.CryptBlocks(sourceCrypted, source)
	return sourceCrypted, err
}

// PKCS5Padding PKCS5Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding PKCS5UnPadding
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//base64编码
func Base64Encode(data []byte) string {
	base64encodeBytes := base64.StdEncoding.EncodeToString(data)
	return base64encodeBytes
}

//base64解码
func Base64Decode(data string) ([]byte, bool) {
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, false
	}
	return decodeBytes, true
}

//url 编码
func UrlEncode(data string) string {
	encode := url.QueryEscape(data)
	return encode
}

//url 解码
func UrlDecode(data string) (string, bool) {
	decodeurl, err := url.QueryUnescape(data)
	if err != nil {
		return "", false
	}
	return decodeurl, true
}
