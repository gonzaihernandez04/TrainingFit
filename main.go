package main

import (
	enums "ayp2_tp/constantes"
	"ayp2_tp/ejercicios"
	"ayp2_tp/rutinas"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"

	"time"
)

func main() {
	fin := false
	// fmt.Println("Hello, world")
	eleccion := 0
	color.HiBlue("---------- Bienvenido/a a la aplicacion de Training Fit ----------\n\n")

	// Quite for que no es necesario, se valida dos veces eleccion a traves del switch y un if.

	for !fin {
		imprimirEleccionesDelMenu()

		fmt.Scanln(&eleccion)
		fmt.Println(eleccion)

		if eleccion < 1 || 10 < eleccion {
			fmt.Println("La eleccion no es valida")
		}

		switch eleccion {
		// Listar rutinas
		case 1:
			listarRutinas()

		// Agregar rutinas
		case 2:
			agregarRutinas()

		// Eliminar rutinas
		case 3:
			eliminarRutina()

		// Listar ejercicios
		case 4:
			listarEjercicios()

		// Agregar ejercicios
		case 5:
			agregarEjercicios()

		// Eliminar ejercicios
		case 6:
			eliminarEjercicios()

		case 7:
			generacionAutomagicaDeRutinas2()

		case 8:
			generacionAutomagicaDeRutinas1()

		case 9:
			generacionAutomagicaDeRutinas3()

		// Salir del programa
		case 10:
			fmt.Println("Gracias por usar la aplicacion de rutinas...")
			fin = true

		default:
			panic("Ha ocurrido un error durante la ejecucion el programa, la eleccion no se tomo de forma correcta.")
		}

	}

}

func imprimirEleccionesDelMenu() {
	fmt.Println()
	color.Green(" 1. Listar rutinas 📜")
	fmt.Println(" 2. Agregar rutinas 🏃")
	color.HiRed(" 3. Eliminar rutinas")
	fmt.Println(" 4. Listar ejercicios 📜")
	fmt.Println(" 5. Agregar ejercicios 🏃")
	color.HiRed(" 6. Eliminar ejercicios")
	color.Yellow(" 7. Generar rutina automagica por calorias a quemar ✨")
	color.HiYellow(" 8. Generar rutina automagica priorizando tiempo, tipo de ejercicio y dificultad ✨")
	fmt.Println(" 9. Generar rutina automagica por maximizacion de tipos y duracion ✨")
	color.Red(" 10. Salir")
	fmt.Println()
	fmt.Printf("Elija una opcion (numero): ")
}

func listarRutinas() {
	fmt.Println("----- Rutinas disponibles: -----")
	rutinasDisponibles := rutinas.ListarRutinas()

	for _, rutina := range rutinasDisponibles {
		fmt.Println("｢", rutina.Nombre, "｣")
		fmt.Println("    ", "Duracion estimada: ", rutina.DuracionEstimada)
		fmt.Println("    ", "Calorias quemadas: ", rutina.CaloriasQuemadas)
		fmt.Println("    ", "Tipo de ejercicio: ", rutina.TipoEjericio)

		fmt.Println("    ", "Tipo de ejercicio y puntaje: ")
		for i := range rutina.TipoEjericio {
			fmt.Println("\t -", rutina.TipoEjericio[i], " -> ", rutina.PuntosPorTipoEjercicio[i])
		}

		// fmt.Println("    ", "Dificultad: ", rutina.Dificultad)
		fmt.Printf("     Dificultad: ")
		if rutina.Dificultad == enums.PRINCIPIANTE {
			fmt.Println("Principiante")
		}
		if rutina.Dificultad == enums.INTERMEDIO {
			fmt.Println("Intermedio")
		}
		if rutina.Dificultad == enums.AVANZADO {
			fmt.Println("Avanzado")
		}

		fmt.Println("    ", "Ejercicios:")
		for _, ejercicio := range rutina.Ejercicios {
			fmt.Println("\t -", ejercicio.Nombre)
		}

		fmt.Println()
	}

	fmt.Println("--------------------------------")
}

