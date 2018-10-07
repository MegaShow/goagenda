package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func hash256(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func getSalt(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Encrypt(password string) (data, salt string) {
	salt = getSalt(6)
	data = hash256(hash256(password) + salt)
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
