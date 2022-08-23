package youdu

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
)

type DecryptResult struct {
	encryptor *encryptor

	AppId  string
	Data   string
	Length int32
}

func (d *DecryptResult) Unmarshal(v interface{}) error {
	return json.Unmarshal([]byte(d.Data), &v)
}

// encryptor is used to encrypt and decrypt messages.
// see: https://gist.github.com/STGDanny/03acf29a90684c2afc9487152324e832
type encryptor struct {
	config *Config
	pkcs7  *Pkcs7
}

func NewEncryptor(config *Config) *encryptor {
	return &encryptor{
		config: config,
		pkcs7:  NewPkcs7(),
	}
}

func (e *encryptor) Encrypt(plaintext string) (string, error) {
	// key
	key, err := base64.StdEncoding.DecodeString(e.config.AesKey)
	if err != nil {
		return "", err
	}

	// plainText
	plainText := make([]byte, 0)

	randBs := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, randBs)
	if err != nil {
		return "", err
	}

	lenBs := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBs, uint32(len([]byte(plaintext))))

	plainText = append(plainText, randBs...)
	plainText = append(plainText, lenBs...)
	plainText = append(plainText, []byte(plaintext)...)
	plainText = append(plainText, []byte(e.config.AppId)...)

	// encrypt
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	plainText = e.pkcs7.Padding(plainText)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (e *encryptor) Decrypt(ciphertext string) (*DecryptResult, error) {
	// key
	key, err := base64.StdEncoding.DecodeString(e.config.AesKey)
	if err != nil {
		return nil, err
	}

	// cipherText
	cipherText, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	// len valid
	if len(cipherText)%len(key) != 0 {
		return nil, errors.New("invalid ciphertext")
	}

	// aes decrypt
	block, err := aes.NewCipher(key)
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
	plainText = e.pkcs7.Unpadding(plainText)

	// rawMessage
	result := &DecryptResult{}
	if err := binary.Read(bytes.NewBuffer(plainText[16:20]), binary.BigEndian, &result.Length); err != nil {
		return nil, err
	}
	result.Data = string(plainText[20 : 20+result.Length])
	result.AppId = string(plainText[20+result.Length:])
	if len(plainText) < int(20+result.Length) {
		return nil, errors.New("invalid ciphertext")
	}

	return result, err
}

// Pkcs7 is used to padding and unpadding messages.
type Pkcs7 struct {
	blockSize int
}

// NewPkcs7 is used to create a new Pkcs7.
func NewPkcs7() *Pkcs7 {
	return &Pkcs7{
		blockSize: 32,
	}
}

// Padding is used to padding messages.
func (p *Pkcs7) Padding(content []byte) []byte {
	padding := p.blockSize - (len(content) % p.blockSize)

	if padding == 0 {
		padding = p.blockSize
	}

	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(content, padtext...)
}

// Unpadding is used to unpadding messages.
func (p *Pkcs7) Unpadding(content []byte) []byte {
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
