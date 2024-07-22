package ejercicios

import (
	enums "ayp2_tp/constantes"
	"ayp2_tp/funcionesCSV"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gocarina/gocsv"
)

type Ejercicio struct {
	Nombre                 string   `csv:"nombre"`
	Descripcion            string   `csv:"descripcion"`
	TiempoEstimado         int      `csv:"tiempoEstimado"`
	CaloriasQuemadas       int      `csv:"caloriasQuemadas"`
	TipoEjericio           []string `csv:"tipoEjericio"`
	PuntosPorTipoEjercicio []int    `csv:"puntosPorTipoEjercicio"`
	Dificultad             int      `csv:"dificultad"`
}

func NuevoEjercicio(
	nombre string,
	descripcion string,
	tiempoEstimado int,
	caloriasQuemadas int,
	tipoEjericio []string,
	puntosPorTipoEjercicio []int,
	dificultad int) *Ejercicio {
	return &Ejercicio{
		nombre,
		descripcion,
		tiempoEstimado,
		caloriasQuemadas,
		tipoEjericio,
		puntosPorTipoEjercicio,
		dificultad,
	}
}

/*
// Toma el nombre de un csv y lo abre como un puntero a archivo
func abrirArchivoCSV(nombreArchivo string) (*os.File, error) {
	archivo, err := os.OpenFile(nombreArchivo, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	return archivo, err
}

// Toma el puntero a un archivo csv y lo cierra
func cerrarArchivoCSV(archivo *os.File) {
	archivo.Close()
}
*/

// Devuelve un slice con todos los ejercicios dentro de un archivo csv abierto
func obtenerSliceEjercicios() []*Ejercicio {
	ejercicios := []*Ejercicio{}
	//ruta := "ejercicios.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("ejercicios.csv")
	archivoEjercicios, err := funcionesCSV.AbrirArchivoCSV(ruta)

	defer funcionesCSV.CerrarArchivoCSV(archivoEjercicios)

	if err != nil {
		panic(err)
	}

	if err := gocsv.UnmarshalFile(archivoEjercicios, &ejercicios); err != nil {
		panic(err)
	}

	return ejercicios
}

func EjercicioExistePorNombre(nuevoEjercicio string) bool {
	ejercicios := obtenerSliceEjercicios()

	for _, ejercicio := range ejercicios {
		if ejercicio.Nombre == nuevoEjercicio {
			return true
		}
	}

	return false
}

// Devuelve true o false si un ejercicio ya existe o no en el csv de ejercicios
func EjercicioExiste(nuevoEjercicio *Ejercicio) bool {
	ejercicios := obtenerSliceEjercicios()

	for _, ejercicio := range ejercicios {
		if reflect.DeepEqual(ejercicio, nuevoEjercicio) {
			return true
		}
	}

	return false
}

// Devuelve true o false, segun si dos ejercicios son iguales o no
func EjerciciosIguales(ejercicio1, ejercicio2 *Ejercicio) bool {
	return reflect.DeepEqual(ejercicio1, ejercicio2)
}

// Devuelve el indice de un ejercicio en la lista de ejercicios del csv, o -1 si el ejercicio no esta
func buscarIndiceEjericio(ejercicioBuscado *Ejercicio) int {
	ejercicios := obtenerSliceEjercicios()

	indiceEjercicio := -1
	for i, ejercicio := range ejercicios {
		if EjerciciosIguales(ejercicio, ejercicioBuscado) {
			indiceEjercicio = i
		}
	}

	return indiceEjercicio
}

// Agrega un ejercicio a la lista de ejercicios disponibles, si es que
// no fue añadido anteriormente
func AltaEjercicio(nuevoEjercicio *Ejercicio) {
	// Checkear si ya existe el ejercicio en la lista
	if EjercicioExiste(nuevoEjercicio) {
		return
	}

	// Abrir el archivo y cerrarlo automaticamente al salir de la funcion
	// ruta := "ejercicios/ejercicios.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("ejercicios.csv")
	archivoEjercicios, _ := funcionesCSV.AbrirArchivoCSV(ruta)

	defer funcionesCSV.CerrarArchivoCSV(archivoEjercicios)

	ejercicios := obtenerSliceEjercicios()

	// Agregar el nuevo ejercicio al slice de ejercicios
	ejercicios = append(ejercicios, nuevoEjercicio)

	// Guardar los datos del slice en el archivo csv
	if _, err := archivoEjercicios.Seek(0, 0); err != nil {
		panic(err)
	}

	gocsv.MarshalFile(&ejercicios, archivoEjercicios)
}