func agregarRutinas() {
	var nombreRutina string
	fmt.Printf("Ingrese el nombre de la nueva rutina: ")

	reader := bufio.NewReader(os.Stdin)
	nombreRutina, _ = reader.ReadString('\n')
	nombreRutina = strings.TrimSpace(nombreRutina)

	var ejericiciosElejidos string
	ejerciciosSonValidos := false

	fmt.Println("Elija los ejercicios para añadir a la rutina separados por comas (ej1,ej2,etc.)")
	fmt.Printf("Ejercicios: ")
	fmt.Scanln(&ejericiciosElejidos)
	listaNombresEjercicios := strings.Split(strings.Trim(ejericiciosElejidos, ","), ",")

	for !ejerciciosSonValidos {
		if !ejerciciosValidos(listaNombresEjercicios) {
			fmt.Println("Los ejercicios no son validos, elija nuevamente...")
		} else {
			var listaEjercicios []*ejercicios.Ejercicio

			for _, nombreEjercicio := range listaNombresEjercicios {
				listaEjercicios = append(listaEjercicios, ejercicios.ConsultarEjercicio(nombreEjercicio)...)
			}

			/*
				nuevaRutina := rutinas.NuevaRutina(nombreRutina, listaEjercicios)
				rutinas.AltaRutina(nuevaRutina)
			*/
			rutinas.AltaRutina(rutinas.NuevaRutina(nombreRutina, listaEjercicios))

			fmt.Println("Rutina ", nombreRutina, " agregada exitosamente...")

			// TODO: se podria borrar la variable ejerciciosSonValidos y dejar el loop como uno infinito
			ejerciciosSonValidos = true

			fmt.Println("--------------------------------")

			return
		}

		fmt.Println("Elija los ejercicios para añadir a la rutina separados por espacios (ej1,ej2,etc.)")
		fmt.Printf("Ejercicios: ")
		fmt.Scanln(&ejericiciosElejidos)
		listaNombresEjercicios = strings.Split(strings.Trim(ejericiciosElejidos, ","), ",")
	}
}

func ejerciciosValidos(listaEjercicios []string) bool {
	for _, ejercicio := range listaEjercicios {
		if !ejercicios.EjercicioExistePorNombre(ejercicio) {
			return false
		}
	}

	return true
}

func eliminarRutina() {
	var nombreRutina string
	fmt.Printf("Ingrese el nombre de la rutina que quiere eliminar: ")
	fmt.Scanln(&nombreRutina)

	rutinasDisponibles := rutinas.ConsultarRutina(nombreRutina)

	if len(rutinasDisponibles) == 0 {
		fmt.Println("No existe una rutina con el nombre ingresado")
		return
	}

	rutinas.BajaRutina(rutinasDisponibles[0])

	fmt.Println("Rutina ", nombreRutina, " eliminada exitosamente...")

	fmt.Println("--------------------------------")
}

func listarEjercicios() {
	fmt.Println("----- Ejercicios disponibles: -----")
	ejerciciosDisponibles := ejercicios.ListarTodosEjercicios()

	for _, ejercicio := range ejerciciosDisponibles {
		fmt.Println("｢", ejercicio.Nombre, "｣")
		fmt.Println("    ", "Descripcion: ", ejercicio.Descripcion)
		fmt.Println("    ", "Duracion estimada: ", ejercicio.TiempoEstimado)
		fmt.Println("    ", "Calorias quemadas: ", ejercicio.CaloriasQuemadas)

		fmt.Println("    ", "Tipo de ejercicio y puntaje: ")
		for i := range ejercicio.TipoEjericio {
			fmt.Println("\t -", ejercicio.TipoEjericio[i], " -> ", ejercicio.PuntosPorTipoEjercicio[i])
		}

		// fmt.Println("    ", "Dificultad: ", rutina.Dificultad)
		fmt.Printf("Dificultad: ")
		if ejercicio.Dificultad == 1 {
			fmt.Println("Principiante")
		}
		if ejercicio.Dificultad == 2 {
			fmt.Println("Intermedio")
		}
		if ejercicio.Dificultad == 3 {
			fmt.Println("Avanzado")
		}

		fmt.Println()
	}

	fmt.Println("--------------------------------")
}

