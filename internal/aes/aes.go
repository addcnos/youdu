package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
)

const PADDING = 32

type RawMsg struct {
	Data   []byte
	Length int32
	AppId  string
}

func Padding(in []byte) []byte {
	padding := PADDING - (len(in) % PADDING)
	if padding == 0 {
		padding = PADDING
	}
	for i := 0; i < padding; i++ {
		in = append(in, byte(padding))
	}
	return in
}

func Unpadding(in []byte) []byte {
	if len(in) == 0 {
		return nil
	}

	padding := in[len(in)-1]
	if int(padding) > len(in) || padding > PADDING {
		return nil
	} else if padding == 0 {
		return nil
	}

	for i := len(in) - 1; i > len(in)-int(padding)-1; i-- {
		if in[i] != padding {
			return nil
		}
	}
	return in[:len(in)-int(padding)]
}

func AesEncrypt(text []byte, key []byte, appId string) (string, error) {
	all := make([]byte, 0)
	randBs := make([]byte, 16)
	io.ReadFull(rand.Reader, randBs)

	lenBs := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBs, uint32(len(text)))

	all = append(all, randBs...)        // 16 rand bytes
	all = append(all, lenBs...)         // 4 length bytes
	all = append(all, []byte(text)...)  // msg content
	all = append(all, []byte(appId)...) // appId content

	enBs, err := aesEncrypt(all, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(enBs), nil
}

func aesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	data = Padding(data)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	enBs := make([]byte, len(data))
	blockMode.CryptBlocks(enBs, data)
	return enBs, nil
}

func AesDecrypt(encText string, key []byte) (*RawMsg, error) {
	cipherData, err := base64.StdEncoding.DecodeString(encText)
	if err != nil {
		return nil, err
	}

	rawData, err := aesDecrypt(cipherData, key)
	if err != nil {
		return nil, err
	}
	if len(rawData) <= 20 {
		return nil, errors.New("Error content length")
	}

	m := &RawMsg{}
	binary.Read(bytes.NewBuffer(rawData[16:20]), binary.BigEndian, &m.Length)
	m.Data = rawData[20 : 20+m.Length]
	m.AppId = string(rawData[20+m.Length:])
	if len(rawData) < int(20+m.Length) {
		return nil, errors.New("Error length content")
	}
	return m, nil
}

func aesDecrypt(data []byte, key []byte) ([]byte, error) {
	keyLen := len(key)
	if len(data)%keyLen != 0 {
		return nil, errors.New("ciphertext size is not multiple of aes key length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	rawData := make([]byte, len(data))
	blockMode.CryptBlocks(rawData, data)
	rawData = Unpadding(rawData)
	if rawData == nil {
		return nil, errors.New("unpadding failed")
	}
	return rawData, nil
}
