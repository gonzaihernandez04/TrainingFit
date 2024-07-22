package funcionesCSV

import (
	"os"
	"path/filepath"
)

/*
func BuscarArchivoCSV(nombreArchivo string) (string, error) {
	// Get the absolute path of the current working directory
	directorioContenedor, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	// Append the filename to the container directory
	ruta := filepath.Join(directorioContenedor, nombreArchivo)

	fmt.Println(ruta)

	return ruta, nil
}
*/

// TODO: encontrar la forma de obtener la ruta de los archivos csv
// independientemente de que archivo se esté ejecutando esta funcion
func BuscarArchivoCSV(nombreArchivo string) (string, error) {
	// Get the absolute path of the current working directory
	rutaAbsoluta, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// fmt.Println("Ruta = ", rutaAbsoluta)

	// Join the current directory path with the filename to get the full file path
	ruta := filepath.Join(rutaAbsoluta, nombreArchivo)
	// fmt.Println("Ruta = ", ruta)

	return ruta, nil
}

func AbrirArchivoCSV(nombreArchivo string) (*os.File, error) {
	archivo, err := os.OpenFile(nombreArchivo, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	// fmt.Println(archivo)
	return archivo, err
}

// Toma el puntero a un archivo csv y lo cierra
func CerrarArchivoCSV(archivo *os.File) {
	archivo.Close()
}