func agregarEjercicios() {
	var nombre, descripcion, tipoEjericioString, puntosPorTipoEjercicioString string
	var tiempoEstimado, caloriasQuemadas, dificultad int
	var tipoEjericio []string
	var puntosPorTipoEjercicio []int

	fmt.Println("Ingrese los datos del nuevo ejercicio rutina: ")
	// Para tomar los espacios en blanco en el string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Nombre: ")
	nombre, _ = reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Printf("Descripcion: ")
	descripcion, _ = reader.ReadString('\n')
	descripcion = strings.TrimSpace(descripcion)

	fmt.Printf("Tiempo estimado: ")
	fmt.Scanln(&tiempoEstimado)
	if tiempoEstimado <= 0 {
		fmt.Println("El tiempo estimado debe ser mayor a 0...")
		return
	}

	fmt.Printf("Calorias quemadas: ")
	fmt.Scanln(&caloriasQuemadas)
	if caloriasQuemadas <= 0 {
		fmt.Println("La cantidad de calorias quemadas debe ser mayor a 0...")
		return
	}

	fmt.Printf("Dificultad (1, 2 o 3): ")
	fmt.Scanln(&dificultad)
	if tiempoEstimado < 1 || 3 < tiempoEstimado {
		fmt.Println("La dificultad solo puede variar entre 1, 2 o 3...")
		return
	}

	fmt.Printf("Categorias del ejercicio separdas por coma (ej1,ej2,etc): ")
	fmt.Scanln(&tipoEjericioString)
	tipoEjericio = strings.Split(tipoEjericioString, ",")

	fmt.Printf("Puntos de cada categoria separdas por coma (p1,p2,etc): ")
	fmt.Scanln(&puntosPorTipoEjercicioString)
	for _, punto := range strings.Split(tipoEjericioString, ",") {
		n, _ := strconv.ParseInt(punto, 10, 64)

		puntosPorTipoEjercicio = append(puntosPorTipoEjercicio, int(n))
	}

	nuevoEjercicio := ejercicios.NuevoEjercicio(
		nombre,
		descripcion,
		tiempoEstimado,
		caloriasQuemadas,
		tipoEjericio,
		puntosPorTipoEjercicio,
		dificultad,
	)

	ejercicios.AltaEjercicio(nuevoEjercicio)

	fmt.Println("Ejercicio ", nombre, " agregado exitosamente...")

	fmt.Println("--------------------------------")
}

func eliminarEjercicios() {
	var nombreEjercicio string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Ingrese el nombre del ejercicio que quiere eliminar: ")
	nombreEjercicio, _ = reader.ReadString('\n')
	nombreEjercicio = strings.TrimSpace(nombreEjercicio)

	ejerciciosDisponibles := ejercicios.ConsultarEjercicio(nombreEjercicio)

	if len(ejerciciosDisponibles) == 0 {
		fmt.Println("No existe un ejercicio con el nombre ingresado...")
		return
	}

	ejercicios.BajaEjercicio(ejerciciosDisponibles[0])

	fmt.Println("Rutina ", nombreEjercicio, " eliminada exitosamente...")

	fmt.Println("--------------------------------")
}

func generacionAutomagicaDeRutinas1() {
	var nombre string
	var duracionTotalRutina int
	var fin bool
	var tipoEjercicio string
	var dificultad int
	var eleccion int

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Ingrese el nombre de su nueva rutina mágica: ")
	nombre, _ = reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	for !fin {
		// Entrada de duración de la rutina
		for {
			fmt.Printf("Ingrese la duración máxima de la rutina (mínimo 45 minutos): ")
			_, err := fmt.Scanf("%d\n", &duracionTotalRutina)
			if err != nil {
				fmt.Println("Error al leer la duración. Por favor, ingrese un número válido.")
				continue
			}
			if duracionTotalRutina < 45 {
				fmt.Println("Duración mínima requerida es de 45 minutos.")
				continue
			}
			break
		}

		// Selección del tipo de ejercicio
		for {
			fmt.Printf("Ingrese el tipo de ejercicio especial (Ej: Fuerza): Solo 1")
			fmt.Scanln(&tipoEjercicio)
			tipos := make([]string, 0)
			tipos = append(tipos, tipoEjercicio)
			if !ejercicios.ExisteTipoDeEjercicio(&tipos) {
				fmt.Printf("No existe el tipo de ejercicio '%v'. Intente nuevamente.\n", tipoEjercicio)

			}
			break
		}

		// Selección de la dificultad
		for {
			fmt.Println("Nivel De dificultad:")
			fmt.Println("1. Principiante")
			fmt.Println("2. Intermedio")
			fmt.Println("3. Avanzado")
			fmt.Scanln(&eleccion)

			switch eleccion {
			case 1:
				dificultad = enums.PRINCIPIANTE
			case 2:
				dificultad = enums.INTERMEDIO
			case 3:
				dificultad = enums.AVANZADO
			default:
				fmt.Println("Opción inválida. Intente nuevamente.")
				continue
			}

			break
		}

		// Mostrar la rutina generada
		color.Cyan("\nEspere, estamos generando una rutina...\n")
		time.Sleep(2 * time.Second)

		fmt.Printf("\nElegiste:\n\n- Nombre de la rutina: %v\n- Duración Máxima: %v minutos\n- Tipo de ejercicios a incluir: %v\n- Dificultad: %v\n", nombre, duracionTotalRutina, tipoEjercicio, dificultad)

		fin = true

		rutina := rutinas.GeneracionAutomagicaDeRutinas1_Aux(nombre, duracionTotalRutina, tipoEjercicio, dificultad)
		rutinas.AltaRutina(rutina)

	}
}

