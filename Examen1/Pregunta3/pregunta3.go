package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run pregunta3.go <cantidad_bloques_memoria>")
		os.Exit(1)
	}

	memoriastr := os.Args[1]
	memoria, err := strconv.Atoi(memoriastr)
	if err != nil || memoria <= 0 {
		fmt.Println("Err: La cantidad de bloques debe ser un número entero positivo.")
		os.Exit(1)
	}

	gestor, err := NuevoManejador(memoria)
	if err != nil {
		fmt.Println("ERROR al inicializar el gestor:", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("======	MENÚ PRINCIPAL ======")
		fmt.Println("\tRESERVAR <cantidad> <nombre>\n\tLIBERAR <nombre>\n\tMOSTRAR\n\tSALIR")
		fmt.Print(">> ")
		entrada, _ := reader.ReadString('\n')
		entrada = strings.TrimSpace(entrada)
		partes := strings.Fields(entrada)
		accion := strings.ToUpper(partes[0])

		switch accion {
		case "RESERVAR":
			if len(partes) == 3 {
				peso, err := strconv.Atoi(partes[1])
				nombre := partes[2]
				if err != nil {
					fmt.Println("Err: La cantidad a reservar debe ser un número entero.")
					continue
				}
				gestor.Reservar(peso, nombre)
			} else {
				fmt.Println("Err: Uso incorrecto. Sintaxis: RESERVAR <cantidad> <nombre>")
			}
		case "LIBERAR":
			if len(partes) == 2 {
				nombre := partes[1]
				gestor.Liberar(nombre)
			} else {
				fmt.Println("Err: Uso incorrecto. Sintaxis: LIBERAR <nombre>")
			}
		case "MOSTRAR":
			gestor.Mostrar()
		case "SALIR":
			fmt.Println("Saliendo del simulador...")
			return
		default:
			fmt.Println("Opción no válida. Intenta de nuevo.")
		}
		fmt.Println()
	}
}
