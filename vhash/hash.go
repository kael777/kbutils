package vhash

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Hash(hash crypto.Hash, str string) (string, error) {
	if !hash.Available() {
		return "", fmt.Errorf("crypto: requested hash function (%s) is unavailable", hash.String())
	}
	h := hash.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil)), nil
}
