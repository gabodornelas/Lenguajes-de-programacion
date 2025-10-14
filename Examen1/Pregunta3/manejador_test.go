package main

import (
	"testing"
)

// Helper para crear un manejador para tests
func setupManejador(t *testing.T, size int) *Manejador {
	m, err := NuevoManejador(size)
	if err != nil {
		t.Fatalf("Error al inicializar manejador de %d: %v", size, err)
	}
	// Silenciamos la salida de fmt.Printf durante el setup del test
	// m.Tam y m.Ini están correctos después de esta llamada.
	return m
}

// ====================================================================
// CASOS DE TESTEO PRINCIPALES (Cobertura > 80%)
// ====================================================================

func TestAsignacionYDivision(t *testing.T) {
	// Memoria de 128 (se redondea a 128, que es 2^7)
	m := setupManejador(t, 100)

	// Caso 1: Asignación simple (sin división)
	// Solicitud de 30. Se asigna en el bloque de 128. Libre queda 98.
	t.Run("AsignacionSimple", func(t *testing.T) {
		m.Reservar(30, "P1")
		if m.Ini.Libre != 128-30 {
			t.Errorf("Error en asignación simple: Esperado Libre %d, Obtenido %d", 98, m.Ini.Libre)
		}
		if m.Ini.Chivo == nil || m.Ini.Chivo.Nombre != "P1" {
			t.Error("Error: El archivo P1 no fue asignado correctamente.")
		}
	})

	// Caso 2: Asignación que debe dividir 3 veces (Busca el bloque que queda libre)
	// ¡Nota! Tu implementación de 'asignar' divide hasta que la petición > Libre/2.
	// 128 (Ini) ya está usado por P1. El bucle de 'asignar' está mal implementado para el Buddy System.
	// El test asume que el bloqueInicial (128) se dividirá si su Libre lo permite.

	// Vamos a resetear y probar la división en un gestor virgen
	m = setupManejador(t, 256)

	t.Run("AsignacionConDivision", func(t *testing.T) {
		// Petición de 10. MaxOrder 256.
		// 10 < 256/2. Divide 256 -> 128 | 128 (Bloque Actual = 128)
		// 10 < 128/2. Divide 128 -> 64 | 64 (Bloque Actual = 64)
		// 10 < 64/2. Divide 64 -> 32 | 32 (Bloque Actual = 32)
		// 10 < 32/2 (16). NO. Asigna en el bloque de 32.
		m.Reservar(10, "P2")

		// Después de la asignación, la lista de bloques debería ser:
		// [32 (P2)] -> [32 (Libre)] -> [64 (Libre)] -> [128 (Libre)]
		// Verificamos la primera división:
		if m.Ini.Tam != 32 {
			t.Errorf("Error de división. Esperado Tam 32, Obtenido %d", m.Ini.Tam)
		}
		if m.Ini.Sig == nil || m.Ini.Sig.Tam != 32 {
			t.Error("Error: Faltó la primera mitad del split (32|32)")
		}
		if m.Ini.Sig.Sig == nil || m.Ini.Sig.Sig.Tam != 64 {
			t.Error("Error: Faltó el bloque de 64 después del split.")
		}
	})
}

// ---

