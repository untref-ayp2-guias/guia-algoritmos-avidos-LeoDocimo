package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	n := len(tareas)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if tareas[j].Tiempo > tareas[j+1].Tiempo {
				tareas[j], tareas[j+1] = tareas[j+1], tareas[j]
			}
		}
	}

}
