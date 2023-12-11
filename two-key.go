package token

type TwoKeys struct {
	// SizeKey uint8// default 16
	execution_count int
	encryption_key  string
}

// retorna "main.encryption_key":"1fgd3..." en ejecución impar cambia la llave
func (k *TwoKeys) GetTwoPublicKeysWasmClientAndGoServer() map[string]string {
	// Incrementar el contador
	k.execution_count++

	// Verificar si el contador es impar
	if k.execution_count%2 != 0 {
		// fmt.Println("Generando una nueva clave:", k.encryption_key)
		k.encryption_key = BuildUniqueKey(16)
	}

	return map[string]string{
		"github.com/cdvelop/token.public_key_client_and_server": k.encryption_key,
	}

}

var public_key_client_and_server = ""

type TwoPublicKeyAdapter interface {
	GetTwoPublicKeysWasmClientAndGoServer() map[string]string
}
