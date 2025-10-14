package main

import (
	"fmt"
	"math"
)

type Archivo struct {
	Empieza int      // direccion donde empieza el bloque
	Tam     int      // tam del archivo
	Nombre  string   // nombre de la reserva
	Sig     *Archivo // el siguiente archivo
}

// Block representa un segmento de memoria en el sistema.
type Bloque struct {
	Empieza int      // direccion donde empieza el bloque
	Tam     int      // talla real del bloque (2^k)
	Libre   int      // espacio libre en el bloque
	Chivo   *Archivo // el archivo que reservo memoria
	Sig     *Bloque  // bloque siguiente
}

// BuddySystemManager maneja la asignación de memoria usando el Buddy System.
type Manejador struct {
	Tam int     // el tamaño total
	Ini *Bloque // bloque inicial
	Fin *Bloque // bloque final
}

func NuevoManejador(tama int) (*Manejador, error) {
	// 1. Calcular el MaxOrder (la menor potencia de 2 >= tama)
	maxOrden := int(math.Ceil(math.Log2(float64(tama))))
	totalbloque := 1 << maxOrden // 2^MaxOrder

	// 2. inicializar el bloque
	Bloqueinicial := Bloque{
		Empieza: 0,
		Tam:     totalbloque,
		Libre:   totalbloque,
		Chivo:   nil,
		Sig:     nil,
	}

	// 3. inicializar manejador
	Gestor := &Manejador{
		Tam: totalbloque,
		Ini: &Bloqueinicial,
	}

	fmt.Printf("\nGestor de Memoria Buddy System inicializado. Peso Total: %d\n", totalbloque)
	return Gestor, nil
}

//-----------------------------------------------------------RESERVAR----------------------------------------------------

func Repetido(m *Manejador, nombre string) *Bloque {
	bloqueActual := m.Ini
	for {
		if bloqueActual == nil {
			return nil
		}
		archivoActual := bloqueActual.Chivo
		for {
			if archivoActual == nil {
				break
			}
			if archivoActual.Nombre == nombre {
				return bloqueActual
			}
			archivoActual = archivoActual.Sig
		}

		bloqueActual = bloqueActual.Sig
	}
}

func NuevoBloque(b *Bloque) *Bloque {
	Bloqueinicial := Bloque{
		Empieza: b.Empieza + b.Tam,
		Tam:     b.Tam,
		Libre:   b.Tam,
		Chivo:   nil,
		Sig:     b.Sig,
	}
	return &Bloqueinicial
}

func NuevoArchivo(b *Bloque, t int, n string) *Archivo {
	Archivoinicial := Archivo{
		Empieza: b.Empieza + b.Tam - b.Libre,
		Tam:     t,
		Nombre:  n,
		Sig:     nil,
	}
	return &Archivoinicial
}

func Asignar(m *Manejador, t int, n string, b int) *Archivo {
	bloqueActual := m.Ini
	for i := 0; i < b; i++ {
		bloqueActual = bloqueActual.Sig
	}
	for {
		if t < bloqueActual.Libre/2 && bloqueActual.Tam/2 < bloqueActual.Libre { //divido
			bloqueActual.Tam = bloqueActual.Tam / 2
			bloqueActual.Libre = bloqueActual.Libre / 2
			nuevobloque := NuevoBloque(bloqueActual)
			bloqueActual.Sig = nuevobloque
		} else { //asigno
			if bloqueActual.Chivo == nil { // no tiene archivo
				bloqueActual.Chivo = NuevoArchivo(bloqueActual, t, n)
				bloqueActual.Libre = bloqueActual.Libre - t
				return bloqueActual.Chivo
			} else { // tiene archivo
				archivoActual := bloqueActual.Chivo
				for archivoActual.Sig != nil {
					archivoActual = archivoActual.Sig
				}
				archivoActual.Sig = NuevoArchivo(bloqueActual, t, n)
				bloqueActual.Libre = bloqueActual.Libre - t
				return archivoActual.Sig
			}
		}
	}
}

