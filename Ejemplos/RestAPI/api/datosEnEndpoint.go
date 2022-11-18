package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Para hacer uso de esta funcion ingresamos a
// http://localhost:3000/#numero#
// por ejemplo: http://localhost:3000/7345684

func LeerDatosEnEndpoint(w http.ResponseWriter, r *http.Request) {
	// La funcion Vars de la libreria mux nos regresa un mapa de tipo
	// map[string]string en el cual accederemos a nuestros valores con el nombre asignado
	// en el endpoint.
	// Por lo cual si usamos datos de tipo numerico como en este ejemplo
	// tendremos que realizar la conversion con la libreria strconv
	x := mux.Vars(r)["VALOR"]
	num, _ := strconv.Atoi(x)

	fmt.Fprintln(w, "El parametro recibido por medio del endpoint es: "+strconv.Itoa(num))
}