// Remueve un ejercicio del csv de ejercicios, si es que hay un ejercicio que concuerde con el dado
func BajaEjercicio(ejercicioParaRemover *Ejercicio) {
	if !EjercicioExiste(ejercicioParaRemover) {
		// fmt.Println("EL EJERCICIO NO ESTA EN LA LISTA")
		return
	}

	// Abrir el archivo y cerrarlo automaticamente al salir de la funcion
	// ruta := "ejercicios/ejercicios.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("ejercicios.csv")
	archivoEjercicios, _ := funcionesCSV.AbrirArchivoCSV(ruta)

	defer funcionesCSV.CerrarArchivoCSV(archivoEjercicios)

	ejercicios := obtenerSliceEjercicios()

	// Eliminar el ejercicio de la lista
	/*
		indiceEjercicio := -1
		for i, ejercicio := range ejercicios {
			if EjerciciosIguales(ejercicio, ejercicioParaRemover) {
				indiceEjercicio = i
			}
		}
	*/

	indiceEjercicio := buscarIndiceEjericio(ejercicioParaRemover)

	if indiceEjercicio >= 0 {
		ejercicios = append(ejercicios[:indiceEjercicio], ejercicios[indiceEjercicio+1:]...)
	}

	// Elimina todos los datos del csv
	if err := archivoEjercicios.Truncate(0); err != nil {
		panic(err)
	}

	// Va hasta arriba del archivo csv
	if _, err := archivoEjercicios.Seek(0, 0); err != nil {
		panic(err)
	}

	// Guarda los datos del slice en el csv
	if err := gocsv.MarshalFile(&ejercicios, archivoEjercicios); err != nil {
		log.Fatalf("Error al escribir en el archivo CSV: %s", err)
	}
}

/*
Recibe un nombre de ejercicio o un numero de calorias y devuelve un slice con los ejercicios que correspondan con ese parametro, o una lista vacia si no hay ningun ejercicio que concuerde
*/
func ConsultarEjercicio(paramBusqueda interface{}) []*Ejercicio {
	// Abrir el archivo y cerrarlo automaticamente al salir de la funcion
	//ruta := "ejercicios/ejercicios.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("ejercicios.csv")
	archivoEjercicios, _ := funcionesCSV.AbrirArchivoCSV(ruta)

	defer funcionesCSV.CerrarArchivoCSV(archivoEjercicios)

	// Crear el slice que contiene los ejercicios buscados
	var ejerciciosBuscados []*Ejercicio
	ejercicios := obtenerSliceEjercicios()

	// Si el parametro es un string, busca por nombre de ejercicio, sino busca
	// por calorias

	if ejercicioBuscado, esString := paramBusqueda.(string); esString {
		for _, ejercicio := range ejercicios {
			// Si el nombre del ejericicio contiene el ejercicio buscado, se agrega al slice
			if strings.Contains(strings.ToLower(ejercicio.Nombre), strings.ToLower(ejercicioBuscado)) {
				ejerciciosBuscados = append(ejerciciosBuscados, ejercicio)
			}
		}

	} else if caloriasMinimas, esInt := paramBusqueda.(int); esInt {
		for _, ejercicio := range ejercicios {
			// Si el nombre del ejericicio contiene el ejercicio buscado, se agrega al slice
			if ejercicio.CaloriasQuemadas >= caloriasMinimas {
				ejerciciosBuscados = append(ejerciciosBuscados, ejercicio)
			}
		}
	}

	return ejerciciosBuscados
}

// Lista en pantalla todos los ejercicios que concuerden con los parametros
func ListarEjercicios(tipoEjercicioBuscado string, dificultadBuscada int) {
	// ruta := "ejercicios/ejercicios.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("ejercicios.csv")
	archivoEjercicios, _ := funcionesCSV.AbrirArchivoCSV(ruta)
	defer funcionesCSV.CerrarArchivoCSV(archivoEjercicios)

	// Crear el slice que contiene los ejercicios buscados
	ejercicios := obtenerSliceEjercicios()

	if dificultadBuscada < 1 || 3 < dificultadBuscada {
		fmt.Println("La dificultad no es valida.")
		return
	}

	for _, ejercicio := range ejercicios {
		if ejercicio.Dificultad == dificultadBuscada {
			for _, tipo := range ejercicio.TipoEjericio {
				if tipo == tipoEjercicioBuscado {
					fmt.Println(ejercicio.Nombre)
				}
			}
		}
	}
}

// Crea funcion para generacion de rutinas automaticas que verifica si existe un tipo de ejercicio
func ExisteTipoDeEjercicio(tiposEjercicioBuscado *[]string) bool {

	// Transformo siempre la primera letra en mayuscula
	/*
		for i, tipo := range tiposEjercicioBuscado {
			runas := []rune(tipo)
			runas[0] = unicode.ToLower(runas[0])
			tiposEjercicioBuscado[i] = string(runas)
		}
	*/

	tiposEjercicioBuscadoLower := *tiposEjercicioBuscado
	for i := range tiposEjercicioBuscadoLower {
		tiposEjercicioBuscadoLower[i] = strings.ToLower(tiposEjercicioBuscadoLower[i])
	}

	tiposEjercicioBuscado = &tiposEjercicioBuscadoLower

	//Un posible refactor es crear un diccionario con todos los tipos y verificar si el que ingresa el usuario existe como tipo
	tipos := []string{
		enums.FUERZA,
		enums.AGILIDAD,
		enums.FLEXIBILIDAD,
		enums.HOMBROS,
		enums.BICEPS,
		enums.PECHO,
		enums.PIERNAS,
		enums.ESPALDA,
		enums.CARDIO,
		enums.RESISTENCIA,
	}

	// Hago funcion contains tipo con un mapa para evitar el n al cuadrado por O(1) al acceso de datos en mapa
	existe := containsTipo(tipos, tiposEjercicioBuscadoLower)

	return existe

}