func Hayespacio(m *Manejador, t int, esp int) int {
	bloqueActual := m.Ini
	minesp := esp
	iteracion := 0
	numbloque := iteracion
	for {
		if bloqueActual == nil {
			if minesp == esp {
				return -1 // no encontro ningun espacio
			} else {
				return numbloque
			}
		}
		if bloqueActual.Libre >= t && minesp > bloqueActual.Libre-t {
			minesp = bloqueActual.Libre - t
			numbloque = iteracion // guarda la "iteracion" del bloque que queremos
		}
		iteracion++
		bloqueActual = bloqueActual.Sig
	}
}

func (m *Manejador) Reservar(tam int, nombre string) {
	if repetido(m, nombre) != nil {
		fmt.Printf("Err: El nombre '%s' ya está en reserva.\n", nombre)
		return
	}

	lugar := hayespacio(m, tam, m.Tam)

	if lugar == -1 {
		fmt.Printf("Err: Solicitud de %d excede el espacio disponible.\n", tam)
		return
	}

	nuevoChivo := asignar(m, tam, nombre, lugar)

	fmt.Printf("ÉXITO: '%s' asignado en Inicio: %d, Peso: %d.\n", nombre, nuevoChivo.Empieza, nuevoChivo.Tam)
}

// -----------------------------------------------------------LIBERAR----------------------------------------------------

func UnionBloques(m *Manejador) {
	b := m.Ini
	iterador := 1
	for b != nil {
		if b.Sig != nil {
			if b.Tam == b.Libre && b.Sig.Tam == b.Sig.Libre && b.Tam == b.Sig.Tam {
				b.Tam = b.Tam * 2
				b.Libre = b.Libre * 2
				b.Sig = b.Sig.Sig //En Go hay recolector de basura
			} else {
				if b.Tam != b.Sig.Tam {
					b = b.Sig
				} else {
					b = b.Sig.Sig
				}
			}
		} else {
			if iterador != 2 {
				iterador++
				b = m.Ini
			} else {
				break
			}
		}
	}
}

func (m *Manejador) Liberar(nombre string) {
	bloqueActual := repetido(m, nombre)
	if bloqueActual == nil {
		fmt.Printf("Err: El nombre '%s' no esta en reserva.\n", nombre)
		return
	}
	archivoActual := bloqueActual.Chivo
	if archivoActual.Nombre == nombre { // es el primer archivo
		bloqueActual.Libre = bloqueActual.Libre + archivoActual.Tam
		bloqueActual.Chivo = archivoActual.Sig //En Go hay recolector de basura
		unionBloques(m)
	} else { // hay otros archivos
		for {
			if archivoActual.Sig.Nombre == nombre {
				bloqueActual.Libre = bloqueActual.Libre + archivoActual.Sig.Tam
				archivoActual.Sig = archivoActual.Sig.Sig //En Go hay recolector de basura
				break
			}
			archivoActual = archivoActual.Sig
		}
	}
}

// -----------------------------------------------------------MOSTRAR----------------------------------------------------
func (m *Manejador) Mostrar() {
	fmt.Printf("\n--- Estado Actual de la Memoria (Total: %d) ---\n", m.Tam)
	bloqueActual := m.Ini
	iteracion := 1
	for {
		if bloqueActual == nil {
			break
		} else {
			fmt.Printf("Bloque de Memoria %d. Rango [%d..%d]\n", iteracion, bloqueActual.Empieza, bloqueActual.Empieza+bloqueActual.Tam-1)
			fmt.Printf("\tPeso Total: %d, Peso Libre %d\n", bloqueActual.Tam, bloqueActual.Libre)
			archivoActual := bloqueActual.Chivo
			fmt.Println("\tArchivos:")
			for {
				if archivoActual == nil {
					fmt.Println("\t\t...")
					break
				} else {
					fmt.Printf("\t\tNombre: %s, Peso: %d Rango: [%d..%d]\n", archivoActual.Nombre, archivoActual.Tam, archivoActual.Empieza, archivoActual.Empieza+archivoActual.Tam-1)
					archivoActual = archivoActual.Sig
				}
			}
			bloqueActual = bloqueActual.Sig
			iteracion++
		}
	}
}
