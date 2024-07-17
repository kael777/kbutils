package xcrypto

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSACrypto(t *testing.T) {
	privateKey, publicKey, err := GenerateRSAKey(2048, RSA_PKCS1)
	assert.Nil(t, err)

	data := "ILoveYouVazhenina"

	pvtKey, err := NewPrivateKeyFromPemBlock(RSA_PKCS1, privateKey)
	assert.Nil(t, err)

	pubKey, err := NewPublicKeyFromPemBlock(RSA_PKCS1, publicKey)
	assert.Nil(t, err)

	cipher, err := pubKey.Encrypt([]byte(data))
	assert.Nil(t, err)

	plain, err := pvtKey.Decrypt(cipher)
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain))

	cipher2, err := pubKey.EncryptOAEP(crypto.SHA256, []byte(data))
	assert.Nil(t, err)

	plain2, err := pvtKey.DecryptOAEP(crypto.SHA256, cipher2)
	assert.Nil(t, err)
	assert.Equal(t, data, string(plain2))
}

func TestRSASign(t *testing.T) {
	publicKey := []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAwWVvD3G+O9N1NuBBz44OLb6aq85w8ahoTRepzydJ2qBcaDh+Zj6M
cybRSGHIGBIG0vyzYiPQhLK+s2kzKJ9rUHkQqRc7zDdVfclJhul1n1oBReyue1q9
AyZXhWssZodeQPG5SnlwziCuVhP6WCLF0M1bkvJr0+VOAfSHeTeYx/S/nH8JErmY
1HQTpkPs/fyabzCKoStWg6D62840HA2gn6Xq1MuPFki+BR8xcaM3Tqp2yN2kkIgO
RcGpTUOMk1L8xXRjTbYT48wyXmeMnR1TtmFE2Xc3sMC8y/mn8V7D4r2alfDHDX4d
13hBzo0oap7tugnr9yA2lak4Nvah03ZprwIDAQAB
-----END RSA PUBLIC KEY-----`)

	privateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAwWVvD3G+O9N1NuBBz44OLb6aq85w8ahoTRepzydJ2qBcaDh+
Zj6McybRSGHIGBIG0vyzYiPQhLK+s2kzKJ9rUHkQqRc7zDdVfclJhul1n1oBReyu
e1q9AyZXhWssZodeQPG5SnlwziCuVhP6WCLF0M1bkvJr0+VOAfSHeTeYx/S/nH8J
ErmY1HQTpkPs/fyabzCKoStWg6D62840HA2gn6Xq1MuPFki+BR8xcaM3Tqp2yN2k
kIgORcGpTUOMk1L8xXRjTbYT48wyXmeMnR1TtmFE2Xc3sMC8y/mn8V7D4r2alfDH
DX4d13hBzo0oap7tugnr9yA2lak4Nvah03ZprwIDAQABAoIBAB80zeHxGaAvs9dC
AnyKUJFjEzQr4J+t6/6cleL+VPV5MNAEZaj76M/f8J88X/w6VG2RJyTr4Ia5DPqI
PCAO8VMP5fdS72w5dYsRgtLJMxieflwZH+J5tsweULsPmx+EMlpKZvq0c9ZfAaKU
IK4+FitmJ6OjiHCtrJO2MHIH3ZhOBxn032BfdyVqhNN+oyn0zSjXvpHg9t/UEsXp
ZA7rHYn7m0RTwynFSaouAhmmZAp2GTYhe0NFu8rCG5afhtw9H2XiIiOhmLcURG+P
oW8v3I/Vt0OoLcqilbjPJs6nd43CAVyGastcBXhDFJJ4mFw5itMV9c+XNsEXPDcD
2g2voqECgYEA38UTnGv1eciGNcYMWUDJIB1c/205GoSpQ2kHXkNbFdN7u9lGlopq
3NwUPpHgbuWR5VxPmZCy1hCpFVXyeF9Ea3mFahiyiFECj4MeYq7i8Yd+UIfDNQ99
4C8TJP2mI4a8DaH7qG1KHfpkgaLsYuIhCmm+aNXsqcSNqRjYJtAE+lECgYEA3UBp
F6asT+ztQXF0QC7JOdaJgW6W4RNaIcU5rdK2vkkfhqQzR/XEFmHqVW7qUnLGm4mW
dTS6QBAoLwyd87KXvTW4y5rW2Un+l0Pc59Kl35BdlwMpXCffeqhamS4B7F4AdVZY
JaCYTCkTuwAx2r5nyOlkTcMIEGeDL676dRHII/8CgYEA3gZq+O9dd2JxV/WT1xMi
/ExmM8IpwJgUYiBaATuPqs5VnQNuuHvKoC11oMeZCi+aXRsEl/gsmZ2aRuMqXCka
eBDxQV4T9pF6mu6cPYoM/11TBZBPLdybJs9OjYtnRySuflBUpL8bpTcGdmIzbcG0
yuI03Uw1MBUoAbn27jvEVKECgYBiWxXc671CMqMuKo9xUNsnmRW7sjvkhsPUq2Z+
vWN7p+oZ4rjhToIDKTgRDqOgT2G3Fy0JoY0CmawjbkpxYX1PIaiq6oSER/6jpAl6
DQysG/NfBIrIavlP/7N20RsNxqQRhXbeE0xg3wnkYavIAEkG6aorX34gPMP22KSC
kosUZQKBgDKPXK4tnOC4HzYFlkiRxBuCMxU8bTG1+qKFvp+O4BbniDcUkZGJP/Gp
t6RsET7ZhCU8m8/6gIS5lZRoJt1aoqL3UyfFdWVA8pZwihDnEHvp1+0yl2BBaAN1
Vv8zI7kt+uZxD5mBGglKs2wzaHqADBXa5kSznIvkcZSg07UQQYU6
-----END RSA PRIVATE KEY-----`)

	data := "ILoveYouVazhenina"

	pvtKey, err := NewPrivateKeyFromPemBlock(RSA_PKCS1, privateKey)
	assert.Nil(t, err)

	pubKey, err := NewPublicKeyFromPemBlock(RSA_PKCS1, publicKey)
	assert.Nil(t, err)

	sign, err := pvtKey.Sign(crypto.SHA1, []byte(data))
	assert.Nil(t, err)
	assert.Equal(t, "kMcktQrGMRN4SvkCNwdsFD+JnszD7V5o9X5BzYKFExlpXMzsWIBWhcSZ+yMtX3iFyJD+5Jm66zCRash844YEAf84JikUPYf2Btg5Pdz2aaSFGEieVroauODlUkQkKp/+WM72ZVOtKxnJ8AmDhBLyOj3n0mlf9q0EZA2UuUM2xQ0fWnbE+qiJ1GhAIO7w/4oBccuu3q76vnTpmPQWGVJ1EIIRAVWU2basbIiY9EqUP2Np+yH50m+5qOgjSHrG6gJLQrIERJz+F4st24oCtBFdjk6lLGpwgk6ziyXusbhGDfn6wNQ76S+lATnc1T6yNweCcE/ZKS0eFmSwdt3F4vD4+w==", base64.StdEncoding.EncodeToString(sign))
	assert.Nil(t, pubKey.Verify(crypto.SHA1, []byte(data), sign))

	sign2, err := pvtKey.Sign(crypto.SHA256, []byte(data))
	assert.Nil(t, err)
	assert.Equal(t, "laxwcVNCxHkducB6YosXAmuzJHNTOZZvMtgw1BmJLljytw1559/SoRbSXT/ycRle9jtuLMPcXdnlrtoUMn+FTDDOha1cRYdJUg2CXLn6M9S7SXo54rUXL0UxAqk0aod8/4V8jeejttfFU/tgMGn9MjSYlRtDJKVzCvcvlGyeIY5v2RfiDp8+3iFCmTtgCxYS+8Vrj+Du50Q/FjR/n5hO24FwRp2TW+mEEkuTJakee6j1IMrNh2HW+PZdhE8he9w5U+01vwDBZwiDTMgpqs77d5qHsReGj+r4ZVRosVDpUisxbB8YuDNc8Oelx6dZFtlAzaTC2FrQNoJyHzBywuxLBQ==", base64.StdEncoding.EncodeToString(sign2))
	assert.Nil(t, pubKey.Verify(crypto.SHA256, []byte(data), sign2))

	signPSS, err := pvtKey.SignPSS(crypto.SHA1, []byte(data), &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash})
	assert.Nil(t, err)
	assert.Nil(t, pubKey.VerifyPSS(crypto.SHA1, []byte(data), signPSS, &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash}))

	signPSS2, err := pvtKey.SignPSS(crypto.SHA256, []byte(data), &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash})
	assert.Nil(t, err)
	assert.Nil(t, pubKey.VerifyPSS(crypto.SHA256, []byte(data), signPSS2, &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthEqualsHash}))
}
