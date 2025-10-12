package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

// transponer la matriz
func transponer(m [][]int, n int) [][]int {
    transpuesta := make([][]int, n)
    for i := range transpuesta {
        transpuesta[i] = make([]int, n)
        for j := 0; j < n; j++ {
            transpuesta[i][j] = m[j][i]
        }
    }
    return transpuesta
}

// multiplicar las dos matrices
func multiplicar(a, b [][]int, n int) [][]int {
    resultado := make([][]int, n)
    for i := range resultado {
        resultado[i] = make([]int, n)
        for j := 0; j < n; j++ {
            for k := 0; k < n; k++ {
                resultado[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return resultado
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Ingresa la matriz en formato [[1,2,3],[4,5,6]]:")

    input, _ := reader.ReadString('\n')

    var matriz [][]int
    err := json.Unmarshal([]byte(input), &matriz)
    if err != nil {
        fmt.Println("Error al parsear la matriz:", err)
        return
    }

    // validar que la martiz sea cuadrada
    n := len(matriz)
    columnas := len(matriz[0])
    if n != columnas{
        fmt.Println("Error: la matriz no es cuadrada.")
        return
    }
    // validar que todas las filas tengan la misma longitud
    for _, fila := range matriz{
        if len(fila) != columnas{
            fmt.Println("Error: todas las filas deben tener la misma cantidad de columnas.")
            return
        }
    }

    // calcular la transpuesta
    transpuesta := transponer(matriz, n)
    
    // multiplicar matriz por su traspuesta
    resultado := multiplicar(matriz, transpuesta, n)

    // mostrar resultado
    fmt.Println("Resultado de la multiplicación:")
    for _, fila := range resultado {
        fmt.Println(fila)
    }
}
