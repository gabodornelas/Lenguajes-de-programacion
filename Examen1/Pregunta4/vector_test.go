package main

import (
	"math"
	"testing"
)

// Constante de tolerancia para comparar flotantes
const tolerance = 1e-9

// Estructura de prueba para todas las funciones que devuelven un Vector
type vectorTest struct {
	name string
	v1   Vector
	v2   Vector
	want Vector
}

// Estructura de prueba para todas las funciones que devuelven un float64 (escalar)
type scalarTest struct {
	name string
	v1   Vector
	v2   Vector
	want float64
}

// --- Pruebas para Operaciones Binarias (Vector + Vector) ---

func TestSuma(t *testing.T) {
	tests := []vectorTest{
		{"Suma Básica", Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{5, 7, 9}},
		{"Suma con Ceros", Vector{0, 0, 0}, Vector{10, 20, 30}, Vector{10, 20, 30}},
		{"Suma con Negativos", Vector{-1, -2, -3}, Vector{1, 2, 3}, Vector{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Suma(tt.v1, tt.v2)
			if !compareVectors(got, tt.want) {
				t.Errorf("Suma(%v, %v) = %v, want %v", tt.v1, tt.v2, got, tt.want)
			}
		})
	}
}

func TestResta(t *testing.T) {
	tests := []vectorTest{
		{"Resta Básica", Vector{5, 5, 5}, Vector{1, 2, 3}, Vector{4, 3, 2}},
		{"Resta a Cero", Vector{10, 10, 10}, Vector{10, 10, 10}, Vector{0, 0, 0}},
		{"Resta Negativa", Vector{1, 1, 1}, Vector{2, 3, 4}, Vector{-1, -2, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Resta(tt.v1, tt.v2)
			if !compareVectors(got, tt.want) {
				t.Errorf("Resta(%v, %v) = %v, want %v", tt.v1, tt.v2, got, tt.want)
			}
		})
	}
}

func TestProductoCruz(t *testing.T) {
	tests := []vectorTest{
		{"Producto Cruz (i, j, k)", Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}},
		{"Producto Cruz Cero", Vector{1, 2, 3}, Vector{1, 2, 3}, Vector{0, 0, 0}},
		{"Producto Cruz General", Vector{2, 3, 4}, Vector{5, 6, 7}, Vector{-3, 6, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductoCruz(tt.v1, tt.v2)
			if !compareVectors(got, tt.want) {
				t.Errorf("ProductoCruz(%v, %v) = %v, want %v", tt.v1, tt.v2, got, tt.want)
			}
		})
	}
}

func TestProductoPunto(t *testing.T) {
	tests := []scalarTest{
		{"Producto Punto Ortogonal", Vector{1, 0, 0}, Vector{0, 1, 0}, 0},
		{"Producto Punto General", Vector{1, 2, 3}, Vector{4, 5, 6}, 32}, // 4 + 10 + 18 = 32
		{"Producto Punto Negativo", Vector{-1, -1, -1}, Vector{1, 1, 1}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProductoPunto(tt.v1, tt.v2)
			if math.Abs(got-tt.want) > tolerance {
				t.Errorf("ProductoPunto(%v, %v) = %f, want %f", tt.v1, tt.v2, got, tt.want)
			}
		})
	}
}

// --- Pruebas para Operaciones con Escalares ---

func TestVectorXEscalar(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		n    float64
		want Vector
	}{
		{"Multiplicación Básica", Vector{1, 2, 3}, 2.0, Vector{2, 4, 6}},
		{"Multiplicación por Cero", Vector{10, 10, 10}, 0.0, Vector{0, 0, 0}},
		{"Multiplicación por Negativo", Vector{1, 2, 3}, -1.0, Vector{-1, -2, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VectorXEscalar(tt.v, tt.n)
			if !compareVectors(got, tt.want) {
				t.Errorf("VectorXEscalar(%v, %f) = %v, want %v", tt.v, tt.n, got, tt.want)
			}
		})
	}
}

func TestVectorMasEscalar(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		n    float64
		want Vector
	}{
		{"Suma Básica", Vector{1, 2, 3}, 10.0, Vector{11, 12, 13}},
		{"Suma con Cero", Vector{1, 2, 3}, 0.0, Vector{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VectorMasEscalar(tt.v, tt.n)
			if !compareVectors(got, tt.want) {
				t.Errorf("VectorMasEscalar(%v, %f) = %v, want %v", tt.v, tt.n, got, tt.want)
			}
		})
	}
}

func TestVectorMenosEscalar(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		n    float64
		want Vector
	}{
		{"Resta Básica", Vector{10, 20, 30}, 5.0, Vector{5, 15, 25}},
		{"Resta con Cero", Vector{1, 2, 3}, 0.0, Vector{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VectorMenosEscalar(tt.v, tt.n)
			if !compareVectors(got, tt.want) {
				t.Errorf("VectorMenosEscalar(%v, %f) = %v, want %v", tt.v, tt.n, got, tt.want)
			}
		})
	}
}

// --- Pruebas para Funciones Auxiliares ---

func TestNorma(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
		{"Norma Base", Vector{3, 0, 0}, 3.0},
		{"Norma 3, 4, 0 (5)", Vector{3, 4, 0}, 5.0},
		{"Norma General (1, 1, 1)", Vector{1, 1, 1}, math.Sqrt(3)},
		{"Norma Cero", Vector{0, 0, 0}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Norma(tt.v)
			if math.Abs(got-tt.want) > tolerance {
				t.Errorf("Norma(%v) = %f, want %f", tt.v, got, tt.want)
			}
		})
	}
}

func TestOperarVN(t *testing.T) {
	v := Vector{1, 1, 1}
	n := 5.0
	tests := []struct {
		name string
		op   rune
		want Vector
	}{
		{"Suma Escalar", '+', Vector{6, 6, 6}},
		{"Resta Escalar", '-', Vector{-4, -4, -4}},
		{"Multiplicación Escalar", '*', Vector{5, 5, 5}},
		{"Operador No Válido", '/', Vector{1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := operarVN(v, tt.op, n)
			if !compareVectors(got, tt.want) {
				t.Errorf("operarVN(%v, '%c', %f) = %v, want %v", v, tt.op, n, got, tt.want)
			}
		})
	}
}

func TestOperarVV(t *testing.T) {
	v1 := Vector{1, 0, 0}
	v2 := Vector{0, 1, 0}
	v3 := Vector{1, 1, 0}
	tests := []struct {
		name string
		op   rune
		want Vector
	}{
		{"Suma Vector", '+', v3},                // 1,0,0 + 0,1,0 = 1,1,0
		{"Resta Vector", '-', Vector{1, -1, 0}}, // 1,0,0 - 0,1,0 = 1,-1,0
		{"Producto Cruz", '*', Vector{0, 0, 1}}, // 1,0,0 x 0,1,0 = 0,0,1
		{"Operador No Válido", '/', v1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := operarVV(v1, tt.op, v2)
			if !compareVectors(got, tt.want) {
				t.Errorf("operarVV(%v, '%c', %v) = %v, want %v", v1, tt.op, v2, got, tt.want)
			}
		})
	}
}

// --- Functión de utilidad para la comparación de vectores con tolerancia ---

func compareVectors(v1, v2 Vector) bool {
	return math.Abs(v1.X-v2.X) < tolerance &&
		math.Abs(v1.Y-v2.Y) < tolerance &&
		math.Abs(v1.Z-v2.Z) < tolerance
}
