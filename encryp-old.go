package token

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Encrypt cifra el contenido y genera una suma de verificación simple.
func (t *Token) EncryptOLD(content string) (out string) {
	// Convertir clave a bytes
	keyBytes := []byte(t.EncryptionKey)
	contentBytes := []byte(content)

	encryptedSha256KeyWithContent := sha256.Sum256(append(keyBytes, contentBytes...))
	// Paso 1: Cifrar contenido con la clave
	// encryptedSha256KeyWithContent := sha256.Sum256([]byte(t.EncryptionKey + content))

	// Paso 2: Suma de verificación (SHA-256 como ejemplo)
	hash := sha256.New()
	hash.Write(encryptedSha256KeyWithContent[:])

	// Paso 3: Obtener la suma de verificación como una cadena hexadecimal
	verificationHex := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("Encrypt verificationHex:", verificationHex)

	// Paso 4: convertir a hex llave mas contenido encriptado
	encryptedKeyWithContentHEX := hex.EncodeToString(encryptedSha256KeyWithContent[:])
	fmt.Println("Encrypt encryptedKeyWithContentHEX:", encryptedKeyWithContentHEX)

	// Paso 5: Concatenar el contenido cifrado y la suma de verificación
	return encryptedKeyWithContentHEX + verificationHex
}
func (t *Token) DecryptOLD(content string) (out, err string) {
	const this = "Decrypt error: "

	// Verificar si la longitud del contenido es suficiente para contener la suma de verificación
	if len(content) < 64 {
		return "", this + "contenido cifrado incorrecto"
	}

	// 5 separar contenido
	//5.1 Tomar los últimos 64 caracteres como la suma de verificación en ex
	verificationHex := content[len(content)-64:]
	fmt.Println("Decrypt verificationHex:", verificationHex)

	//5.2 obtener llave + contenido cifrado en hex
	encryptedKeyWithContentHEX := content[:len(content)-64]
	fmt.Println("Decrypt encryptedKeyWithContentHEX:", encryptedKeyWithContentHEX)

	// Paso 4: Decodificar a byte la cadena hex llave mas contenido encriptado
	encryptedSha256KeyWithContent, er := hex.DecodeString(encryptedKeyWithContentHEX)
	if er != nil {
		return "", this + "decodificación hex fallida: " + er.Error()
	}

	// Paso 3: Verificar la suma de verificación (SHA-256 como ejemplo)
	hash := sha256.New()
	hash.Write(encryptedSha256KeyWithContent)
	calculatedVerification := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("Decrypt calculatedVerification:", calculatedVerification)

	if calculatedVerification != verificationHex {
		return "", this + "suma de verificación no válida"
	}

	// Paso 2: Descifrar el contenido sin la clave
	outBytes := encryptedSha256KeyWithContent[len(t.EncryptionKey)/2:]

	// Paso 1: Convertir a cadena el contenido descifrado
	out = string(outBytes)

	return out, ""
}
