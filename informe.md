# **Aplicación de Rutinas de Entrenamiento:**

El programa está implementado en dos estructuras principales, "ejercicios" y "rutinas", y la creación automática de rutinas está implementada con diferentes técnicas de diseño de código. Esta es accesible por el usuario desde una interfaz donde tiene diferentes opciones para el uso de la aplicación.

## Implementación de Ejercicios:

1. `func NuevoEjercicio(nombre, descripción, tiempoEstimado, caloríasQuemadas, tipoEjercicio, puntosPorTipoEjercicio, dificultad)`
    > Genera una rutina a partir de los parámetros solicitados.
2. `func EjercicioExistePorNombre(nuevoEjercicio)`
    > Comprueba en la base de ejercicios si existe ya un ejercicio con ese nombre.
3. `func EjerciciosIguales(ejercicio1, ejercicio2)`
    > Comprueba si los dos ejercicios por parámetro son iguales en su totalidad.
4. `func AltaEjercicio(nuevoEjercicio)`
   > Agrega un ejercicio a la base de ejercicios disponibles, si es que no fue añadido anteriormente.
5. `func BajaEjercicio(ejercicioParaRemover)`
   > Remueve un ejercicio de la base de ejercicios, si es que encuentra uno igual en la misma.
6. `func ConsultarEjercicio(paramBusqueda)`
   > Recibe un parámetro y devuelve un slice con los ejercicios que correspondan con ese parámetro, o una lista vacía si no los hay.
7. `func ListarEjercicios(tipoEjercicioBuscado, dificultadBuscada)`
   > Lista en pantalla todos los ejercicios que concuerden con los parámetros.
8. `func ListarTodosEjercicios()`
   > Devuelve una lista con todos los ejercicios disponibles en la base de ejercicios.
9. `func ModificarEjercicio(nombre, nuevoEjercicio)`
   > Reemplaza el ejercicio con el nombre dado por el nuevo ejercicio.

## Implementación de Rutinas:

1. `func NuevaRutina(nombre, listadoDeEjercicios)`
    > Genera una rutina a partir de un array de ejercicios y el nombre que defina el usuario.
2. `func RutinaExiste(nuevaRutina)`
    > Comprueba si existe una rutina igual en la base de rutinas.
3. `func RutinasIguales(rutina1, rutina2)`
    > Comprueba si las dos rutinas por parámetro son iguales en su totalidad.
4. `func AltaRutina(nuevaRutina)`
    > Agrega una rutina a la base de rutinas disponibles, si es que no fue añadida anteriormente.
5. `func BajaRutina(rutinaParaRemover)`
    > Remueve una rutina de la base de rutinas, si es que encuentra una igual en la misma.
6. `func ConsultarRutina(nombreBuscado)`
    > Busca una rutina por nombre y si está en la base de rutinas la retorna.
7. `func ListarRutinas()`
   > Devuelve una lista con las rutinas disponibles.
8. `func ModificarRutina(nombre, nuevaRutina)`
   > Reemplaza la rutina con el nombre dado por la nueva rutina.

## Creación de Rutinas Automágicas:

1. `GeneracionAutomagicaDeRutinas1()`
   > Se genera una rutina a partir de una duración total, los tipos de ejercicios a elección y nivel de dificultad establecido. Luego se registra en la base de rutinas bajo un nombre y se retorna por pantalla.
2. `GeneracionAutomagicaDeRutinas2()`
   > Se genera una rutina a partir de la cantidad de calorías a quemar como objetivo. Luego se registra en la base de rutinas bajo un nombre y se retorna por pantalla.
3. `GeneracionAutomagicaDeRutinas3()`
   > Se genera una rutina a partir de un tipo de ejercicio a maximizar y una duración total. Luego se registra en la base de rutinas bajo un nombre y se retorna por pantalla.

## Conclusión
   > Un proyecto de este tipo y tamaño debe ser planificado con antelación y cuidado para evitar problemas a largo plazo relacionados al escalado, compresión del del codigo, e inclusión de nuevas funcionalidades, asi como retoques en las ya existentes. Lo ideal seria documentar el codigo conforme se va creando y actualizar constantemente esta documentación de forma que todos los integrantes del proyecto conozcan las funcionalidades de cada parte, o les sea mas sencillo entenderlas.