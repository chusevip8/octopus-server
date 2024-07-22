package octopus

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

const aesPassword = "9a1d2e3f4b5c6d7e8f9a0b1c2d3e4f5a"

func DataIn(inData []byte) (rawData []byte, err error) {
	data, err := decompress(inData)
	if err != nil {
		return nil, err
	}
	rawData, err = decrypt(data, []byte(aesPassword))
	return
}

func DataOut(rawData []byte) (outData []byte, err error) {
	data, err := encrypt(rawData, []byte(aesPassword))
	if err != nil {
		return nil, err
	}
	outData, err = compress(data)
	return
}

// 加密数据
func encrypt(plaintext, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// 解密数据
func decrypt(ciphertext, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
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
