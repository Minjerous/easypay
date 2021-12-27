package tool

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

const SaltSize = 16

//加盐加密

func HashWithSalted(plain string) string {
	buf := make([]byte, SaltSize, SaltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		fmt.Println("random read failed ->", err)
	}

	h := sha1.New()
	h.Write(buf)
	h.Write([]byte(plain))

	return base64.URLEncoding.EncodeToString(h.Sum(buf))
}

//解密判断

func Match(secret, plain string) bool {
	data, _ := base64.URLEncoding.DecodeString(secret)
	if len(data) != SaltSize+sha1.Size {
		fmt.Println("wrong length of data")
		return false
	}
	h := sha1.New()
	h.Write(data[:SaltSize])
	h.Write([]byte(plain))
	return bytes.Equal(h.Sum(nil), data[SaltSize:])
}
