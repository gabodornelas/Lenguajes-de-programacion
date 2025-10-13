package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("=== MENÚ PRINCIPAL ===")
        fmt.Println("RESERVAR")
        fmt.Println("LIBERAR")
        fmt.Println("MOSTRAR")
		fmt.Println("SALIR")

        opcion, _ := reader.ReadString('\n')
        opcion = strings.TrimSpace(opcion)

        switch opcion {
        case "RESERVAR":
            fmt.Println("Quieres reservar memoria")
        case "LIBERAR":
            fmt.Println("quieres liberar")
        case "MOSTRAR":
            fmt.Println("quieres mostrar a quienes han sido asignados en memoria")
		case "SALIR":
            fmt.Println("Saliendo del programa...")
            return
        default:
            fmt.Println("Opción no válida. Intenta de nuevo.")
        }

        fmt.Println()
    }
}
