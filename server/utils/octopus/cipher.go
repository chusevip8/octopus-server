package octopus

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)

var (
	aesKey = []byte("0123456789OCTOPU") // 16字节密钥
	aesIv  = []byte("1234567890ABCDEF") // 16字节IV偏移量
)

func DataIn(inData []byte) (rawData []byte, err error) {
	data, err := decompress(inData)
	if err != nil {
		return nil, err
	}
	rawData, err = decrypt(aesKey, aesIv, data)
	return
}

func DataOut(rawData []byte) (outData []byte, err error) {
	data, err := encrypt(aesKey, aesIv, rawData)
	if err != nil {
		return nil, err
	}
	outData, err = compress(data)
	return
}
func encrypt(key, iv, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddedPlaintext := pkcs7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedPlaintext)
	return ciphertext, nil
}
func decrypt(key, iv, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decryptedData := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decryptedData, ciphertext)
	return pkcs7Unpadding(decryptedData), nil
}

// 使用PKCS7填充方式对数据进行填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 对使用PKCS7填充方式的数据进行去填充
func pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Printf("Failed to close gzip writer: %s\n", err)
		}
	}()

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := r.Close(); err != nil {
			fmt.Printf("Failed to close gzip reader: %s\n", err)
		}
	}()

	return io.ReadAll(r)
}
