package utils

import (
	"AuthTemplate/src"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
)

func EncryptAES(data map[string]interface{}) (string, error) {
	block, err := aes.NewCipher(src.Config.EncryptionKey)
	if err != nil {
		return "", err
	}

	dataBytes, _ := json.Marshal(data)

	ciphertext := make([]byte, aes.BlockSize+len(dataBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(dataBytes))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES(encodedCipher string) (map[string]interface{}, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encodedCipher)
	if err != nil {
		return map[string]interface{}{}, err
	}

	block, err := aes.NewCipher(src.Config.EncryptionKey)
	if err != nil {
		return map[string]interface{}{}, err
	}

	if len(ciphertext) < aes.BlockSize {
		return map[string]interface{}{}, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	var dataMap map[string]interface{}
	_ = json.Unmarshal(ciphertext, &dataMap)
	return dataMap, nil
}
