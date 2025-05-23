package ejercicios

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	mapa := make(map[Objeto]float32, len(objetos))
	n := len(objetos)

	// Se ordena el arreglo según el valor por peso de los objetos de mayor a menor
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			valorPorPeso := float32(objetos[j].Valor) / float32(objetos[j].Peso) //valor por peso que valen los objetos
			valorPorPeso1 := float32(objetos[j+1].Valor) / float32(objetos[j+1].Peso)
			if valorPorPeso < valorPorPeso1 {
				objetos[j], objetos[j+1] = objetos[j+1], objetos[j]
			}
		}
	}

	capacidadActual := capacidad
	for _, objeto := range objetos {
		if objeto.Peso <= capacidadActual {
			mapa[objeto] = 1
			capacidadActual -= objeto.Peso
		} else {
			mapa[objeto] = float32(capacidadActual) / float32(objeto.Peso)
			break
		}

	}

	return mapa

}
