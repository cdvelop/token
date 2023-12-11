package token_test

import (
	"testing"

	"github.com/cdvelop/token"
)

func TestTokenEncryptDecrypt(t *testing.T) {
	// Definir casos de prueba en un mapa
	testCases := map[string]struct {
		Content string
		Expect  string
	}{
		"Case1": {"Hello, secret message.", "Hello, secret message."},
	}

	// Iterar sobre los casos de prueba
	for testName, c := range testCases {
		t.Run(testName, func(t *testing.T) {

			// Crear una instancia de Token
			token := token.AddCipherAdapter(token.BuildUniqueKey(16))

			// Encriptar el contenido
			encrypted_content, err := token.Encrypt([]byte(c.Content))
			if err != "" {
				t.Fatal(err)
				return
			}

			// fmt.Println(encrypted_content)

			// Probar la desencriptación
			result, err := token.Decrypt(encrypted_content)
			if err != "" {
				result = err
			}

			// Verificar el resultado
			if result != c.Expect {
				t.Fatalf("Error :%v\n-Se esperaba: %s\n-pero se Obtuvo: %s", testName, c.Expect, result)
				return
			}

		})
	}
}
