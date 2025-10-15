package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

//------------------------------------------------ Operaciones Binarias (Vector + Vector)------------------------------------------------

// Suma realiza la suma de dos vectores: v1 + v2.
func Suma(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

// Resta realiza la resta de dos vectores: v1 - v2.
func Resta(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

// ProductoCruz realiza el producto cruz de dos vectores: v1 x v2.
func ProductoCruz(v1 Vector, v2 Vector) Vector {
	return Vector{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

// ProductoPunto calcula el producto punto (escalar) de dos vectores: v1 . v2.
func ProductoPunto(v1 Vector, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

//----------------------------------------------Operaciones con Escalares por la Derecha (Vector ⊕ Escalar)-------------------------------------

// VectorXEscalar multiplica una escala al vector: v * n.
func VectorXEscalar(v Vector, n float64) Vector {
	return Vector{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

// VectorMasEscalar suma un escalar al vector: v + n.
func VectorMasEscalar(v Vector, n float64) Vector {
	return Vector{
		X: v.X + n,
		Y: v.Y + n,
		Z: v.Z + n,
	}
}

// VectorMenosEscalar suma un escalar al vector: v + n.
func VectorMenosEscalar(v Vector, n float64) Vector {
	return Vector{
		X: v.X - n,
		Y: v.Y - n,
		Z: v.Z - n,
	}
}

// ----------------------------------Norma (&) ----------------------------------

// Norma calcula la norma del vector: |v|.
func Norma(v Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// String hace que el vector se pueda imprimir de forma legible.
func String(v Vector) string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f)", v.X, v.Y, v.Z)
}
