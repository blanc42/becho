package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

func GenerateSignature(secret string, expire int64) string {
	h := hmac.New(sha256.New, []byte(secret))
	_, _ = h.Write([]byte(strconv.FormatInt(expire, 10)))
	return hex.EncodeToString(h.Sum(nil))
}

type SignedParams struct {
	Signature string `json:"signature"`
	Expire    string `json:"expire"`
	PublicKey string `json:"public_key"`
}

func GetUploadcareSignedParams() SignedParams {
	secret := os.Getenv("UPLOADCARE_SECRET")
	expire := time.Now().Add(30 * time.Minute).Unix()
	expireStr := fmt.Sprintf("%d", expire) // Convert int64 to string
	signature := GenerateSignature(secret, expire)

	return SignedParams{
		Signature: signature,
		Expire:    expireStr,
		PublicKey: os.Getenv("UPLOADCARE_PUBLIC"),
	}
}
