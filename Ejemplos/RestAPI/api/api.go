package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewApi() {
	// Lo primero que haremos es cerar un enrutador de gorilla
	r := mux.NewRouter().StrictSlash(true)

	// El siguiente es un ejemplo sencillo sobre como crear un endpoint del tipo GET
	// Como primer parametro colocaremos el endpoint y como segundo parametro la funcion a ejecutar
	// Es obligatorio que la funcion reciba los parametros http.ResponseWriter y *http.Request
	r.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		// Para escribir respuestas utilizaremos la funcion fmt.Fprintln
		fmt.Fprintln(w, "hola mundo desde la api en Go!")
		// al momento de levantar nuestro servidor e ingresar a http://localhost:3000/holamundo
		// nos escribira en pantalla la respuesta
	}).Methods("GET") // <- podemos ver que con el metodo Methods le decimos que tipo de peticion sera

	// Si queremos que nuestro codigo sea mas ordenado, podemos crear funciones y utilizarlas como parametro
	// las funciones las encontraremos en el archivo functions.go
	r.HandleFunc("/map", PostConMap).Methods("POST")
	r.HandleFunc("/struct", PostConStruct).Methods("POST")

	// Recibir datos por medio del endpoint:
	// utilizaremos expresiones regulares para determinar los posibles valores
	// y le colocaremos un nombre para luego poder acceder a ellos
	r.HandleFunc("/{VALOR:[0-9]+}", LeerDatosEnEndpoint).Methods("GET")

	// creamos un handler colocandole los permisos CORS para evitar inconvenientes en la comunicacion
	handler := handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)

	// la siguiente linea se encarga de levantar el servidor, como primer parametro colocamos el puerto
	// que deseamos utilizar, en este caso utilizaremos el 3000
	// como segundo parametro utilizaremos el handler que creamos anteriormente
	http.ListenAndServe(":3000", handler)
}
