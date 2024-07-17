package vhash

import (
	"crypto"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HMacSHA1(key, str string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func HMacSHA256(key, str string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func HMac(hash crypto.Hash, key, str string) (string, error) {
	if !hash.Available() {
		return "", fmt.Errorf("crypto: requested hash function (%s) is unavailable", hash.String())
	}
	h := hmac.New(hash.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil)), nil
}
