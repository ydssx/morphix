package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// 定义一个加密函数，用于对数据进行加密
func Encrypt(data []byte, key []byte) ([]byte, error) {
	// 创建一个加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 检查密钥长度
	if len(key) != block.BlockSize() {
		return nil, fmt.Errorf("invalid key size, expected %d bytes but got %d", block.BlockSize(), len(key))
	}
	// 对数据进行填充
	data = padding(data, block.BlockSize())

	// 创建一个CBC模式的加密器
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCBCEncrypter(block, iv)

	// 对数据进行加密
	encrypted := make([]byte, len(data))
	stream.CryptBlocks(encrypted, data)

	return encrypted, nil
}

// 定义一个解密函数，用于对加密数据进行解密
func Decrypt(encrypted []byte, key []byte) ([]byte, error) {
	// 创建一个加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建一个CBC模式的加密器
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCBCDecrypter(block, iv)

	// 对加密数据进行解密
	decrypted := make([]byte, len(encrypted))
	stream.CryptBlocks(decrypted, encrypted)

	// 去除填充的数据
	decrypted = unpadding(decrypted)

	return decrypted, nil
}

// 对数据进行填充
func padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padData...)
}

// 去除填充的数据
func unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
