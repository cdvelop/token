package token

import (
	"crypto/rand"
	"encoding/base64"
)

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
// size ej: 16, 32, 64
func BuildUniqueKey(sizeKey uint8) string {
	b := make([]byte, sizeKey)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