func generacionAutomagicaDeRutinas2() {
	var nombre string
	var caloriasTotales int

	fmt.Printf("Ingrese el nombre de la nueva rutina: ")
	fmt.Scanln(&nombre)

	fmt.Printf("Calorias totales que desea quemar: ")
	fmt.Scanln(&caloriasTotales)

	nuevaRutina := rutinas.GeneracionAutomagicaDeRutinas2_Aux(nombre, caloriasTotales)
	// fmt.Print(nuevaRutina)
	fmt.Print("Rutina ", nombre, " añadida correctamente...")

	rutinas.AltaRutina(nuevaRutina)
}

func generacionAutomagicaDeRutinas3() {
	var nombre string
	var tipoEjercicio string
	var duracionMaximaRutina int = 0
	var tiposElegidos []string
	var fin bool = false

	// Se crea un reader y un ReadString para que el usuario pueda insertar espacios y no haya errores
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("Ingrese el nombre de su nueva rutina magica: ")
		nombre, _ = reader.ReadString('\n')
		nombre = strings.TrimSpace(nombre)
		if nombre != "" {
			break
		}
		continue
	}
	for !fin {

		for {
			fmt.Printf("Ingrese la duracion maxima de la rutina de un minimo de 45 minutos(ingresar numero entero ej: 45) \n")
			// Quita elementos del buffer, eliminando comportamientos como repeticion de mensajes
			_, err := fmt.Scanf("%d\n", &duracionMaximaRutina)

			if err != nil {
				fmt.Println("Error al leer la duración. Por favor, ingrese un número.")
				continue
			}

			if duracionMaximaRutina+1 < rutinas.ObtenerMinimoTiempoRutina() {
				fmt.Println("Duración invalida. Ingrese una nuevamente...")
				continue
			}
			break
		}
		//

		// Verificar que existan los tipos
		for {
			fmt.Printf("Que desea maximizar? (Fuerza,Resistencia,Flexibilidad,Cardio,Hombros,Biceps,Pecho,Piernas,Espalda,Gemelos,Gluteos): Separados por coma \n")
			fmt.Scanln(&tipoEjercicio)
			tiposElegidos = strings.Split(tipoEjercicio, ",")
			existe := ejercicios.ExisteTipoDeEjercicio(&tiposElegidos)
			if !existe {
				fmt.Printf("No existe el tipo de ejercicio %v ", tipoEjercicio)
				continue
			}

			fin = true
			break
		}

	}

	color.Cyan("\nEspere, estamos generando una rutina...\n ")
	time.Sleep(2 * time.Second)

	fmt.Printf("\n Elegiste \n\n - Nombre de la rutina: -> %v  \n\n - Duracion Maxima: -> %v minutos \n\n - Tipo de maximizacion: -> %v \n \n", nombre, duracionMaximaRutina, tipoEjercicio)

	nuevaRutina := rutinas.GenerarRutinaAutomagica3_Aux(nombre, tiposElegidos, duracionMaximaRutina)

	fmt.Print("Rutina ", nombre, " añadida correctamente...")

	rutinas.AltaRutina(nuevaRutina)

}
