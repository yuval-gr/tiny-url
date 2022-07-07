package main

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

func getSHA1Hash(s *string) string {

	hasher := sha1.New()
	hasher.Write([]byte(*s))
	base64Hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return strings.ToLower(base64Hash[:7])
}
