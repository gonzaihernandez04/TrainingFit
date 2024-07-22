package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNuevoEjercicio(t *testing.T) {
	ejercicio := NuevoEjercicio("Press frances", "Se recuesta sobre un banco con la cabeza sobrepasando el mismo. Agarra una barra en W con las palmas y brazos estirados y flexiona los codos hasta sobre pasar la cabeza. Se deben flexionar los codos. Repetir", 10, 50, []string{"Triceps", "Brazos"}, []int{5, 2}, 3)

	assert.Equal(t, ejercicio.Nombre, "Press frances")
	assert.NotNil(t, ejercicio, "El ejercicio no deberia ser nulo")
}

func TestEjerciciosIguales(t *testing.T) {
	ejercicio := NuevoEjercicio("Press frances", "Se recuesta sobre un banco con la cabeza sobrepasando el mismo. Agarra una barra en W con las palmas y brazos estirados y flexiona los codos hasta sobre pasar la cabeza. Se deben flexionar los codos. Repetir", 10, 50, []string{"Triceps", "Brazos"}, []int{5, 2}, 3)
	ejercicio2 := NuevoEjercicio("Press frances", "Se recuesta sobre un banco con la cabeza sobrepasando el mismo. Agarra una barra en W con las palmas y brazos estirados y flexiona los codos hasta sobre pasar la cabeza. Se deben flexionar los codos. Repetir", 10, 50, []string{"Triceps", "Brazos"}, []int{5, 2}, 3)
	assert.True(t, EjerciciosIguales(ejercicio, ejercicio2))

}

func TestEjerciciosNoIguales(t *testing.T) {
	ejercicio := NuevoEjercicio("Press frances", "Se recuesta sobre un banco con la cabeza sobrepasando el mismo. Agarra una barra en W con las palmas y brazos estirados y flexiona los codos hasta sobre pasar la cabeza. Se deben flexionar los codos. Repetir", 10, 50, []string{"Triceps", "Brazos"}, []int{5, 2}, 3)
	// Pasa el test si los ejercicios son diferentes
	ejercicio3 := NuevoEjercicio("Press francesssss", "Se recuesta sobre un banco con la cabeza sobrepasando el mismo. Agarra una barra en W con las palmas y brazos estirados y flexiona los codos hasta sobre pasar la cabeza. Se deben flexionar los codos. Repetir", 10, 50, []string{"Triceps", "Brazos"}, []int{5, 2}, 3)
	assert.False(t, EjerciciosIguales(ejercicio, ejercicio3))
}

// func TestYaExisteEjercicio(t *testing.T) {
// 	ejercicio := NuevoEjercicio("sentadillas", "Hacer fuerza con las piernas", 40, 200, []string{"fuerza", "piernas"}, []int{2, 2}, 2)
// 	assert.True(t, EjercicioExiste(ejercicio), "El ejercicio deberia existir")
// }
