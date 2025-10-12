package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Funcion que rota k espacios la cadena de caracteres w y retorna la cadena rotada
func rotar(w string, k int) string {
	// convierte a runes para poder saber la longitud de la cadena
  	runes := []rune(w)
	longitud := len(runes)
	if longitud > 0{
	    modulo := k % longitud
	    // se une la segunda parte de la cadena con la primera
	    rotado := append(runes[modulo:],runes[:modulo]...)
	    return string(rotado)
	}else{ // caso base, la cadena es vacia
		return w
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	// Leer cadena de caracteres
	fmt.Println("Ingresa una cadena de caracteres: ")
	// lee hasta conseguir salto de linea
	cadena, _ := reader.ReadString('\n')
	//limpia la cadena
	cadena = strings.TrimSpace(cadena)
	
	// Leer número entero
	fmt.Println("Ingresa un número entero no negativo: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	numero, err := strconv.Atoi(numStr)
	if err != nil || numero < 0 {
		fmt.Println("Error: no ingresaste un número válido.")
	    return
	}
	fmt.Println(rotar(cadena,numero))
}
