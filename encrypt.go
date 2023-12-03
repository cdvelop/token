package token

import (
	"encoding/hex"
)

// Encrypt
func (t *Token) Encrypt(content string) (out, err string) {
	const this = "Encrypt error "

	bytes, err := compress([]byte(content))
	if err != "" {
		return "", this + err
	}
	// fmt.Println("EncryptionKey:", t.EncryptionKey)

	return hex.EncodeToString(bytes), ""
}

// Decrypt
func (t *Token) Decrypt(content string) (out, err string) {
	const this = "Decrypt error: "

	bytes, er := hex.DecodeString(content)
	if er != nil {
		return "", this + er.Error()
	}

	return decompress(bytes)
}
