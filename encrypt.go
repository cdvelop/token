package token

import (
	"encoding/hex"
)

// Encrypt
func (t *config) Encrypt(content string) (out, err string) {
	const this = "Encrypt error "

	// fmt.Println("PRUEBA LLAVE SETEADA:", t.encryption_key)

	bytes, err := compress([]byte(content))
	if err != "" {
		return "", this + err
	}
	// fmt.Println("EncryptionKey:", t.EncryptionKey)

	return hex.EncodeToString(bytes), ""
}

// Decrypt
func (t *config) Decrypt(content string) (out, err string) {
	const this = "Decrypt error: "

	bytes, er := hex.DecodeString(content)
	if er != nil {
		return "", this + er.Error()
	}

	return decompress(bytes)
}