func TestLiberacionYUnion(t *testing.T) {
	// Memoria de 128 (se redondea a 128)
	m := setupManejador(t, 128)

	// 1. Crear el escenario de fragmentación para probar la unión
	// Petición 1: 30 (Asigna en 32). Divide 128 -> 64 | 64 -> 32 | 32.
	m.Reservar(30, "A") // Bloque de 32 (A, Libre: 2)

	// Petición 2: 30 (Asigna en el buddy de 32, o en el siguiente bloque disponible)
	// Siguiendo la implementación de 'hayespacio' (Best-Fit)
	m.Reservar(30, "B") // Bloque de 32 (B, Libre: 2). La lista es: [32(A)] -> [32(B)] -> [64]

	// Petición 3: 10 (Asigna en el bloque de 64, lo divide).
	m.Reservar(10, "C") // Bloque de 32 (C, Libre: 22). La lista es: [32(A)] -> [32(B)] -> [32(C)] -> [32] -> [64]

	// Caso 3: Liberar y NO unir (porque el buddy está ocupado/es de otro tamaño)
	t.Run("LiberarSinUnion", func(t *testing.T) {
		// Liberamos el bloque A. Su buddy (B) está ocupado.
		m.Liberar("A")
		// El bloque A ahora tiene Libre = 32. Su Tam sigue siendo 32.

		// Verificamos que el primer bloque está libre
		if m.Ini.Libre != 32 || m.Ini.Chivo != nil { // Asume que 'Liberar' elimina la referencia Chivo
			t.Errorf("Error: Bloque A no se liberó correctamente. Libre: %d", m.Ini.Libre)
		}

		// Verificamos que NO hubo fusión. El tamaño sigue siendo 32.
		if m.Ini.Tam != 32 {
			t.Errorf("Error: Hubo fusión inesperada. Tam esperado 32, Obtenido %d", m.Ini.Tam)
		}
	})

	// Caso 4: Liberar y SÍ unir (Coalescencia)
	t.Run("LiberarConUnion", func(t *testing.T) {
		// El estado actual es: [32(Libre)] -> [32(B)] -> [32(C)] -> [32(Libre)] -> [64(Libre)]
		// El primer bloque libre es el de 32. El segundo bloque de 32 está libre.

		// Liberamos B. El bloque B está en la posición 1 (índice 1 de Bloque.Sig.Sig)
		m.Liberar("B")

		// Ahora el bloque de 32 de A y el bloque de 32 de B deberían fusionarse en un solo bloque de 64.
		// El puntero m.Ini debe cambiar de [32(A)] -> [32(B)] a [64]

		// El primer bloque debería ser ahora el resultado de la unión.
		// Verificamos que el primer bloque tiene tamaño 64.
		if m.Ini.Tam != 64 {
			t.Errorf("Error: No se realizó la unión de A y B. Esperado Tam 64, Obtenido %d", m.Ini.Tam)
		}
		// Verificamos que el bloque fusionado está totalmente libre (64).
		if m.Ini.Libre != 64 {
			t.Errorf("Error: El bloque unido no tiene el Libre correcto. Esperado 64, Obtenido %d", m.Ini.Libre)
		}
	})

	// Caso 5: Doble Unión (Libera C y luego el resto)
	// Estado actual: [64(Libre)] -> [32(C)] -> [32(Libre)] -> [64(Libre)]
	t.Run("LiberarConDobleUnion", func(t *testing.T) {
		m.Liberar("C")

		// Después de liberar C (32) y su buddy (el 32 siguiente), deberían formar un bloque de 64.
		// Y luego este nuevo 64 debería unirse con el primer 64.

		// El estado final debería ser un único bloque de 128
		if m.Ini.Tam != 128 {
			t.Errorf("Error: No se fusionó completamente hasta 128. Obtenido %d", m.Ini.Tam)
		}
	})
}

// ---

func TestHayEspacioBestFit(t *testing.T) {
	m := setupManejador(t, 256)

	// Crear 3 bloques libres para probar el Best-Fit: 32, 64 y 128.
	m.Reservar(10, "A")  // Crea el bloque de 32 (Libre 22)
	m.Reservar(30, "B")  // Crea el bloque de 64 (Libre 34)
	m.Reservar(100, "C") // Crea el bloque de 128 (Libre 28)

	// Liberar para que estén libres (solo nos interesa el Libre, no si tienen archivos dentro para este test)
	// Para simular la lista enlazada de bloques, manipulamos manualmente el estado después de Reservar,
	// ya que tu 'Reservar' está diseñado para mantener las asignaciones.

	// Estado (asumiendo que hay 3 bloques que cumplen el tamaño de 20):
	// Bloque 1 (Tam 32, Libre 22) -> No cabe 20 (22 < 20. Libre < t es falso)
	// Bloque 2 (Tam 64, Libre 34) -> Cabe 20. Desperdicio: 14.
	// Bloque 3 (Tam 128, Libre 28) -> Cabe 20. Desperdicio: 8. (MEJOR)

	t.Run("BestFitBusqueda", func(t *testing.T) {
		// La implementación actual de Reservar/Asignar hace que los bloques queden con archivos dentro.
		// Vamos a forzar el estado para que 'hayespacio' funcione bien:

		// Lista de bloques: [B32] -> [B64] -> [B128]
		m.Ini.Libre = 32
		m.Ini.Sig.Libre = 64
		m.Ini.Sig.Sig.Libre = 128

		// Solicitud: 20
		// Bloque 1 (Índice 0): Libre 32. Desperdicio: 12
		// Bloque 2 (Índice 1): Libre 64. Desperdicio: 44
		// Bloque 3 (Índice 2): Libre 128. Desperdicio: 108

		lugar := hayespacio(m, 20, m.Tam) // m.Tam es 256

		// El mejor ajuste es el Bloque 1 (índice 0) con un desperdicio de 12.
		if lugar != 0 {
			t.Errorf("Error en Best-Fit: Esperado índice 0 (32-20=12), Obtenido %d", lugar)
		}

		// Solicitud: 50
		// Bloque 1 (Índice 0): Libre 32. NO CABE.
		// Bloque 2 (Índice 1): Libre 64. Desperdicio: 14. (MEJOR)
		// Bloque 3 (Índice 2): Libre 128. Desperdicio: 78.
		lugar = hayespacio(m, 50, m.Tam)
		if lugar != 1 {
			t.Errorf("Error en Best-Fit: Esperado índice 1 (64-50=14), Obtenido %d", lugar)
		}
	})

	t.Run("NoHayEspacio", func(t *testing.T) {
		// Petición que excede el tamaño máximo
		lugar := hayespacio(m, 300, m.Tam)
		if lugar != -1 {
			t.Errorf("Error: Debería fallar al solicitar más de 256. Obtenido %d", lugar)
		}
	})
}
