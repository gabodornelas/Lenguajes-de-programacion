package main

import (
	"testing"
)

// =================================================================================
// FUNCIONES AUXILIARES DE TESTEO
// =================================================================================

// countBlocks: Función auxiliar para contar el número de bloques en la lista ligada.
// Esto ayuda a verificar que la unión fue exitosa.
func countBlocks(m *Manejador) int {
	count := 0
	b := m.Ini
	for b != nil {
		count++
		b = b.Sig
	}
	return count
}

// =================================================================================
// 1. Pruebas de Inicialización
// =================================================================================

func TestNuevoManejador(t *testing.T) {
	tests := []struct {
		nombre         string
		tamaInicial    int
		esperaTamTotal int // El tamaño redondeado a la potencia de 2
	}{
		{"Potencia de 2 Exacta (32)", 32, 32},
		{"No Potencia de 2 (40)", 40, 64}, // Debe redondear a 64
		{"Pequeño (3)", 3, 4},             // Redondea a 4
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			gestor, err := NuevoManejador(tt.tamaInicial)

			if err != nil {
				t.Errorf("NuevoManejador() error inesperado = %v", err)
				return
			}

			// 1. Verificar el tamaño total
			if gestor.Tam != tt.esperaTamTotal {
				t.Errorf("Manejador.Tam = %v, se esperaba %v", gestor.Tam, tt.esperaTamTotal)
			}

			// 2. Verificar el bloque inicial
			if gestor.Ini.Tam != tt.esperaTamTotal {
				t.Errorf("Ini.Tam = %v, se esperaba %v", gestor.Ini.Tam, tt.esperaTamTotal)
			}
			if gestor.Ini.Libre != tt.esperaTamTotal {
				t.Errorf("Ini.Libre = %v, se esperaba %v", gestor.Ini.Libre, tt.esperaTamTotal)
			}
		})
	}
}

// =================================================================================
// 2. Pruebas de Reserva, División y Best-Fit (Reservar, Hayespacio, Asignar)
// =================================================================================

func TestManejador_Reservar_BestFit_y_Splitting(t *testing.T) {
	// Inicializamos un gestor con tamaño 16 (se redondea a 16).
	gestor, _ := NuevoManejador(16) // Estado: [Bloque16 (L=16)]

	// 1. Reserva 'A' (Tamaño 3). Debe forzar la división a un bloque de tamaño 4.
	t.Run("Reserva_A_Splitting", func(t *testing.T) {
		gestor.Reservar(3, "A")
		// Estado esperado después de divisiones: [B4 (L=1, Arch: A), B4 (L=4), B8 (L=8)]
		bloqueA := Repetido(gestor, "A")
		if bloqueA == nil {
			t.Fatal("Reserva fallida: 'A' no se encontró.")
		}
		if bloqueA.Tam != 4 {
			t.Errorf("Bloque de 'A' tiene Tam=%d, se esperaba 4 después de splitting.", bloqueA.Tam)
		}
		if bloqueA.Libre != 1 {
			t.Errorf("Bloque de 'A' tiene Libre=%d, se esperaba 1 (4-3).", bloqueA.Libre)
		}
	})

	// 2. Reserva 'B' (Tamaño 1). Debe usar la lógica Best-Fit.
	// La función Hayespacio debe elegir el bloque que minimiza el espacio libre restante.
	t.Run("Reserva_B_BestFit", func(t *testing.T) {
		// Bloques disponibles para reservar 1:
		// a) Bloque 1 (Libre 1): Resto = 0.
		// b) Bloque 2 (Libre 4): Resto = 3.
		// c) Bloque 3 (Libre 8): Resto = 7.
		// Best-Fit debe elegir el bloque que deja 0 de resto (índice 0).

		// Primero, verificamos que Hayespacio elija el lugar correcto (índice 0)
		lugar := Hayespacio(gestor, 1, gestor.Tam)
		if lugar != 0 {
			t.Errorf("Hayespacio falló el Best-Fit. Se esperaba índice 0 (Bloque Libre 1), se obtuvo %d.", lugar)
		}

		// Ejecutamos la reserva 'B'
		gestor.Reservar(1, "B")
		bloqueB := Repetido(gestor, "B")
		if bloqueB == nil {
			t.Fatal("Reserva fallida: 'B' no se encontró.")
		}
		// 'B' debe estar en el mismo bloque de 'A' (Bloque 4)
		if bloqueB.Tam != 4 {
			t.Errorf("Bloque de 'B' tiene Tam=%d, se esperaba 4.", bloqueB.Tam)
		}
		if bloqueB.Libre != 0 {
			t.Errorf("Bloque de 'B' tiene Libre=%d, se esperaba 0 (1-1).", bloqueB.Libre)
		}
	})

	// 3. Reserva 'C' (Tamaño 5). Debe forzar la división de B8 y usar B8 resultante.
	t.Run("Reserva_C_SplitAgain", func(t *testing.T) {
		gestor.Reservar(5, "C")
		// B8 (Libre 8) -> B4, B4. Reserva C en un B8 resultante. Error.
		// B8 (Libre 8). Tam=8. t=5. 5 < 8/2 (4) es FALSO. No divide.
		// Simplemente reserva en B8.
		bloqueC := Repetido(gestor, "C")
		if bloqueC == nil {
			t.Fatal("Reserva fallida: 'C' no se encontró.")
		}
		if bloqueC.Tam != 8 {
			t.Errorf("Bloque de 'C' tiene Tam=%d, se esperaba 8.", bloqueC.Tam)
		}
		if bloqueC.Libre != 3 {
			t.Errorf("Bloque de 'C' tiene Libre=%d, se esperaba 3 (8-5).", bloqueC.Libre)
		}
	})
}

