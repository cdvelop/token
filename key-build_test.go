package token_test

import (
	"testing"

	"github.com/cdvelop/token"
)

func TestBuildUniqueKey(t *testing.T) {
	const numKeys = 1000
	uniqueKeys := make(map[string]struct{})

	for i := 0; i < numKeys; i++ {
		key := token.BuildUniqueKey(16) // Cambia el tamaño según sea necesario
		if key == "" {
			t.Error("Error generando la clave única")
		}

		// Verificar que la clave no se repita
		if _, exists := uniqueKeys[key]; exists {
			t.Errorf("Clave duplicada encontrada: %s", key)
		}

		// Almacenar la clave en el mapa para futuras verificaciones
		uniqueKeys[key] = struct{}{}
	}

	// fmt.Println(uniqueKeys)
}

func BenchmarkBuildUniqueKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = token.BuildUniqueKey(16) // Cambia el tamaño según sea necesario
	}
}
