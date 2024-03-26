package youdu

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

type RawData struct {
	AppID  string
	Data   []byte
	Length int32
}

type Encryptor struct {
	key   []byte
	appid string
	pkcs7 *pkcs7
}

func NewEncryptor(key []byte, appid string) *Encryptor {
	return &Encryptor{
		key:   key,
		appid: appid,
		pkcs7: newPkcs7(),
	}
}

func NewEncryptorWithConfig(config *Config) *Encryptor {
	key, err := base64.StdEncoding.DecodeString(config.AesKey)
	if err != nil {
		panic(err)
	}

	return NewEncryptor(key, config.AppID)
}

func (e *Encryptor) Encrypt(plaintext []byte) (string, error) {
	plainText := make([]byte, 0)

	randBs := make([]byte, 16) // nolint:gomnd
	_, err := io.ReadFull(rand.Reader, randBs)
	if err != nil {
		return "", err
	}

	lenBs := make([]byte, 4) // nolint:gomnd
	binary.BigEndian.PutUint32(lenBs, uint32(len(plaintext)))

	plainText = append(plainText, randBs...)
	plainText = append(plainText, lenBs...)
	plainText = append(plainText, plaintext...)
	plainText = append(plainText, []byte(e.appid)...)

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCEncrypter(block, e.key[:block.BlockSize()])
	plainText = e.pkcs7.padding(plainText)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (e *Encryptor) Decrypt(ciphertext string) (*RawData, error) {
	cipherText, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	if len(cipherText)%len(e.key) != 0 {
		return nil, errors.New("invalid ciphertext")
	}

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = e.pkcs7.unpadding(plainText)

	result := &RawData{}
	if err := binary.Read(bytes.NewBuffer(plainText[16:20]), binary.BigEndian, &result.Length); err != nil {
		return nil, err
	}

	result.Data = plainText[20 : 20+result.Length]
	result.AppID = string(plainText[20+result.Length:])
	if len(plainText) < int(20+result.Length) { // nolint:gomnd
		return nil, errors.New("invalid ciphertext")
	}

	return result, err
}

// pkcs7 is used to padding and unpadding messages.
type pkcs7 struct {
	blockSize int
}

// newPkcs7 is used to create a new pkcs7.
func newPkcs7() *pkcs7 {
	return &pkcs7{
		blockSize: 32, // nolint:gomnd
	}
}

// padding is used to padding messages.
func (p *pkcs7) padding(content []byte) []byte {
	padding := p.blockSize - (len(content) % p.blockSize)

	if padding == 0 {
		padding = p.blockSize
	}

	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(content, padtext...)
}

// unpadding is used to unpadding messages.
func (p *pkcs7) unpadding(content []byte) []byte {
	if len(content) == 0 {
		return nil
	}

	padding := content[len(content)-1]
	if int(padding) > len(content) || int(padding) > p.blockSize {
		return nil
	} else if padding == 0 {
		return nil
	}

	for i := len(content) - 1; i > len(content)-int(padding)-1; i-- {
		if content[i] != padding {
			return nil
		}
	}

	return content[:len(content)-int(padding)]
}