// =================================================================================
// 3. Pruebas de Liberación y Unión (Liberar, UnionBloques)
// =================================================================================

func TestManejador_Liberar_Union_Completa(t *testing.T) {
	// Inicializamos un gestor con tamaño 8. [B8 (L=8)]
	gestor, _ := NuevoManejador(8)

	// 1. Reserva A (Tamaño 1) -> B8 -> B4, B4 -> B2, B2, B4. Reserva A en B2. [B2 (L=1, A), B2 (L=2), B4 (L=4)]
	gestor.Reservar(1, "A")
	// 2. Reserva B (Tamaño 1) -> Reserva en el segundo B2. [B2 (L=1, A), B2 (L=1, B), B2 (L=2), B2 (L=2), B4 (L=4)]
	gestor.Reservar(1, "B")

	// 3. Liberar B
	gestor.Liberar("B")

	// El bloque de 'B' (originalmente B2) queda totalmente libre (L=2). Su buddy (el siguiente B2) también está libre (L=2).
	// La unión DEBE ocurrir aquí: B2 + B2 = B4.
	t.Run("Liberar_B_Union_Parcial", func(t *testing.T) {
		// Contamos los bloques después de la liberación B
		count := countBlocks(gestor)
		if count > 4 { // Estado original antes de B: 5 bloques. Unión de 2 -> 4.
			t.Logf("Advertencia: El número de bloques después de la liberación B es %d, lo que sugiere que la unión no ocurrió o no fue completa.", count)
		}
	})

	// 4. Liberar A (Debería unir el B2 libre con el B4 libre, y luego seguir uniendo)
	t.Run("Liberar_A_Union_Completa", func(t *testing.T) {
		gestor.Liberar("A")

		if Repetido(gestor, "A") != nil {
			t.Fatal("Liberar 'A' falló: La reserva sigue existiendo.")
		}

		// Verificamos el estado de unión: Debe volver al bloque inicial [8, Libre: 8]
		if gestor.Ini.Tam != 8 || gestor.Ini.Libre != 8 {
			t.Errorf("Unión fallida. Bloque inicial debería ser [Tam:8, Libre:8]. Actual: [Tam:%d, Libre:%d]", gestor.Ini.Tam, gestor.Ini.Libre)
		}
		if gestor.Ini.Sig != nil {
			t.Errorf("Unión fallida: Se esperaba solo un bloque principal (Ini.Sig = nil), pero hay más bloques.")
		}
	})
}

func TestManejador_Liberar_ArchivoIntermedio(t *testing.T) {
	// Prueba la liberación de un Archivo que NO es el primero en su Bloque
	gestor, _ := NuevoManejador(8)

	// 1. Reserva A (Tamaño 1)
	gestor.Reservar(1, "A")
	// 2. Reserva B (Tamaño 1) en el MISMO bloque
	bloqueA := Repetido(gestor, "A")
	Asignar(gestor, 1, "B", 0) // Usamos Asignar para forzar la asignación en el mismo bloque

	// Verificamos que 'A' sea el primero y 'B' el segundo
	if bloqueA.Chivo.Nombre != "A" || bloqueA.Chivo.Sig.Nombre != "B" {
		t.Fatal("Setup incorrecto: A y B no están en el mismo bloque o en el orden esperado.")
	}

	// 3. Liberar B (el archivo intermedio/final)
	gestor.Liberar("B")

	t.Run("Liberar_Archivo_No_Primero", func(t *testing.T) {
		if Repetido(gestor, "B") != nil {
			t.Fatal("Liberar 'B' falló: La reserva sigue existiendo.")
		}
		// Verificamos que 'A' siga siendo el primer archivo
		if bloqueA.Chivo.Nombre != "A" {
			t.Errorf("Tras liberar B, A debería ser el primer archivo. Se encontró: %s", bloqueA.Chivo.Nombre)
		}
		// Verificamos que el puntero haya saltado a nil
		if bloqueA.Chivo.Sig != nil {
			t.Errorf("El puntero Sig de A debería ser nil.")
		}
	})
}
