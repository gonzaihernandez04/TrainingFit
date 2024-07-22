package rutinas_test

import (
	enums "ayp2_tp/constantes"
	"ayp2_tp/ejercicios"
	"ayp2_tp/rutinas"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*VALIDA SI SE CREA UNA RUTINA CORRECTAMENTE*/
func TestNuevaRutina(t *testing.T) {
	nuevoEjercicio := ejercicios.NuevoEjercicio("Sentadilla", "Hacer fuerza con las piernas", 40, 200, []string{"fuerza", "flexibilidad"}, []int{2, 2}, enums.INTERMEDIO)
	nuevoEjercicio2 := ejercicios.NuevoEjercicio("Press de banca", "Empujar desde un asiento acostado hacia arriba, bajar lentamente al pecho y empujar nuevamente.", 40, 200, []string{"fuerza", "pecho"}, []int{2, 2}, enums.INTERMEDIO)

	ejercicios := []*ejercicios.Ejercicio{nuevoEjercicio, nuevoEjercicio2}

	nuevaRutina := rutinas.NuevaRutina("Fuerza", ejercicios)

	assert.NotNil(t, nuevaRutina, "La rutina no deberia ser nula")
	assert.Equal(t, nuevaRutina.Nombre, "Fuerza")
}

/*VALIDA SI HAY EJERCICIOS EN UNA RUTINA*/
func TestExistenEjerciciosEnRutina(t *testing.T) {
	nuevoEjercicio := ejercicios.NuevoEjercicio("Sentadilla", "Hacer fuerza con las piernas", 40, 200, []string{"fuerza", "flexibilidad"}, []int{2, 2}, enums.INTERMEDIO)
	nuevoEjercicio2 := ejercicios.NuevoEjercicio("Press de banca", "Empujar desde un asiento acostado hacia arriba, bajar lentamente al pecho y empujar nuevamente.", 40, 200, []string{"fuerza", "pecho"}, []int{2, 2}, enums.INTERMEDIO)

	ejercicios := []*ejercicios.Ejercicio{nuevoEjercicio, nuevoEjercicio2}

	nuevaRutina := rutinas.NuevaRutina("Fuerza", ejercicios)

	for _, ejercicio := range nuevaRutina.Ejercicios {
		assert.NotNil(t, ejercicio, "Debe haber ejercicios en la rutina")
	}

}

/*Evalua si YA EXISTE RUTINA, si no existe, falla el test*/
//el test falla, ya que no encuentra la rutina.
/*
func TestYaExisteRutina(t *testing.T) {
	nuevoEjercicio := ejercicios.NuevoEjercicio("Sentadilla", "Hacer fuerza con las piernas", 40, 200, []string{"fuerza", "flexibilidad"}, []int{2, 2}, enums.INTERMEDIO)

	nuevoEjercicio2 := ejercicios.NuevoEjercicio("Press de banca", "Empujar desde un asiento acostado hacia arriba, bajar lentamente al pecho y empujar nuevamente.", 40, 200, []string{"fuerza", "pecho"}, []int{2, 2}, enums.INTERMEDIO)
	ejercicios.AltaEjercicio(nuevoEjercicio2)
	ejercicios.AltaEjercicio(nuevoEjercicio2)

	ejercicios := []*ejercicios.Ejercicio{nuevoEjercicio, nuevoEjercicio2}


	nuevaRutina := NuevaRutina("Fuerza", ejercicios)

	existe := RutinasExiste(nuevaRutina)
	fmt.Println(existe)



	assert.False(t, existe, "No existe la rutina")


}
*/

func TestYaExisteRutina(t *testing.T) {
	nuevoEjercicio := ejercicios.NuevoEjercicio("Sentadilla", "Hacer fuerza con las piernas", 40, 200, []string{"fuerza", "flexibilidad"}, []int{2, 2}, enums.INTERMEDIO)

	ejercicios.AltaEjercicio(nuevoEjercicio)

	ejercicios := []*ejercicios.Ejercicio{nuevoEjercicio}

	nuevaRutina := rutinas.NuevaRutina("Fuerza", ejercicios)
	rutinas.AltaRutina(nuevaRutina)

	booleano := rutinas.RutinasExiste(nuevaRutina)

	assert.True(t, booleano)
}
