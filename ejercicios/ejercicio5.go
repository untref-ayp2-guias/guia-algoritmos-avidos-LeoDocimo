package ejercicios

var billetes = []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

func Cambiar(cantidad int) map[int]int {
	cambio := make(map[int]int)
	for _, denominacion := range billetes {
		if cantidad >= denominacion {
			cantidad_billetes := cantidad / denominacion
			cambio[denominacion] = cantidad_billetes
			cantidad = cantidad % denominacion
		}
	}
	return cambio
}
