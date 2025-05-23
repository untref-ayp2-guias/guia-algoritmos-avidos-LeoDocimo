package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjercicio3(t *testing.T) {
	objetos := []Objeto{
		{"o1", 6, 20},
		{"o2", 3, 18},
		{"o3", 2, 14},
		{"o4", 5, 25},
	}
	capacidadMochila := 7

	salida := Ejercicio3(objetos, capacidadMochila)

	assert.Len(t, salida, 3, "Se esperaban 3 objetos seleccionados")

	assert.Equal(t, float32(1), salida[Objeto{"o2", 3, 18}], "Se esperaba el objeto o2 con valor 1")

	assert.Equal(t, float32(1), salida[Objeto{"o3", 2, 14}], "Se esperaba el objeto o3 con valor 1")

	assert.InDeltaf(t, float32(0.4), salida[Objeto{"o4", 5, 25}], 1e-6, "Se esperaba el objeto o4 con valor 0.4")
}
