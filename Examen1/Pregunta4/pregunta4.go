package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func operarVN(v1 Vector, op rune, n float64) Vector {
	switch op {
	case '+':
		return VectorMasEscalar(v1, n)
	case '-':
		return VectorMenosEscalar(v1, n)
	case '*':
		return VectorXEscalar(v1, n)
	}
	return v1
}

func operarVV(v1 Vector, op rune, v2 Vector) Vector {
	switch op {
	case '+':
		return Suma(v1, v2)
	case '-':
		return Resta(v1, v2)
	case '*':
		return ProductoCruz(v1, v2)
	}
	return v1
}

func opera(cadena []rune, ini int) (Vector, int, float64) {
	var v1, v2 Vector
	var n, decimal float64
	n = 0
	decimal = 0
	var NoV float64
	operador := '0'
	for i := 0; i < len(cadena); i++ {
		if cadena[i] == '(' {
			v2, i, NoV = opera(cadena[i+1:], i+1)
			//volvi del parentesis
			if operador == '0' { //no hay operador
				v1 = Suma(v1, v2)
			} else {
				if operador == '%' {
					NoV = ProductoPunto(v1, v2)
				} else {
					if NoV > 0 {
						v1 = operarVN(v1, operador, NoV)
					} else {
						v1 = operarVV(v1, operador, v2)
						NoV = 0
					}
				}
			}
		}
		if i < len(cadena) {
			if cadena[i] == ')' {
				return v1, i + 1 + ini, NoV
			}
			if cadena[i] == '+' || cadena[i] == '-' || cadena[i] == '*' {
				operador = cadena[i]
			}
			if cadena[i] == '[' {
				v2.X = float64(cadena[i+1] - '0')
				v2.Y = float64(cadena[i+3] - '0')
				v2.Z = float64(cadena[i+5] - '0')
				if i != 0 && cadena[i-1] == '&' {
					n = Norma(v2)
					NoV = n
				}
				if operador == '0' { //no hay operador
					v1 = Suma(v1, v2)
				} else {
					if cadena[i-1] == '&' {
						v1 = operarVN(v1, operador, n)
						NoV = 0
					} else {
						if operador == '%' {
							NoV = ProductoPunto(v1, v2)
						} else {
							v1 = operarVV(v1, operador, v2)
							NoV = 0
						}
					}
				}
				i = i + 6
				n = 0
			}
			if cadena[i] >= '0' && cadena[i] <= '9' {

				for (cadena[i] >= '0' && cadena[i] <= '9') || cadena[i] == '.' {
					if cadena[i] >= '0' && cadena[i] <= '9' {
						n = n*10 + float64(cadena[i]-'0')
					} else {
						decimal = float64(i)
					}
					if i+1 == len(cadena) {
						break
					}
					i++
				}
				if decimal > 0 {
					n = n / (10 * (float64(i) - decimal - 1))
				}
				if operador != '0' { // hay operador (siempre deberia haber poreque las operaciones son por la derecha)
					v1 = operarVN(v1, operador, n)
					n = 0
					decimal = 0
					NoV = 0
				}
			}
		}
	}
	return v1, 0, NoV
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var v Vector
	var i int
	var NoV float64
	// leer cadena de caracteres
	fmt.Print("Ingresa una operacion con vectores, los vectores deben estar en formato [1,2,3]: ")
	// lee hasta conseguir salto de linea
	cadena, _ := reader.ReadString('\n')
	// limpia la cadena
	cadena = strings.TrimSpace(cadena)
	runas := []rune(cadena)
	v, i, NoV = opera(runas, 0)
	if NoV > 0 {
		fmt.Printf("%.2f", NoV)
	} else {
		fmt.Printf("[%.2f,%.2f,%.2f]", v.X, v.Y, v.Z)
	}
	NoV = float64(i)
}