func containsTipo(tipos []string, tiposEjercicioBuscado []string) bool {
	tipoMap := make(map[string]bool, len(tiposEjercicioBuscado))

	for _, tipoBuscado := range tiposEjercicioBuscado {
		tipoMap[tipoBuscado] = true
	}

	// Verificar si alguno de los tipos está en el mapa
	for _, tipo := range tipos {
		// utilizo la clave tipo para verificar que exista en el mapa
		if _, found := tipoMap[tipo]; found {
			return true
		}
	}
	return false
}
func ListarTodosEjercicios() []*Ejercicio {
	// ruta := "ejercicios/ejercicios.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("ejercicios.csv")
	archivoEjercicios, _ := funcionesCSV.AbrirArchivoCSV(ruta)
	defer funcionesCSV.CerrarArchivoCSV(archivoEjercicios)

	// Crear el slice que contiene los ejercicios buscados
	ejercicios := obtenerSliceEjercicios()

	return ejercicios
}

// Reemplaza el ejercicio con el nombre dado por el nuevo ejercicio
func ModificarEjercicio(nombre string, nuevoEjercicio *Ejercicio) {
	/*ESPECIFICO CORRECTAMENTE LA RUTA.*/

	// ruta := "ejercicios/ejercicios.csv"
	ruta, _ := funcionesCSV.BuscarArchivoCSV("ejercicios.csv")
	archivoEjercicios, _ := funcionesCSV.AbrirArchivoCSV(ruta)
	defer funcionesCSV.CerrarArchivoCSV(archivoEjercicios)

	// Crear el slice que contiene los ejercicios buscados
	ejercicios := obtenerSliceEjercicios()
	indiceEjercicio := buscarIndiceEjericio(ConsultarEjercicio(nombre)[0])

	ejercicios[indiceEjercicio] = nuevoEjercicio

	if _, err := archivoEjercicios.Seek(0, 0); err != nil {
		panic(err)
	}

	gocsv.MarshalFile(&ejercicios, archivoEjercicios)
}

// Creo una funcion para obtener el tiempo minimo de ejercicio para una rutina automagica

// func main() {
// 	nuevoEjercicio1 := NuevoEjercicio("Sentadilla", "Hacer fuerza con las piernas", 40, 200, []string{enums.FUERZA, enums.PIERNAS}, []int{2, 2}, enums.INTERMEDIO)

// 	nuevoEjercicio2 := NuevoEjercicio("Press de hombro", "Hacer fuerza con los hombros", 70, 500, []string{enums.HOMBROS}, []int{2}, enums.PRINCIPIANTE)
// 	nuevoEjercicio3 := NuevoEjercicio("Biceps", "Hacer fuerza con los biceps", 20, 100, []string{enums.BRAZOS, enums.FUERZA}, []int{1, 2}, enums.PRINCIPIANTE)

// 	AltaEjercicio(nuevoEjercicio1)
// 	AltaEjercicio(nuevoEjercicio1)
// 	AltaEjercicio(nuevoEjercicio2)
// 	AltaEjercicio(nuevoEjercicio2)
// 	// AltaEjercicio(nuevoEjercicio3)
// 	AltaEjercicio(nuevoEjercicio3)

// 	// BajaEjercicio(nuevoEjercicio1)

// 	// Consultar ejercicios por nombre
// 	ejercicios := ConsultarEjercicio(" ")

// 	for _, ejercicio := range ejercicios {
// 		fmt.Println(ejercicio.Nombre)
// 		fmt.Println()
// 	}

// 	// Consultar ejercicios por calorias
// 	ejercicios = ConsultarEjercicio(200)

// 	fmt.Println("POR CALORIAS")
// 	for _, ejercicio := range ejercicios {
// 		fmt.Println(ejercicio.Nombre)
// 		fmt.Println(ejercicio.CaloriasQuemadas)
// 		fmt.Println()
// 	}

// 	ListarEjercicios("dsasdd", 12)

// 	ModificarEjercicio("sentadilla",
// 		NuevoEjercicio("Sentadilla", "Hacer fuerza con las piernas", 40, 300, []string{"fuerza", "piernas"}, []int{2, 2}, enums.INTERMEDIO),
// 	)

// }
