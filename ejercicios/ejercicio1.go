package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorRecursivo(actividades []Actividad) []Actividad {

	return selectorIndiceRecursivo(actividades, 0, 0)

}

func selectorIndiceRecursivo(actividades []Actividad, indice, fin int) []Actividad {
	if indice >= len(actividades) {
		return []Actividad{}
	}

	actual := actividades[indice]

	if actual.Inicio >= fin {
		return append([]Actividad{actividades[indice]}, selectorIndiceRecursivo(actividades, indice+1, actual.Fin)...)
	}

	return selectorIndiceRecursivo(actividades, indice+1, fin)
}
