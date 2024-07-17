package xcrypto

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCBC(t *testing.T) {
	key := "AES256Key-32Characters1234567890"
	iv := key[:16]
	data := "ILoveYouVazhenina"

	cipher, err := AESEncryptCBC([]byte(key), []byte(iv), []byte(data))
	assert.Nil(t, err)
	assert.Equal(t, "48cuX6vB10b6hrnAo0yppnkaz3BFEMfCfNRvpatF51A=", cipher.String())

	plain, err := AESDecryptCBC([]byte(key), []byte(iv), cipher.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain))

	cipher2, err := AESEncryptCBC([]byte(key), []byte(iv), []byte(data), 32)
	assert.Nil(t, err)
	assert.Equal(t, "48cuX6vB10b6hrnAo0yppnkaz3BFEMfCfNRvpatF51A=", cipher2.String())

	plain2, err := AESDecryptCBC([]byte(key), []byte(iv), cipher2.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain2))
}

func TestAesECB(t *testing.T) {
	key := "AES256Key-32Characters1234567890"
	data := "ILoveYouVazhenina"

	cipher, err := AESEncryptECB([]byte(key), []byte(data))
	assert.Nil(t, err)
	assert.Equal(t, "0k38YYk27Sv9Nc6248PIGXr7hYc0gL0ID8EXNf08Pb8=", cipher.String())

	plain, err := AESDecryptECB([]byte(key), cipher.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain))

	cipher2, err := AESEncryptECB([]byte(key), []byte(data), 32)
	assert.Nil(t, err)
	assert.Equal(t, "0k38YYk27Sv9Nc6248PIGXr7hYc0gL0ID8EXNf08Pb8=", cipher2.String())

	plain2, err := AESDecryptECB([]byte(key), cipher.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain2))
}

func TestAesCFB(t *testing.T) {
	key := "AES256Key-32Characters1234567890"
	iv := key[:16]
	data := "ILoveYouVazhenina"

	cipher, err := AESEncryptCFB([]byte(key), []byte(iv), []byte(data))
	assert.Nil(t, err)
	assert.Equal(t, "KP7OnZjqJ8S3mgTz3gYDO60=", cipher.String())

	plain, err := AESDecryptCFB([]byte(key), []byte(iv), cipher.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain))
}

func TestAesOFB(t *testing.T) {
	key := "AES256Key-32Characters1234567890"
	iv := key[:16]
	data := "ILoveYouVazhenina"

	cipher, err := AESEncryptOFB([]byte(key), []byte(iv), []byte(data))
	assert.Nil(t, err)
	assert.Equal(t, "KP7OnZjqJ8S3mgTz3gYDO+E=", cipher.String())

	plain, err := AESDecryptOFB([]byte(key), []byte(iv), cipher.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain))
}

func TestAesCTR(t *testing.T) {
	key := "AES256Key-32Characters1234567890"
	iv := key[:16]
	data := "ILoveYouVazhenina"

	cipher, err := AESEncryptCTR([]byte(key), []byte(iv), []byte(data))
	assert.Nil(t, err)
	assert.Equal(t, "KP7OnZjqJ8S3mgTz3gYDO28=", cipher.String())

	plain, err := AESDecryptCTR([]byte(key), []byte(iv), cipher.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain))
}

func TestAesGCM(t *testing.T) {
	key := "AES256Key-32Characters1234567890"
	nonce := key[:12]
	data := "ILoveYouVazhenina"
	aad := "vazhenina"

	cipher, err := AESEncryptGCM([]byte(key), []byte(nonce), []byte(data), []byte(aad), &GCMOption{})
	assert.Nil(t, err)
	assert.Equal(t, "qciumnRZL5IHE3QsBqN/9Bmo3NtCsKGNGneu2pxghru/", cipher.String())
	assert.Equal(t, "qciumnRZL5IHE3QsBqN/9Bk=", base64.StdEncoding.EncodeToString(cipher.Data()))
	assert.Equal(t, "qNzbQrChjRp3rtqcYIa7vw==", base64.StdEncoding.EncodeToString(cipher.Tag()))

	plain, err := AESDecryptGCM([]byte(key), []byte(nonce), cipher.Bytes(), []byte(aad), nil)
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain))
}
